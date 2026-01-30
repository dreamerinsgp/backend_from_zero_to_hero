package main

import (
	"cache-demo/model"
	"fmt"
	"log"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
)

// resetExperiment 重置实验环境
func resetExperiment() {
	// 1. 加载配置
	var c Config
	conf.MustLoad("config.yaml", &c)

	fmt.Println("========== 重置实验环境 ==========")

	// 2. 初始化数据库连接
	db, err := initDB(c)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 3. 初始化Redis连接
	rds, err := initRedis(c)
	if err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}

	// 4. 清理 Redis 缓存
	fmt.Println("\n【步骤1】清理 Redis 缓存...")
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
	fmt.Printf("  剩余缓存数量: %d\n", len(remainingKeys))

	// 5. 重置数据库数据
	fmt.Println("\n【步骤2】重置数据库数据...")
	userRepo := model.NewUserRepo(db)

	// 删除所有用户
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		log.Fatalf("删除用户数据失败: %v", err)
	}

	// 重新插入测试数据
	testUsers := []*model.User{
		{Username: "alice", Email: "alice@example.com", Age: 25},
		{Username: "bob", Email: "bob@example.com", Age: 30},
		{Username: "charlie", Email: "charlie@example.com", Age: 28},
	}

	for _, user := range testUsers {
		if err := userRepo.Create(user); err != nil {
			log.Printf("创建用户失败: %v", err)
		}
	}

	// 验证数据
	var count int64
	db.Model(&model.User{}).Count(&count)
	fmt.Printf("✓ 数据库数据已重置\n")
	fmt.Printf("  用户数量: %d\n", count)

	fmt.Println("\n========== 重置完成 ==========")
	fmt.Println("\n现在可以运行程序进行新的实验：")
	fmt.Println("  go run main.go")
	fmt.Println()
}

func main() {
	// 检查是否是重置命令
	if len(os.Args) > 1 && os.Args[1] == "reset" {
		resetExperiment()
		return
	}

	// 否则显示帮助信息
	fmt.Println("使用方法:")
	fmt.Println("  重置实验环境: go run reset.go reset")
	fmt.Println("  或者使用脚本: bash reset.sh")
	fmt.Println()
	fmt.Println("重置功能:")
	fmt.Println("  1. 清理 Redis 中的所有用户缓存 (user:*)")
	fmt.Println("  2. 重置数据库数据（删除并重新插入测试数据）")
}
