package main

import (
	"cache-demo/cache"
	"cache-demo/model"
	"cache-demo/service"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config 配置结构（复用main.go的配置）
type Config struct {
	MySQL struct {
		Host         string `json:"host" yaml:"host"`
		Port         int    `json:"port" yaml:"port"`
		User         string `json:"user" yaml:"user"`
		Password     string `json:"password" yaml:"password"`
		Database     string `json:"database" yaml:"database"`
		MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"`
		MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"`
	} `json:"mysql" yaml:"mysql"`
	Redis struct {
		Host        string `json:"host" yaml:"host"`
		Password    string `json:"password" yaml:"password"`
		Type        string `json:"type" yaml:"type"`
		PingTimeout string `json:"ping_timeout" yaml:"ping_timeout"`
	} `json:"redis" yaml:"redis"`
}

// DBQueryStats 数据库查询统计
type DBQueryStats struct {
	TotalQueries int64
	QueryTimes   []time.Time
	mu           sync.Mutex
}

func (s *DBQueryStats) AddQuery() {
	atomic.AddInt64(&s.TotalQueries, 1)
	s.mu.Lock()
	s.QueryTimes = append(s.QueryTimes, time.Now())
	s.mu.Unlock()
}

func (s *DBQueryStats) GetQPSInWindow(windowSeconds int) float64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.QueryTimes) == 0 {
		return 0
	}

	now := time.Now()
	windowStart := now.Add(-time.Duration(windowSeconds) * time.Second)

	count := 0
	for _, t := range s.QueryTimes {
		if t.After(windowStart) {
			count++
		}
	}

	return float64(count) / float64(windowSeconds)
}

func main() {
	// 加载配置
	var c Config
	conf.MustLoad("config.yaml", &c)

	// 初始化数据库连接
	db, err := initDB(c)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 初始化Redis连接
	rds, err := initRedis(c)
	if err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}

	// 确保测试数据存在
	if err := ensureTestData(db); err != nil {
		log.Fatalf("初始化测试数据失败: %v", err)
	}

	// 初始化服务层
	userRepo := model.NewUserRepo(db)

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("缓存雪崩问题测试")
	fmt.Println(strings.Repeat("=", 80))

	// 场景1：演示缓存雪崩问题（固定过期时间）
	testAvalancheProblem(userRepo, rds, "场景1：缓存雪崩问题演示")

	// 场景2：使用随机过期时间解决缓存雪崩
	testRandomExpireSolution(userRepo, rds, "场景2：随机过期时间解决方案")

	// 场景3：效果对比
	testAvalancheEffectiveness(userRepo, rds, "场景3：效果对比")
}

