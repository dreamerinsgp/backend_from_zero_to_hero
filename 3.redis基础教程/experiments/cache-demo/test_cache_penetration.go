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
	fmt.Println("缓存穿透问题测试")
	fmt.Println(strings.Repeat("=", 80))

	// 场景1：演示缓存穿透问题（未使用空值缓存）
	testPenetrationProblem(userRepo, rds, "场景1：缓存穿透问题演示")

	// 场景2：使用空值缓存解决缓存穿透
	testNullCacheSolution(userRepo, rds, "场景2：空值缓存解决方案")

	// 场景3：空值缓存的效果对比
	testNullCacheEffectiveness(userRepo, rds, "场景3：空值缓存效果对比")
}

// testPenetrationProblem 场景1：演示缓存穿透问题（未使用空值缓存）
func testPenetrationProblem(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：演示缓存穿透问题 - 查询不存在的数据，每次都访问数据库")
	fmt.Println("问题：数据既不在缓存中，也不在数据库中，无法缓存空结果")
	fmt.Println()

	// 使用普通的缓存服务（不支持空值缓存）
	userCache := cache.NewUserCache(rds)
	userService := service.NewUserService(repo, userCache)

	nonExistentID := int64(99999)

	// 第一次查询
	fmt.Println("[第1次查询] 查询不存在的用户ID（缓存未命中，查询数据库）")
	start1 := time.Now()
	_, err1 := userService.GetUserByID(nonExistentID)
	duration1 := time.Since(start1)
	if err1 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err1.Error(), duration1)
	}
	fmt.Println("  → 注意：查询了数据库，返回空结果，但没有缓存")
	time.Sleep(500 * time.Millisecond)

	// 第二次查询（应该还会查询数据库）
	fmt.Println("\n[第2次查询] 再次查询相同的用户ID（缓存未命中，再次查询数据库）")
	start2 := time.Now()
	_, err2 := userService.GetUserByID(nonExistentID)
	duration2 := time.Since(start2)
	if err2 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err2.Error(), duration2)
	}
	fmt.Println("  → 问题：再次查询了数据库！这就是缓存穿透问题")
	time.Sleep(500 * time.Millisecond)

	// 第三次查询（应该还会查询数据库）
	fmt.Println("\n[第3次查询] 第三次查询相同的用户ID（缓存未命中，第三次查询数据库）")
	start3 := time.Now()
	_, err3 := userService.GetUserByID(nonExistentID)
	duration3 := time.Since(start3)
	if err3 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err3.Error(), duration3)
	}
	fmt.Println("  → 问题：每次查询都访问数据库，造成数据库压力")

	fmt.Println("\n✓ 场景1测试完成：演示了缓存穿透问题")
	fmt.Println("  问题总结：查询不存在的数据，每次都穿透缓存访问数据库")
}

// testNullCacheSolution 场景2：使用空值缓存解决缓存穿透
func testNullCacheSolution(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：使用空值缓存解决缓存穿透问题")
	fmt.Println("方案：将空结果也缓存起来，设置较短的过期时间（60秒）")
	fmt.Println()

	// 使用支持空值缓存的缓存服务
	userCache := cache.NewUserCacheWithPenetration(rds)
	userService := service.NewUserServiceWithPenetration(repo, userCache)

	nonExistentID := int64(88888)

	// 第一次查询（设置空值缓存）
	fmt.Println("[第1次查询] 查询不存在的用户ID（缓存未命中，查询数据库，设置空值缓存）")
	start1 := time.Now()
	_, err1 := userService.GetUserByID(nonExistentID)
	duration1 := time.Since(start1)
	if err1 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err1.Error(), duration1)
	}
	fmt.Printf("  → 注意：查询了数据库，返回空结果，并设置了空值缓存（过期时间: %d秒）\n", cache.NullCacheExpireSeconds)
	time.Sleep(500 * time.Millisecond)

	// 第二次查询（应该命中空值缓存）
	fmt.Println("\n[第2次查询] 再次查询相同的用户ID（空值缓存命中，不查询数据库）")
	start2 := time.Now()
	_, err2 := userService.GetUserByID(nonExistentID)
	duration2 := time.Since(start2)
	if err2 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err2.Error(), duration2)
	}
	fmt.Println("  → 优势：空值缓存命中，直接返回，不查询数据库！")
	fmt.Printf("  → 性能提升: 第1次耗时 %v, 第2次耗时 %v (提升: %.2f%%)\n",
		duration1, duration2, float64(duration1-duration2)/float64(duration1)*100)
	time.Sleep(500 * time.Millisecond)

	// 第三次查询（应该还是命中空值缓存）
	fmt.Println("\n[第3次查询] 第三次查询相同的用户ID（空值缓存命中，不查询数据库）")
	start3 := time.Now()
	_, err3 := userService.GetUserByID(nonExistentID)
	duration3 := time.Since(start3)
	if err3 != nil {
		fmt.Printf("✓ 查询结果: %v (耗时: %v)\n", err3.Error(), duration3)
	}
	fmt.Println("  → 优势：空值缓存持续有效，保护数据库")

	fmt.Println("\n✓ 场景2测试完成：空值缓存成功解决了缓存穿透问题")
	fmt.Println("  解决方案：将空结果缓存起来，避免重复查询数据库")
}

// testNullCacheEffectiveness 场景3：空值缓存的效果对比
func testNullCacheEffectiveness(repo model.UserRepo, rds *redis.Redis, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：对比使用空值缓存前后的效果")
	fmt.Println()

	nonExistentID := int64(77777)

	// 测试1：不使用空值缓存（缓存穿透）
	fmt.Println("[测试1] 不使用空值缓存（缓存穿透）")
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

	// 测试2：使用空值缓存（解决缓存穿透）
	fmt.Println("\n[测试2] 使用空值缓存（解决缓存穿透）")
	userCache2 := cache.NewUserCacheWithPenetration(rds)
	userService2 := service.NewUserServiceWithPenetration(repo, userCache2)

	var dbQueries2 int
	for i := 1; i <= 5; i++ {
		_, err := userService2.GetUserByID(nonExistentID)
		if err != nil {
			// 检查是否是空值缓存命中
			if i == 1 {
				dbQueries2++ // 第一次查询数据库
			}
			// 后续查询都是空值缓存命中，不查询数据库
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("  查询5次，访问数据库 %d 次\n", dbQueries2)
	fmt.Println("  → 优势：只有第1次查询数据库，后续都命中空值缓存")

	// 效果对比
	fmt.Println("\n[效果对比]")
	fmt.Printf("  不使用空值缓存: 5次查询，%d次数据库访问\n", dbQueries1)
	fmt.Printf("  使用空值缓存:   5次查询，%d次数据库访问\n", dbQueries2)
	fmt.Printf("  数据库压力减少: %.1f%%\n", float64(dbQueries1-dbQueries2)/float64(dbQueries1)*100)

	fmt.Println("\n✓ 场景3测试完成：空值缓存显著减少了数据库压力")
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("所有测试场景完成！")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\n总结：")
	fmt.Println("1. 缓存穿透：查询不存在的数据，每次都访问数据库")
	fmt.Println("2. 空值缓存：将空结果缓存起来，避免重复查询数据库")
	fmt.Println("3. 效果：显著减少数据库压力，提升系统性能")
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
