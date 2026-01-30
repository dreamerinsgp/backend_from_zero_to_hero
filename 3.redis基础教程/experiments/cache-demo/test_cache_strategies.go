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
	userCache := cache.NewUserCache(rds)

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("缓存更新策略测试")
	fmt.Println(strings.Repeat("=", 80))

	// 场景1：读多写少 - 更新缓存策略
	testScenario1(userRepo, userCache, "场景1：读多写少 - 更新缓存策略")

	// 场景2：写多读少 - 删除缓存策略
	testScenario2(userRepo, userCache, "场景2：写多读少 - 删除缓存策略")

	// 场景3：数据一致性要求高 - 删除缓存策略
	testScenario3(userRepo, userCache, "场景3：数据一致性要求高 - 删除缓存策略")
}

// testScenario1 场景1：读多写少 - 更新缓存策略
// 特点：读操作频繁，写操作较少
// 策略：更新缓存（减少数据库查询）
func testScenario1(repo model.UserRepo, userCache cache.UserCache, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：读多写少场景，使用更新缓存策略")
	fmt.Println("优势：下次查询直接命中缓存，减少数据库查询压力")
	fmt.Println()

	// 创建使用更新缓存策略的服务
	userService := service.NewUserServiceWithStrategy(repo, userCache, service.UpdateCache)

	userID := int64(1)

	// 步骤1：首次查询（缓存未命中，从数据库加载）
	fmt.Println("[步骤1] 首次查询用户（缓存未命中）")
	user1, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Username=%s, Email=%s\n", user1.ID, user1.Username, user1.Email)
	time.Sleep(500 * time.Millisecond)

	// 步骤2：再次查询（缓存命中）
	fmt.Println("\n[步骤2] 再次查询用户（缓存命中）")
	user2, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Username=%s, Email=%s\n", user2.ID, user2.Username, user2.Email)
	fmt.Println("  → 注意：这次查询应该命中缓存，不会查询数据库")
	time.Sleep(500 * time.Millisecond)

	// 步骤3：更新用户（使用更新缓存策略）
	fmt.Println("\n[步骤3] 更新用户信息（策略：更新缓存）")
	user1.Email = "newemail2@example.com"
	user1.Age = 30
	if err := userService.UpdateUser(user1); err != nil {
		log.Printf("更新失败: %v", err)
		return
	}
	fmt.Printf("✓ 更新成功: Email=%s, Age=%d\n", user1.Email, user1.Age)
	fmt.Println("  → 注意：缓存已更新，下次查询会直接返回最新数据")
	time.Sleep(500 * time.Millisecond)

	// 步骤4：查询更新后的数据（应该从缓存获取最新数据）
	fmt.Println("\n[步骤4] 查询更新后的用户（缓存命中，返回最新数据）")
	user3, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Email=%s, Age=%d\n", user3.ID, user3.Email, user3.Age)
	fmt.Println("  → 注意：这次查询命中缓存，直接返回更新后的数据，无需查询数据库")

	fmt.Println("\n✓ 场景1测试完成：更新缓存策略适合读多写少场景")
}

// testScenario2 场景2：写多读少 - 删除缓存策略
// 特点：写操作频繁，读操作较少
// 策略：删除缓存（避免频繁更新缓存）
func testScenario2(repo model.UserRepo, userCache cache.UserCache, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：写多读少场景，使用删除缓存策略")
	fmt.Println("优势：避免频繁更新缓存，节省缓存更新成本")
	fmt.Println()

	// 创建使用删除缓存策略的服务
	userService := service.NewUserServiceWithStrategy(repo, userCache, service.DeleteCache)

	userID := int64(2)

	// 步骤1：首次查询（缓存未命中，从数据库加载）
	fmt.Println("[步骤1] 首次查询用户（缓存未命中）")
	user1, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Username=%s, Email=%s\n", user1.ID, user1.Username, user1.Email)
	time.Sleep(500 * time.Millisecond)

	// 步骤2：再次查询（缓存命中）
	fmt.Println("\n[步骤2] 再次查询用户（缓存命中）")
	user2, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Username=%s\n", user2.ID, user2.Username)
	time.Sleep(500 * time.Millisecond)

	// 步骤3：更新用户（使用删除缓存策略）
	fmt.Println("\n[步骤3] 更新用户信息（策略：删除缓存）")
	user1.Email = "newemail2@example.com"
	user1.Age = 30
	if err := userService.UpdateUser(user1); err != nil {
		log.Printf("更新失败: %v", err)
		return
	}
	fmt.Printf("✓ 更新成功: Email=%s, Age=%d\n", user1.Email, user1.Age)
	fmt.Println("  → 注意：缓存已删除，下次查询会重新从数据库加载")
	time.Sleep(500 * time.Millisecond)

	// 步骤4：查询更新后的数据（缓存未命中，从数据库重新加载）
	fmt.Println("\n[步骤4] 查询更新后的用户（缓存未命中，从数据库重新加载）")
	user3, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Email=%s, Age=%d\n", user3.ID, user3.Email, user3.Age)
	fmt.Println("  → 注意：这次查询缓存未命中，从数据库重新加载最新数据")

	fmt.Println("\n✓ 场景2测试完成：删除缓存策略适合写多读少场景")
}