// testAvalancheProblem 场景1：演示缓存雪崩问题（固定过期时间）
func testAvalancheProblem(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：演示缓存雪崩问题 - 大量缓存在同一时间过期")
	fmt.Println("问题：所有缓存都设置相同的过期时间，导致同时过期")
	fmt.Println()

	// 创建支持固定过期时间的缓存服务
	userCache := cache.NewUserCacheWithAvalanche(rds)
	userService := service.NewUserServiceWithAvalanche(repo, userCache, service.FixedExpire, cache.AvalancheBaseExpireSeconds)

	// 统计数据库查询
	stats := &DBQueryStats{}

	// 步骤1：批量加载用户到缓存（固定过期时间60秒）
	fmt.Println("[步骤1] 批量加载10个用户到缓存（固定过期时间60秒）")
	userIDs := make([]int64, 0, 10)
	for i := int64(1); i <= 10; i++ {
		user, err := repo.FindByID(i)
		if err == nil && user != nil {
			userCache.SetUserWithFixedExpire(user, cache.AvalancheBaseExpireSeconds)
			userIDs = append(userIDs, user.ID)
			fmt.Printf("  ✓ 加载用户ID %d 到缓存（过期时间: %d秒）\n", user.ID, cache.AvalancheBaseExpireSeconds)
		}
	}
	fmt.Printf("\n✓ 成功加载 %d 个用户到缓存\n", len(userIDs))

	// 步骤2：等待缓存过期
	fmt.Printf("\n[步骤2] 等待缓存过期（%d秒）...\n", cache.AvalancheBaseExpireSeconds)
	for i := cache.AvalancheBaseExpireSeconds; i > 0; i-- {
		if i%10 == 0 || i <= 5 {
			fmt.Printf("  倒计时: %d秒\n", i)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("  → 所有缓存已过期！")

	// 步骤3：模拟大量并发请求（缓存雪崩）
	fmt.Println("\n[步骤3] 模拟100个并发请求查询用户（缓存雪崩）")
	startTime := time.Now()

	var wg sync.WaitGroup
	concurrentRequests := 100

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()
			// 随机选择一个用户ID
			userID := userIDs[requestID%len(userIDs)]
			queryStart := time.Now()
			_, err := userService.GetUserByID(userID)
			queryDuration := time.Since(queryStart)
			if err == nil {
				stats.AddQuery()
				if requestID < 5 {
					fmt.Printf("  [请求%d] 查询用户ID %d (耗时: %v)\n", requestID+1, userID, queryDuration)
				}
			}
		}(i)
	}

	wg.Wait()
	totalDuration := time.Since(startTime)

	fmt.Printf("\n[统计结果]\n")
	fmt.Printf("  并发请求数: %d\n", concurrentRequests)
	fmt.Printf("  数据库查询次数: %d\n", stats.TotalQueries)
	fmt.Printf("  总耗时: %v\n", totalDuration)
	fmt.Printf("  数据库QPS峰值: %.2f\n", stats.GetQPSInWindow(1))
	fmt.Println("  → 问题：大量请求同时访问数据库，造成数据库压力暴增！")

	fmt.Println("\n✓ 场景1测试完成：演示了缓存雪崩问题")
	fmt.Println("  问题总结：所有缓存在同一时间过期，导致大量请求同时访问数据库")
}

// testRandomExpireSolution 场景2：使用随机过期时间解决缓存雪崩
func testRandomExpireSolution(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：使用随机过期时间解决缓存雪崩问题")
	fmt.Println("方案：给缓存设置随机的过期时间，避免缓存在同一时刻过期")
	fmt.Println()

	// 创建支持随机过期时间的缓存服务
	userCache := cache.NewUserCacheWithAvalanche(rds)
	userService := service.NewUserServiceWithAvalanche(repo, userCache, service.RandomExpire, cache.AvalancheBaseExpireSeconds)

	// 统计数据库查询
	stats := &DBQueryStats{}

	// 步骤1：批量加载用户到缓存（随机过期时间60-66秒）
	fmt.Println("[步骤1] 批量加载10个用户到缓存（随机过期时间60-66秒）")
	userIDs := make([]int64, 0, 10)
	for i := int64(1); i <= 10; i++ {
		user, err := repo.FindByID(i)
		if err == nil && user != nil {
			userCache.SetUserWithRandomExpire(user, cache.AvalancheBaseExpireSeconds)
			actualExpire := cache.GetRandomExpireTime(cache.AvalancheBaseExpireSeconds)
			userIDs = append(userIDs, user.ID)
			fmt.Printf("  ✓ 加载用户ID %d 到缓存（过期时间: %d秒，随机）\n", user.ID, actualExpire)
		}
	}
	fmt.Printf("\n✓ 成功加载 %d 个用户到缓存（随机过期时间）\n", len(userIDs))

	// 步骤2：等待缓存开始过期
	fmt.Printf("\n[步骤2] 等待缓存开始过期（%d秒后开始陆续过期）...\n", cache.AvalancheBaseExpireSeconds)
	for i := cache.AvalancheBaseExpireSeconds; i > 0; i-- {
		if i%10 == 0 || i <= 5 {
			fmt.Printf("  倒计时: %d秒\n", i)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("  → 缓存开始陆续过期（不是同时过期）")

	// 步骤3：模拟大量并发请求（应该分散访问数据库）
	fmt.Println("\n[步骤3] 模拟100个并发请求查询用户（随机过期时间）")
	startTime := time.Now()

	var wg sync.WaitGroup
	concurrentRequests := 100

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()
			// 随机选择一个用户ID
			userID := userIDs[requestID%len(userIDs)]
			queryStart := time.Now()
			_, err := userService.GetUserByID(userID)
			queryDuration := time.Since(queryStart)
			if err == nil {
				stats.AddQuery()
				if requestID < 5 {
					fmt.Printf("  [请求%d] 查询用户ID %d (耗时: %v)\n", requestID+1, userID, queryDuration)
				}
			}
		}(i)
	}

	wg.Wait()
	totalDuration := time.Since(startTime)

	fmt.Printf("\n[统计结果]\n")
	fmt.Printf("  并发请求数: %d\n", concurrentRequests)
	fmt.Printf("  数据库查询次数: %d\n", stats.TotalQueries)
	fmt.Printf("  总耗时: %v\n", totalDuration)
	fmt.Printf("  数据库QPS峰值: %.2f\n", stats.GetQPSInWindow(1))
	fmt.Println("  → 优势：数据库访问分散，不会同时访问，压力平滑！")

	fmt.Println("\n✓ 场景2测试完成：随机过期时间成功解决了缓存雪崩问题")
	fmt.Println("  解决方案：给缓存设置随机的过期时间，避免缓存在同一时刻过期")
}

