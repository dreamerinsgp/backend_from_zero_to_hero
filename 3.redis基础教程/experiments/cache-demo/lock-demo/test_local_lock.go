package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// Config 配置结构
type LockConfig struct {
	Redis struct {
		Host        string `json:"host" yaml:"host"`
		Password    string `json:"password" yaml:"password"`
		Type        string `json:"type" yaml:"type"`
		PingTimeout string `json:"ping_timeout" yaml:"ping_timeout"`
	} `json:"redis" yaml:"redis"`
}

// ❌ 错误：使用普通锁（mutex）- 只能控制单进程内的线程
var mu sync.Mutex

func main() {
	// 获取进程ID
	processID := os.Getpid()

	// 获取命令行参数（进程标识）
	processName := fmt.Sprintf("进程-%d", processID)
	if len(os.Args) > 1 {
		processName = os.Args[1]
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("【实验1】多进程普通锁（mutex）- 展示问题\n")
	fmt.Printf("%s 启动 (PID: %d)\n", processName, processID)
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("⚠️  注意：普通锁（mutex）只能控制单进程内的线程，无法跨进程控制！")
	fmt.Println(strings.Repeat("=", 80))

	// 加载配置
	var c LockConfig
	err := conf.Load("../config.yaml", &c)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		return
	}

	// 初始化Redis（仅用于存储库存数据）
	pingTimeout, _ := time.ParseDuration(c.Redis.PingTimeout)
	if pingTimeout == 0 {
		pingTimeout = 10 * time.Second
	}

	redisConf := redis.RedisConf{
		Host:        c.Redis.Host,
		Pass:        c.Redis.Password,
		Type:        c.Redis.Type,
		PingTimeout: pingTimeout,
	}

	redisClient, err := redis.NewRedis(redisConf)
	if err != nil {
		fmt.Printf("连接Redis失败: %v\n", err)
		return
	}

	// 场景：库存扣减（模拟超卖问题）
	fmt.Printf("\n【场景】库存扣减 - 使用普通锁（mutex）\n")
	fmt.Println(strings.Repeat("-", 80))

	productID := "product:1001"

	// 初始化库存（如果不存在）
	stockKey := fmt.Sprintf("stock:%s", productID)
	exists, _ := redisClient.Exists(stockKey)
	if !exists {
		redisClient.Set(stockKey, "100") // 初始库存100
		fmt.Printf("%s: 初始化库存: %s = 100\n", processName, productID)
	}

	// 获取当前库存
	currentStock, _ := redisClient.Get(stockKey)
	stockValue, _ := strconv.Atoi(currentStock)
	fmt.Printf("%s: 📦 当前库存: %d\n", processName, stockValue)

	// 尝试扣减库存（购买数量）
	purchaseQuantity := 10
	fmt.Printf("%s: 🛒 尝试购买 %d 件商品...\n", processName, purchaseQuantity)

	// ❌ 使用普通锁（mutex）- 只能锁住当前进程内的线程
	fmt.Printf("%s: 🔒 获取本地锁（mutex）...\n", processName)
	mu.Lock() // 只能锁住当前进程内的线程，无法锁住其他进程！
	fmt.Printf("%s: ✅ 获取本地锁成功（但其他进程的锁是独立的！）\n", processName)

	// 模拟一些处理时间
	time.Sleep(500 * time.Millisecond)

	// 重新读取库存
	currentStock, _ = redisClient.Get(stockKey)
	stockValue, _ = strconv.Atoi(currentStock)
	fmt.Printf("%s: 📖 重新读取库存: %d\n", processName, stockValue)

	if stockValue < purchaseQuantity {
		fmt.Printf("%s: ❌ 库存不足（当前: %d, 需要: %d）\n", processName, stockValue, purchaseQuantity)
		mu.Unlock()
		return
	}

	// 模拟业务处理时间（这段时间内，其他进程可能也在操作库存）
	fmt.Printf("%s: ⏳ 处理订单中（模拟耗时操作）...\n", processName)
	time.Sleep(1 * time.Second)

	// 扣减库存
	newStock := stockValue - purchaseQuantity
	redisClient.Set(stockKey, strconv.Itoa(newStock))
	fmt.Printf("%s: ✅ 扣减库存: %d - %d = %d\n", processName, stockValue, purchaseQuantity, newStock)

	mu.Unlock()
	fmt.Printf("%s: 🔓 释放本地锁\n", processName)

	// 最终库存
	finalStock, _ := redisClient.Get(stockKey)
	finalStockValue, _ := strconv.Atoi(finalStock)
	fmt.Printf("%s: 📊 最终库存: %d\n", processName, finalStockValue)

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%s 完成\n", processName)
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\n⚠️  问题分析：")
	fmt.Println("   1. 每个进程都有自己独立的 mutex，互不影响")
	fmt.Println("   2. 多个进程可以同时获取各自的锁")
	fmt.Println("   3. 导致多个进程同时读取到相同的库存值")
	fmt.Println("   4. 最终导致超卖问题！")
	fmt.Println(strings.Repeat("=", 80))
}
