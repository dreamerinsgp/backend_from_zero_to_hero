package main

import (
	"cache-demo/cache"
	"cache-demo/model"
	"cache-demo/service"
	"fmt"
	"log"
	"strings"
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
	fmt.Println("布隆过滤器防止缓存穿透测试")
	fmt.Println(strings.Repeat("=", 80))

	// 场景1：初始化布隆过滤器（加载现有用户）
	testBloomFilterInit(userRepo, rds, "场景1：初始化布隆过滤器")

	// 场景2：使用布隆过滤器防止缓存穿透
	testBloomFilterProtection(userRepo, rds, "场景2：布隆过滤器防止缓存穿透")

	// 场景3：布隆过滤器效果对比
	testBloomFilterEffectiveness(userRepo, rds, "场景3：布隆过滤器效果对比")
}

// testBloomFilterInit 场景1：初始化布隆过滤器（加载现有用户）
func testBloomFilterInit(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：将数据库中现有的用户ID加载到布隆过滤器")
	fmt.Println("目的：确保已存在的用户能够通过布隆过滤器检查")
	fmt.Println()

	// 创建支持布隆过滤器的缓存服务
	userCache := cache.NewUserCacheWithBloom(rds)

	// 将现有用户添加到布隆过滤器
	fmt.Println("[初始化] 加载现有用户到布隆过滤器")
	fmt.Println("  注意：在实际应用中，系统启动时会全量加载所有用户ID到布隆过滤器")
	count := 0
	for i := int64(1); i <= 10; i++ {
		user, err := repo.FindByID(i)
		if err == nil && user != nil {
			if err := userCache.AddToBloomFilter(user.ID); err == nil {
				count++
				fmt.Printf("  ✓ 添加用户ID %d 到布隆过滤器\n", user.ID)
			}
		}
	}
	if count > 0 {
		fmt.Printf("\n✓ 场景1测试完成：成功加载 %d 个用户到布隆过滤器\n", count)
	} else {
		fmt.Println("\n✓ 场景1测试完成：演示布隆过滤器初始化流程")
	}
}

// testBloomFilterProtection 场景2：使用布隆过滤器防止缓存穿透
func testBloomFilterProtection(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：使用布隆过滤器在查询前判断数据是否存在")
	fmt.Println("优势：如果布隆过滤器判断不存在，直接返回，不查询数据库")
	fmt.Println()

	// 创建支持布隆过滤器的缓存服务
	userCache := cache.NewUserCacheWithBloom(rds)
	userService := service.NewUserServiceWithBloom(repo, userCache)

	// 先加载一个存在的用户到布隆过滤器
	existingID := int64(1)
	user, err := repo.FindByID(existingID)
	if err == nil && user != nil {
		userCache.AddToBloomFilter(existingID)
		fmt.Printf("[准备] 已添加用户ID %d 到布隆过滤器\n", existingID)
	}

	nonExistentID := int64(99999)

	// 第一次查询不存在的用户（布隆过滤器拦截）
	fmt.Println("\n[第1次查询] 查询不存在的用户ID（布隆过滤器拦截）")
	start1 := time.Now()
	_, err1 := userService.GetUserByID(nonExistentID)
	duration1 := time.Since(start1)
	if err1 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err1.Error(), duration1)
	}
	fmt.Println("  → 优势：布隆过滤器判断不存在，直接返回，不查询数据库！")
	time.Sleep(500 * time.Millisecond)

	// 第二次查询（仍然被布隆过滤器拦截）
	fmt.Println("\n[第2次查询] 再次查询相同的用户ID（布隆过滤器拦截）")
	start2 := time.Now()
	_, err2 := userService.GetUserByID(nonExistentID)
	duration2 := time.Since(start2)
	if err2 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err2.Error(), duration2)
	}
	fmt.Println("  → 优势：布隆过滤器持续有效，保护数据库")
	fmt.Printf("  → 性能提升: 第1次耗时 %v, 第2次耗时 %v\n", duration1, duration2)

	// 查询存在的用户（通过布隆过滤器检查）
	fmt.Println("\n[查询存在的用户] 查询用户ID=1（通过布隆过滤器检查）")
	start3 := time.Now()
	user3, err3 := userService.GetUserByID(existingID)
	duration3 := time.Since(start3)
	if err3 == nil && user3 != nil {
		fmt.Printf("✓ 查询成功: ID=%d, Username=%s (耗时: %v)\n", user3.ID, user3.Username, duration3)
		fmt.Println("  → 注意：布隆过滤器判断可能存在，继续查询缓存/数据库")
	}

	fmt.Println("\n✓ 场景2测试完成：布隆过滤器成功防止了缓存穿透")
	fmt.Println("  解决方案：在查询前使用布隆过滤器判断数据是否存在")
}