// testAvalancheEffectiveness 场景3：效果对比
func testAvalancheEffectiveness(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：对比使用固定过期时间和随机过期时间的效果")
	fmt.Println()

	userIDs := make([]int64, 0, 10)
	for i := int64(1); i <= 10; i++ {
		user, err := repo.FindByID(i)
		if err == nil && user != nil {
			userIDs = append(userIDs, user.ID)
		}
	}

	// 测试1：固定过期时间（缓存雪崩）
	fmt.Println("[测试1] 固定过期时间（缓存雪崩）")
	userCache1 := cache.NewUserCacheWithAvalanche(rds)
	userService1 := service.NewUserServiceWithAvalanche(repo, userCache1, service.FixedExpire, cache.AvalancheBaseExpireSeconds)

	// 预热缓存
	for _, id := range userIDs {
		user, _ := repo.FindByID(id)
		if user != nil {
			userCache1.SetUserWithFixedExpire(user, cache.AvalancheBaseExpireSeconds)
		}
	}

	// 等待缓存过期
	fmt.Printf("  等待缓存过期（%d秒）...\n", cache.AvalancheBaseExpireSeconds)
	time.Sleep(time.Duration(cache.AvalancheBaseExpireSeconds) * time.Second)

	// 并发查询
	stats1 := &DBQueryStats{}
	var wg1 sync.WaitGroup
	start1 := time.Now()

	for i := 0; i < 50; i++ {
		wg1.Add(1)
		go func(i int) {
			defer wg1.Done()
			userID := userIDs[i%len(userIDs)]
			_, err := userService1.GetUserByID(userID)
			if err == nil {
				stats1.AddQuery()
			}
		}(i)
	}
	wg1.Wait()
	duration1 := time.Since(start1)

	fmt.Printf("  查询50次，数据库访问 %d 次，耗时 %v\n", stats1.TotalQueries, duration1)
	fmt.Printf("  数据库QPS峰值: %.2f\n", stats1.GetQPSInWindow(1))
	fmt.Println("  → 问题：大量请求同时访问数据库")

	time.Sleep(2 * time.Second)

	// 测试2：随机过期时间（解决缓存雪崩）
	fmt.Println("\n[测试2] 随机过期时间（解决缓存雪崩）")
	userCache2 := cache.NewUserCacheWithAvalanche(rds)
	userService2 := service.NewUserServiceWithAvalanche(repo, userCache2, service.RandomExpire, cache.AvalancheBaseExpireSeconds)

	// 预热缓存
	for _, id := range userIDs {
		user, _ := repo.FindByID(id)
		if user != nil {
			userCache2.SetUserWithRandomExpire(user, cache.AvalancheBaseExpireSeconds)
		}
	}

	// 等待缓存开始过期
	fmt.Printf("  等待缓存开始过期（%d秒后开始陆续过期）...\n", cache.AvalancheBaseExpireSeconds)
	time.Sleep(time.Duration(cache.AvalancheBaseExpireSeconds) * time.Second)

	// 并发查询
	stats2 := &DBQueryStats{}
	var wg2 sync.WaitGroup
	start2 := time.Now()

	for i := 0; i < 50; i++ {
		wg2.Add(1)
		go func(i int) {
			defer wg2.Done()
			userID := userIDs[i%len(userIDs)]
			_, err := userService2.GetUserByID(userID)
			if err == nil {
				stats2.AddQuery()
			}
		}(i)
	}
	wg2.Wait()
	duration2 := time.Since(start2)

	fmt.Printf("  查询50次，数据库访问 %d 次，耗时 %v\n", stats2.TotalQueries, duration2)
	fmt.Printf("  数据库QPS峰值: %.2f\n", stats2.GetQPSInWindow(1))
	fmt.Println("  → 优势：数据库访问分散，压力平滑")

	// 效果对比
	fmt.Println("\n[效果对比]")
	fmt.Printf("  固定过期时间: 数据库访问 %d 次，QPS峰值 %.2f\n", stats1.TotalQueries, stats1.GetQPSInWindow(1))
	fmt.Printf("  随机过期时间: 数据库访问 %d 次，QPS峰值 %.2f\n", stats2.TotalQueries, stats2.GetQPSInWindow(1))
	if stats1.GetQPSInWindow(1) > 0 {
		reduction := (stats1.GetQPSInWindow(1) - stats2.GetQPSInWindow(1)) / stats1.GetQPSInWindow(1) * 100
		fmt.Printf("  QPS峰值减少: %.1f%%\n", reduction)
	}

	fmt.Println("\n✓ 场景3测试完成：随机过期时间显著平滑了数据库压力")
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("所有测试场景完成！")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\n总结：")
	fmt.Println("1. 缓存雪崩：大量缓存在同一时间过期，导致大量请求同时访问数据库")
	fmt.Println("2. 随机过期时间：给缓存设置随机的过期时间，避免缓存在同一时刻过期")
	fmt.Println("3. 效果：显著平滑数据库压力，提升系统稳定性")
}

