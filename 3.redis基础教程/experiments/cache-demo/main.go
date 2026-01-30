package main

import (
	"cache-demo/cache"
	"cache-demo/model"
	"cache-demo/service"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config 配置结构
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
	// 检查命令参数
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "reset":
			resetCache()
			return
		case "reset-db":
			resetDatabase()
			return
		case "init-db":
			initDatabase()
			return
		case "help":
			showHelp()
			return
		}
	}

	// 1. 加载配置
	var c Config
	conf.MustLoad("config.yaml", &c)

	// 打印配置信息（用于调试）
	log.Printf("MySQL配置: Host=%s, Port=%d, User=%s, Database=%s",
		c.MySQL.Host, c.MySQL.Port, c.MySQL.User, c.MySQL.Database)
	log.Printf("Redis配置: Host=%s", c.Redis.Host)

	// 2. 初始化数据库连接
	db, err := initDB(c)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	log.Println("数据库连接成功")

	// 3. 初始化Redis连接
	rds, err := initRedis(c)
	if err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}
	log.Println("Redis连接成功")

	// 4. 检查并初始化测试数据
	if err := ensureTestData(db); err != nil {
		log.Fatalf("初始化测试数据失败: %v", err)
	}

	// 5. 初始化服务层
	userRepo := model.NewUserRepo(db)
	userCache := cache.NewUserCache(rds)
	userService := service.NewUserService(userRepo, userCache)

	// 6. 演示缓存的基本使用
	demonstrateCache(userService)
}

// ensureTestData 确保测试数据存在
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

// initTestData 初始化测试数据
func initTestData(db *gorm.DB) error {
	userRepo := model.NewUserRepo(db)

	// 检查是否已有数据
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		log.Printf("数据库已有 %d 条数据，跳过插入", count)
		return nil
	}

	testUsers := []*model.User{
		{Username: "alice", Email: "alice@example.com", Age: 25},
		{Username: "bob", Email: "bob@example.com", Age: 30},
		{Username: "charlie", Email: "charlie@example.com", Age: 28},
	}

	for _, user := range testUsers {
		if err := userRepo.Create(user); err != nil {
			return fmt.Errorf("创建用户 %s 失败: %w", user.Username, err)
		}
		log.Printf("✓ 创建测试用户: %s (ID: %d)", user.Username, user.ID)
	}

	log.Println("✓ 测试数据初始化完成")
	return nil
}

// initDB 初始化数据库连接
func initDB(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(c.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// initRedis 初始化Redis连接
func initRedis(c Config) (*redis.Redis, error) {
	pingTimeout, err := time.ParseDuration(c.Redis.PingTimeout)
	if err != nil {
		pingTimeout = 10 * time.Second // 默认10秒
	}

	// 构建 Redis 配置
	redisConf := redis.RedisConf{
		Host:        c.Redis.Host,
		Type:        c.Redis.Type,
		PingTimeout: pingTimeout,
	}

	// 如果配置了密码，则设置密码（空字符串表示无密码）
	if c.Redis.Password != "" {
		redisConf.Pass = c.Redis.Password
	}

	rds := redis.MustNewRedis(redisConf)

	if !rds.Ping() {
		return nil, fmt.Errorf("Redis ping失败")
	}

	return rds, nil
}

// demonstrateCache 演示缓存的基本使用
func demonstrateCache(userService service.UserService) {
	fmt.Println("\n========== 缓存基本用法演示 ==========")

	// 场景1: 第一次查询（缓存未命中，查数据库）
	fmt.Println("\n【场景1】第一次查询用户ID=1（缓存未命中）")
	user1, err := userService.GetUserByID(1)
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("查询结果: ID=%d, Username=%s, Email=%s, Age=%d\n",
			user1.ID, user1.Username, user1.Email, user1.Age)
	}

	// 等待1秒，便于观察日志
	time.Sleep(1 * time.Second)

	// 场景2: 第二次查询（缓存命中）
	fmt.Println("\n【场景2】第二次查询用户ID=1（缓存命中）")
	user2, err := userService.GetUserByID(1)
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("查询结果: ID=%d, Username=%s, Email=%s, Age=%d\n",
			user2.ID, user2.Username, user2.Email, user2.Age)
	}

	// 等待1秒
	time.Sleep(1 * time.Second)

	// 场景3: 查询另一个用户（缓存未命中）
	fmt.Println("\n【场景3】查询用户ID=2（缓存未命中）")
	user3, err := userService.GetUserByID(2)
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("查询结果: ID=%d, Username=%s, Email=%s, Age=%d\n",
			user3.ID, user3.Username, user3.Email, user3.Age)
	}

	// 等待1秒
	time.Sleep(1 * time.Second)

	// 场景4: 再次查询用户ID=2（缓存命中）
	fmt.Println("\n【场景4】再次查询用户ID=2（缓存命中）")
	user4, err := userService.GetUserByID(2)
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("查询结果: ID=%d, Username=%s, Email=%s, Age=%d\n",
			user4.ID, user4.Username, user4.Email, user4.Age)
	}

	// 等待1秒
	time.Sleep(1 * time.Second)

	// 场景5: 更新用户（更新缓存）
	fmt.Println("\n【场景5】更新用户ID=1的信息（更新缓存）")
	user1.Age = 26
	if err := userService.UpdateUser(user1); err != nil {
		log.Printf("更新失败: %v", err)
	} else {
		fmt.Printf("更新成功: ID=%d, Age=%d\n", user1.ID, user1.Age)
	}

	// 等待1秒
	time.Sleep(1 * time.Second)

	// 场景6: 查询更新后的用户（应该从缓存获取最新数据）
	fmt.Println("\n【场景6】查询更新后的用户ID=1（从缓存获取最新数据）")
	user5, err := userService.GetUserByID(1)
	if err != nil {
		log.Printf("查询失败: %v", err)
	} else {
		fmt.Printf("查询结果: ID=%d, Username=%s, Email=%s, Age=%d\n",
			user5.ID, user5.Username, user5.Email, user5.Age)
	}

	// 场景7: 查询不存在的用户（演示后续缓存穿透场景的基础）
	fmt.Println("\n【场景7】查询不存在的用户ID=99999（演示缓存穿透场景）")
	_, err = userService.GetUserByID(99999)
	if err != nil {
		fmt.Printf("查询结果: %v\n", err)
	}

	fmt.Println("\n========== 演示完成 ==========")
	fmt.Println("\n提示:")
	fmt.Println("1. 观察日志输出，可以看到缓存命中/未命中的情况")
	fmt.Println("2. 可以使用 redis-cli 查看缓存数据: redis-cli")
	fmt.Println("3. 查看缓存Key: KEYS user:*")
	fmt.Println("4. 查看具体缓存: GET user:1")
	fmt.Println("5. 查看TTL: TTL user:1")
}