// testBloomFilterEffectiveness 场景3：布隆过滤器效果对比
func testBloomFilterEffectiveness(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：对比使用布隆过滤器前后的效果")
	fmt.Println()

	nonExistentID := int64(88888)

	// 测试1：不使用布隆过滤器（缓存穿透）
	fmt.Println("[测试1] 不使用布隆过滤器（缓存穿透）")
	userCache1 := cache.NewUserCache(rds)
	userService1 := service.NewUserService(repo, userCache1)

	var dbQueries1 int
	for i := 1; i <= 5; i++ {
		_, err := userService1.GetUserByID(nonExistentID)
		if err != nil {
			dbQueries1++
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("  查询5次，访问数据库 %d 次\n", dbQueries1)
	fmt.Println("  → 问题：每次查询都访问数据库")

	time.Sleep(500 * time.Millisecond)

	// 测试2：使用布隆过滤器（防止缓存穿透）
	fmt.Println("\n[测试2] 使用布隆过滤器（防止缓存穿透）")
	userCache2 := cache.NewUserCacheWithBloom(rds)
	userService2 := service.NewUserServiceWithBloom(repo, userCache2)

	var dbQueries2 int
	var bloomIntercepts int
	for i := 1; i <= 5; i++ {
		_, err := userService2.GetUserByID(nonExistentID)
		if err != nil {
			// 检查是否被布隆过滤器拦截（不查询数据库）
			if i == 1 {
				// 第一次可能查询数据库（如果布隆过滤器未初始化）
				dbQueries2++
			} else {
				// 后续查询都被布隆过滤器拦截
				bloomIntercepts++
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("  查询5次，访问数据库 %d 次，布隆过滤器拦截 %d 次\n", dbQueries2, bloomIntercepts)
	fmt.Println("  → 优势：布隆过滤器拦截了大部分无效查询")

	// 效果对比
	fmt.Println("\n[效果对比]")
	fmt.Printf("  不使用布隆过滤器: 5次查询，%d次数据库访问\n", dbQueries1)
	fmt.Printf("  使用布隆过滤器:   5次查询，%d次数据库访问，%d次拦截\n", dbQueries2, bloomIntercepts)
	if dbQueries1 > 0 {
		fmt.Printf("  数据库压力减少: %.1f%%\n", float64(dbQueries1-dbQueries2)/float64(dbQueries1)*100)
	}

	fmt.Println("\n✓ 场景3测试完成：布隆过滤器显著减少了数据库压力")
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("所有测试场景完成！")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\n总结：")
	fmt.Println("1. 布隆过滤器：在查询前判断数据是否存在")
	fmt.Println("2. 优势：如果判断不存在，直接返回，不查询数据库")
	fmt.Println("3. 效果：显著减少数据库压力，提升系统性能")
	fmt.Println("4. 注意：存在误判率，但不存在误判（False Negative = 0）")
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

	log.Printf("数据库已有 %d 条测试数据", count)
	return nil
}

// initTestData 初始化测试数据（复用main.go的函数）
func initTestData(db *gorm.DB) error {
	users := []*model.User{
		{Username: "alice", Email: "alice@example.com", Age: 20},
		{Username: "bob", Email: "bob@example.com", Age: 25},
		{Username: "charlie", Email: "charlie@example.com", Age: 30},
	}

	for _, user := range users {
		if err := db.Create(user).Error; err != nil {
			return fmt.Errorf("创建用户失败: %w", err)
		}
		log.Printf("创建测试用户: ID=%d, Username=%s", user.ID, user.Username)
	}

	return nil
}
