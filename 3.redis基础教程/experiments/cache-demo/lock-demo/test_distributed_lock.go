package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// ✅ 正确：使用分布式锁（Redis）
func Lock(ctx context.Context, kv *redis.Redis, key string, expire int) (lock *redis.RedisLock, ok bool, err error) {
	lock = redis.NewRedisLock(kv, key)
	lock.SetExpire(expire)
	ok, err = lock.AcquireCtx(ctx)
	if err != nil {
		err = fmt.Errorf("lock AcquireCtx err: %w", err)
		return
	}
	return
}

func ReleaseLock(lock *redis.RedisLock) {
	if lock != nil {
		_, err := lock.Release()
		if err != nil {
			fmt.Printf("释放锁失败: %v\n", err)
		}
	}
}

func main() {
	// 获取进程ID
	processID := os.Getpid()

	// 获取命令行参数（进程标识）
	processName := fmt.Sprintf("进程-%d", processID)
	if len(os.Args) > 1 {
		processName = os.Args[1]
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("【实验2】多进程分布式锁 - 正确解决方案\n")
	fmt.Printf("%s 启动 (PID: %d)\n", processName, processID)
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("✅ 分布式锁可以跨进程/跨服务器控制，所有进程共享同一个锁！")
	fmt.Println(strings.Repeat("=", 80))

	// 加载配置
	var c LockConfig
	err := conf.Load("../config.yaml", &c)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		return
	}

	// 初始化Redis
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

	// 场景：库存扣减（防止超卖）
	fmt.Printf("\n【场景】库存扣减 - 使用分布式锁\n")
	fmt.Println(strings.Repeat("-", 80))

	productID := "product:1001"
	lockKey := fmt.Sprintf("lock:stock:%s", productID)

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

	// ✅ 使用分布式锁（Redis）- 所有进程共享同一个锁
	fmt.Printf("%s: 🔒 尝试获取分布式锁 '%s'...\n", processName, lockKey)
	ctx := context.Background()
	lock, ok, err := Lock(ctx, redisClient, lockKey, 10)
	if err != nil {
		fmt.Printf("%s: ❌ 获取锁时出错: %v\n", processName, err)
		return
	}

	if !ok {
		fmt.Printf("%s: ⏳ 获取锁失败（锁已被其他进程持有），等待中...\n", processName)

		// 重试获取锁（最多等待30秒）
		maxRetries := 30
		for i := 0; i < maxRetries; i++ {
			time.Sleep(1 * time.Second)
			lock, ok, err = Lock(ctx, redisClient, lockKey, 10)
			if err != nil {
				fmt.Printf("%s: ❌ 获取锁时出错: %v\n", processName, err)
				return
			}
			if ok {
				fmt.Printf("%s: ✅ 获取分布式锁成功（等待了 %d 秒）\n", processName, i+1)
				break
			}
			if i == maxRetries-1 {
				fmt.Printf("%s: ❌ 等待超时，无法获取锁\n", processName)
				return
			}
		}
	} else {
		fmt.Printf("%s: ✅ 获取分布式锁成功\n", processName)
	}

	// 确保释放锁
	defer func() {
		ReleaseLock(lock)
		fmt.Printf("%s: 🔓 释放分布式锁\n", processName)
	}()

	// 重新读取库存（双重检查）
	currentStock, _ = redisClient.Get(stockKey)
	stockValue, _ = strconv.Atoi(currentStock)
	fmt.Printf("%s: 📖 重新读取库存: %d\n", processName, stockValue)

	if stockValue < purchaseQuantity {
		fmt.Printf("%s: ❌ 库存不足（当前: %d, 需要: %d）\n", processName, stockValue, purchaseQuantity)
		return
	}

	// 模拟业务处理时间
	fmt.Printf("%s: ⏳ 处理订单中（模拟耗时操作）...\n", processName)
	time.Sleep(1 * time.Second)

	// 扣减库存
	newStock := stockValue - purchaseQuantity
	redisClient.Set(stockKey, strconv.Itoa(newStock))
	fmt.Printf("%s: ✅ 扣减库存: %d - %d = %d\n", processName, stockValue, purchaseQuantity, newStock)

	// 最终库存
	finalStock, _ := redisClient.Get(stockKey)
	finalStockValue, _ := strconv.Atoi(finalStock)
	fmt.Printf("%s: 📊 最终库存: %d\n", processName, finalStockValue)

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%s 完成\n", processName)
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("\n✅ 优势分析：")
	fmt.Println("   1. 所有进程共享同一个Redis锁")
	fmt.Println("   2. 同一时刻只有一个进程能获取锁")
	fmt.Println("   3. 其他进程必须等待锁释放")
	fmt.Println("   4. 有效防止超卖问题！")
	fmt.Println(strings.Repeat("=", 80))
}