// resetCache 只重置缓存（不重置数据库）
func resetCache() {
	// 1. 加载配置
	var c Config
	conf.MustLoad("config.yaml", &c)

	fmt.Println("========== 重置缓存 ==========")

	// 2. 初始化Redis连接
	rds, err := initRedis(c)
	if err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}

	// 3. 清理 Redis 缓存
	fmt.Println("\n清理 Redis 缓存...")
	keys, err := rds.Keys("user:*")
	if err != nil {
		log.Printf("获取缓存Key失败: %v", err)
	} else {
		if len(keys) > 0 {
			for _, key := range keys {
				rds.Del(key)
			}
			fmt.Printf("✓ 已清理 %d 个缓存Key\n", len(keys))
		} else {
			fmt.Println("✓ 缓存已为空，无需清理")
		}
	}

	// 验证缓存是否已清理
	remainingKeys, _ := rds.Keys("user:*")
	fmt.Printf("剩余缓存数量: %d\n", len(remainingKeys))

	fmt.Println("\n========== 缓存重置完成 ==========")
	fmt.Println("\n现在可以运行程序进行新的实验：")
	fmt.Println("  go run main.go")
	fmt.Println()
}

// resetDatabase 重置数据库（删除并重新创建测试数据）
func resetDatabase() {
	// 1. 加载配置
	var c Config
	conf.MustLoad("config.yaml", &c)

	fmt.Println("========== 重置数据库 ==========")

	// 2. 初始化数据库连接
	db, err := initDB(c)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 3. 删除所有用户数据并重置自增ID
	fmt.Println("\n删除所有用户数据并重置自增ID...")
	if err := db.Exec("TRUNCATE TABLE users").Error; err != nil {
		log.Fatalf("重置用户表失败: %v", err)
	}
	fmt.Println("✓ 已删除所有用户数据并重置自增ID")

	// 4. 重新插入测试数据
	fmt.Println("\n重新插入测试数据...")
	if err := initTestData(db); err != nil {
		log.Fatalf("初始化测试数据失败: %v", err)
	}

	// 验证数据
	var count int64
	db.Model(&model.User{}).Count(&count)
	fmt.Printf("\n✓ 数据库数据已重置\n")
	fmt.Printf("  用户数量: %d\n", count)

	fmt.Println("\n========== 数据库重置完成 ==========")
	fmt.Println("\n现在可以运行程序进行新的实验：")
	fmt.Println("  go run main.go")
	fmt.Println()
}

// initDatabase 初始化数据库（创建表并插入测试数据）
func initDatabase() {
	// 1. 加载配置
	var c Config
	conf.MustLoad("config.yaml", &c)

	fmt.Println("========== 初始化数据库 ==========")

	// 2. 初始化数据库连接
	db, err := initDB(c)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 3. 自动迁移（创建表）
	fmt.Println("\n创建数据表...")
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("创建表失败: %v", err)
	}
	fmt.Println("✓ 数据表已创建")

	// 4. 检查并插入测试数据
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		fmt.Println("\n插入测试数据...")
		if err := initTestData(db); err != nil {
			log.Fatalf("初始化测试数据失败: %v", err)
		}
	} else {
		fmt.Printf("\n数据库已有 %d 条数据，跳过插入\n", count)
	}

	fmt.Println("\n========== 数据库初始化完成 ==========")
	fmt.Println("\n现在可以运行程序：")
	fmt.Println("  go run main.go")
	fmt.Println()
}

// showHelp 显示帮助信息
func showHelp() {
	fmt.Println("缓存演示程序 - 使用说明")
	fmt.Println()
	fmt.Println("命令:")
	fmt.Println("  go run main.go         运行缓存演示程序")
	fmt.Println("  go run main.go reset   重置缓存（清理所有 user:* 缓存）")
	fmt.Println("  go run main.go reset-db 重置数据库（删除并重新插入测试数据）")
	fmt.Println("  go run main.go init-db  初始化数据库（创建表并插入测试数据）")
	fmt.Println("  go run main.go help     显示此帮助信息")
	fmt.Println()
	fmt.Println("说明:")
	fmt.Println("  - reset: 只清理 Redis 缓存，不影响数据库")
	fmt.Println("  - reset-db: 重置数据库数据，不影响缓存")
	fmt.Println("  - init-db: 创建表并插入测试数据（如果表已存在则跳过）")
	fmt.Println()
}