// testScenario3 场景3：数据一致性要求高 - 删除缓存策略
// 特点：对数据一致性要求极高，不能容忍缓存和数据库不一致
// 策略：删除缓存（保证下次查询时获取最新数据）
func testScenario3(repo model.UserRepo, userCache cache.UserCache, title string) {
	fmt.Println("\n" + strings.Repeat("-", 80))
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("说明：数据一致性要求高场景，使用删除缓存策略")
	fmt.Println("优势：保证数据一致性，下次查询时获取数据库最新数据")
	fmt.Println()

	// 创建使用删除缓存策略的服务
	userService := service.NewUserServiceWithStrategy(repo, userCache, service.DeleteCache)

	userID := int64(3)

	// 步骤1：首次查询（缓存未命中，从数据库加载）
	fmt.Println("[步骤1] 首次查询用户（缓存未命中）")
	user1, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Username=%s, Email=%s\n", user1.ID, user1.Username, user1.Email)
	time.Sleep(500 * time.Millisecond)

	// 步骤2：再次查询（缓存命中）
	fmt.Println("\n[步骤2] 再次查询用户（缓存命中）")
	user2, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Username=%s\n", user2.ID, user2.Username)
	time.Sleep(500 * time.Millisecond)

	// 步骤3：模拟外部系统直接更新数据库（绕过应用层）
	fmt.Println("\n[步骤3] 模拟外部系统直接更新数据库（绕过应用层）")
	// 直接更新数据库，不通过服务层
	user1.Email = "external_update@example.com"
	user1.Age = 35
	if err := repo.Update(user1); err != nil {
		log.Printf("直接更新数据库失败: %v", err)
		return
	}
	fmt.Printf("✓ 数据库已更新: Email=%s, Age=%d\n", user1.Email, user1.Age)
	fmt.Println("  → 注意：数据库已更新，但缓存中仍是旧数据")
	time.Sleep(500 * time.Millisecond)

	// 步骤4：通过服务层更新（删除缓存策略）
	fmt.Println("\n[步骤4] 通过服务层更新用户（策略：删除缓存）")
	user1.Email = "service_update@example.com"
	user1.Age = 40
	if err := userService.UpdateUser(user1); err != nil {
		log.Printf("更新失败: %v", err)
		return
	}
	fmt.Printf("✓ 更新成功: Email=%s, Age=%d\n", user1.Email, user1.Age)
	fmt.Println("  → 注意：缓存已删除，保证下次查询时获取最新数据")
	time.Sleep(500 * time.Millisecond)

	// 步骤5：查询更新后的数据（缓存未命中，从数据库重新加载）
	fmt.Println("\n[步骤5] 查询更新后的用户（缓存未命中，从数据库重新加载）")
	user3, err := userService.GetUserByID(userID)
	if err != nil {
		log.Printf("查询失败: %v", err)
		return
	}
	fmt.Printf("✓ 查询成功: ID=%d, Email=%s, Age=%d\n", user3.ID, user3.Email, user3.Age)
	fmt.Println("  → 注意：这次查询缓存未命中，从数据库重新加载，保证数据一致性")

	fmt.Println("\n✓ 场景3测试完成：删除缓存策略适合数据一致性要求高的场景")
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("所有测试场景完成！")
	fmt.Println(strings.Repeat("=", 80))
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