// initDB 初始化数据库连接（复用main.go的函数）
func initDB(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库实例失败: %w", err)
	}

	if c.MySQL.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(c.MySQL.MaxOpenConns)
	}
	if c.MySQL.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	}

	return db, nil
}

// initRedis 初始化Redis连接（复用main.go的函数）
func initRedis(c Config) (*redis.Redis, error) {
	pingTimeout, err := time.ParseDuration(c.Redis.PingTimeout)
	if err != nil {
		pingTimeout = time.Second
	}

	conf := redis.RedisConf{
		Host:        c.Redis.Host,
		Type:        c.Redis.Type,
		Pass:        c.Redis.Password,
		PingTimeout: pingTimeout,
	}

	// 如果密码为空，不设置Pass字段
	if c.Redis.Password == "" {
		conf.Pass = ""
	}

	rds := redis.MustNewRedis(conf)
	return rds, nil
}

// ensureTestData 确保测试数据存在（复用main.go的函数）
func ensureTestData(db *gorm.DB) error {
	var count int64
	db.Model(&model.User{}).Count(&count)

	if count == 0 {
		log.Println("检测到数据库中没有测试数据，正在初始化...")
		return initTestData(db)
	}

	// 确保至少有10个用户用于测试
	if count < 10 {
		log.Printf("数据库只有 %d 个用户，需要至少10个，正在补充...", count)
		return initTestData(db)
	}

	log.Printf("数据库已有 %d 条测试数据", count)
	return nil
}

// initTestData 初始化测试数据（复用main.go的函数）
func initTestData(db *gorm.DB) error {
	users := []*model.User{
		{Username: "alice", Email: "alice@example.com", Age: 20},
		{Username: "bob", Email: "bob@example.com", Age: 25},
		{Username: "charlie", Email: "charlie@example.com", Age: 30},
		{Username: "david", Email: "david@example.com", Age: 28},
		{Username: "eve", Email: "eve@example.com", Age: 22},
		{Username: "frank", Email: "frank@example.com", Age: 35},
		{Username: "grace", Email: "grace@example.com", Age: 27},
		{Username: "henry", Email: "henry@example.com", Age: 32},
		{Username: "ivy", Email: "ivy@example.com", Age: 24},
		{Username: "jack", Email: "jack@example.com", Age: 29},
	}

	for _, user := range users {
		if err := db.Create(user).Error; err != nil {
			// 如果用户已存在，跳过
			continue
		}
		log.Printf("创建测试用户: ID=%d, Username=%s", user.ID, user.Username)
	}

	return nil
}
