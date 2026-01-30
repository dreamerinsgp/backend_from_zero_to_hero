package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
)

// Config 配置结构
type EvictionConfig struct {
	Redis struct {
		Host        string `json:"host" yaml:"host"`
		Password    string `json:"password" yaml:"password"`
		Type        string `json:"type" yaml:"type"`
		PingTimeout string `json:"ping_timeout" yaml:"ping_timeout"`
	} `json:"redis" yaml:"redis"`
}

// 支持的淘汰策略
var evictionPolicies = []string{
	"allkeys-lru",
	"volatile-lru",
	"allkeys-lfu",
	"volatile-lfu",
	"allkeys-random",
	"volatile-random",
	"volatile-ttl",
	"noeviction",
}

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	policy := os.Args[1]
	if !isValidPolicy(policy) {
		fmt.Printf("❌ 无效的淘汰策略: %s\n", policy)
		fmt.Println("\n支持的策略：")
		for _, p := range evictionPolicies {
			fmt.Printf("  - %s\n", p)
		}
		return
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("【内存淘汰策略实验】%s\n", policy)
	fmt.Println(strings.Repeat("=", 80))

	// 加载配置
	var c EvictionConfig
	err := conf.Load("config.yaml", &c)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		return
	}

	// 初始化Redis（使用go-redis直接连接，便于执行CONFIG命令）
	addr := c.Redis.Host
	if addr == "" {
		addr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.Redis.Password,
		DB:       0,
	})

	ctx := context.Background()

	// 测试连接
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接Redis失败: %v\n", err)
		return
	}

	// 执行实验
	runEvictionTest(ctx, rdb, policy)
}

func showUsage() {
	fmt.Println("使用方法：")
	fmt.Println("  go run test_eviction_policy.go <policy>")
	fmt.Println("")
	fmt.Println("支持的策略：")
	for _, p := range evictionPolicies {
		fmt.Printf("  - %s\n", p)
	}
	fmt.Println("")
	fmt.Println("示例：")
	fmt.Println("  go run test_eviction_policy.go allkeys-lru")
	fmt.Println("  go run test_eviction_policy.go volatile-ttl")
}

func isValidPolicy(policy string) bool {
	for _, p := range evictionPolicies {
		if p == policy {
			return true
		}
	}
	return false
}

func runEvictionTest(ctx context.Context, rdb *redis.Client, policy string) {
	fmt.Printf("\n【步骤1】清理测试数据\n")
	fmt.Println(strings.Repeat("-", 80))
	rdb.FlushAll(ctx)
	fmt.Println("✅ 已清理所有数据")

	fmt.Printf("\n【步骤2】设置内存限制和淘汰策略\n")
	fmt.Println(strings.Repeat("-", 80))

	// 设置较小的内存限制（5MB，便于测试）
	maxMemory := "5mb"
	err := rdb.ConfigSet(ctx, "maxmemory", maxMemory).Err()
	if err != nil {
		fmt.Printf("❌ 设置maxmemory失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 设置 maxmemory = %s\n", maxMemory)

	// 设置淘汰策略
	err = rdb.ConfigSet(ctx, "maxmemory-policy", policy).Err()
	if err != nil {
		fmt.Printf("❌ 设置maxmemory-policy失败: %v\n", err)
		return
	}
	fmt.Printf("✅ 设置 maxmemory-policy = %s\n", policy)

	// 验证配置
	maxMemoryResult, _ := rdb.ConfigGet(ctx, "maxmemory").Result()
	maxMemoryPolicyResult, _ := rdb.ConfigGet(ctx, "maxmemory-policy").Result()
	fmt.Printf("📊 当前配置: maxmemory=%v, policy=%v\n", maxMemoryResult, maxMemoryPolicyResult)

	fmt.Printf("\n【步骤3】填充数据直到内存满\n")
	fmt.Println(strings.Repeat("-", 80))

	// 记录初始key数量
	initialKeys := getKeyCount(ctx, rdb)
	fmt.Printf("初始key数量: %d\n", initialKeys)

	// 填充数据（每个key约100KB）
	keySize := 100 * 1024 // 100KB
	value := strings.Repeat("x", keySize)
	keys := []string{}

	fmt.Println("开始填充数据...")
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("test:key:%d", i)

		// 根据策略设置过期时间
		if strings.HasPrefix(policy, "volatile-") {
			// volatile策略需要设置过期时间
			// 为了演示TTL策略，设置不同的过期时间
			var expire time.Duration
			if policy == "volatile-ttl" {
				// TTL策略：设置不同的过期时间，便于观察
				expire = time.Duration(1000+i*10) * time.Second // 不同的TTL
			} else {
				expire = time.Hour // 1小时过期
			}
			err := rdb.SetEX(ctx, key, value, expire).Err()
			if err != nil {
				// 内存可能已满，记录被淘汰的key
				fmt.Printf("⚠️  写入失败 (可能触发淘汰): key=%s, error=%v\n", key, err)
				break
			}
		} else {
			// allkeys策略不需要过期时间
			err := rdb.Set(ctx, key, value, 0).Err()
			if err != nil {
				fmt.Printf("⚠️  写入失败 (可能触发淘汰): key=%s, error=%v\n", key, err)
				break
			}
		}

		keys = append(keys, key)
		if (i+1)%10 == 0 {
			memoryInfo := getMemoryInfo(ctx, rdb)
			fmt.Printf("已写入 %d 个key, 内存使用: %s / %s\n", i+1, memoryInfo["used"], memoryInfo["max"])
		}
	}

	currentKeys := getKeyCount(ctx, rdb)
	fmt.Printf("\n✅ 填充完成，当前key数量: %d\n", currentKeys)

	// 模拟访问模式（用于LRU/LFU策略）
	if strings.Contains(policy, "lru") || strings.Contains(policy, "lfu") {
		fmt.Printf("\n【步骤3.5】模拟访问模式（用于LRU/LFU策略）\n")
		fmt.Println(strings.Repeat("-", 80))

		// 访问前10个key多次（模拟热点数据）
		fmt.Println("访问前10个key（模拟热点数据）...")
		for i := 0; i < 10 && i < len(keys); i++ {
			rdb.Get(ctx, keys[i]) // 访问key，更新LRU/LFU信息
		}
		fmt.Println("✅ 已访问前10个key（这些key应该被保留）")

		// 对于LFU策略，需要多次访问某些key
		if strings.Contains(policy, "lfu") {
			fmt.Println("多次访问key 0-4（提高访问频率）...")
			for j := 0; j < 5; j++ {
				for i := 0; i < 5 && i < len(keys); i++ {
					rdb.Get(ctx, keys[i])
				}
			}
			fmt.Println("✅ 已多次访问key 0-4（这些key访问频率高，应该被保留）")
		}
	}

	fmt.Printf("\n【步骤4】继续写入新数据，观察淘汰行为\n")
	fmt.Println(strings.Repeat("-", 80))

	// 记录当前存在的key
	existingKeysBefore := getExistingKeys(ctx, rdb, keys)
	fmt.Printf("写入新数据前存在的key数量: %d\n", len(existingKeysBefore))

	// 尝试写入新数据（触发淘汰）
	newKeys := []string{}
	for i := 100; i < 120; i++ {
		key := fmt.Sprintf("test:new:key:%d", i)
		newKeys = append(newKeys, key)

		if strings.HasPrefix(policy, "volatile-") {
			err := rdb.SetEX(ctx, key, value, time.Hour).Err()
			if err != nil {
				if policy == "noeviction" {
					fmt.Printf("❌ 写入失败 (noeviction策略): key=%s, error=%v\n", key, err)
				} else {
					fmt.Printf("⚠️  写入失败: key=%s, error=%v\n", key, err)
				}
				break
			}
		} else {
			err := rdb.Set(ctx, key, value, 0).Err()
			if err != nil {
				if policy == "noeviction" {
					fmt.Printf("❌ 写入失败 (noeviction策略): key=%s, error=%v\n", key, err)
				} else {
					fmt.Printf("⚠️  写入失败: key=%s, error=%v\n", key, err)
				}
				break
			}
		}

		fmt.Printf("✅ 写入新key: %s\n", key)
		time.Sleep(100 * time.Millisecond) // 短暂延迟，便于观察
	}

	// 检查哪些key被淘汰了
	existingKeysAfter := getExistingKeys(ctx, rdb, keys)
	evictedKeys := findEvictedKeys(existingKeysBefore, existingKeysAfter)

	fmt.Printf("\n【步骤5】分析淘汰结果\n")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("写入新数据前存在的key数量: %d\n", len(existingKeysBefore))
	fmt.Printf("写入新数据后存在的key数量: %d\n", len(existingKeysAfter))
	fmt.Printf("被淘汰的key数量: %d\n", len(evictedKeys))

	if len(evictedKeys) > 0 {
		fmt.Println("\n被淘汰的key（前10个）：")
		for i, key := range evictedKeys {
			if i >= 10 {
				break
			}
			fmt.Printf("  - %s\n", key)
		}
		if len(evictedKeys) > 10 {
			fmt.Printf("  ... 还有 %d 个key被淘汰\n", len(evictedKeys)-10)
		}
	}

	// 显示内存信息
	memoryInfo := getMemoryInfo(ctx, rdb)
	fmt.Printf("\n📊 最终内存使用: %s / %s\n", memoryInfo["used"], memoryInfo["max"])

	// 显示统计信息
	stats := getStats(ctx, rdb)
	fmt.Printf("📈 淘汰统计: evicted_keys=%s\n", stats["evicted_keys"])

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("实验完成！策略: %s\n", policy)
	fmt.Println(strings.Repeat("=", 80))
}

func getKeyCount(ctx context.Context, rdb *redis.Client) int {
	count, err := rdb.DBSize(ctx).Result()
	if err != nil {
		return 0
	}
	return int(count)
}

func getExistingKeys(ctx context.Context, rdb *redis.Client, keys []string) []string {
	existing := []string{}
	for _, key := range keys {
		exists, _ := rdb.Exists(ctx, key).Result()
		if exists > 0 {
			existing = append(existing, key)
		}
	}
	return existing
}

func findEvictedKeys(before, after []string) []string {
	afterMap := make(map[string]bool)
	for _, key := range after {
		afterMap[key] = true
	}

	evicted := []string{}
	for _, key := range before {
		if !afterMap[key] {
			evicted = append(evicted, key)
		}
	}
	return evicted
}

func getMemoryInfo(ctx context.Context, rdb *redis.Client) map[string]string {
	result := make(map[string]string)

	info, err := rdb.Info(ctx, "memory").Result()
	if err == nil {
		lines := strings.Split(info, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "used_memory_human:") {
				parts := strings.Split(line, ":")
				if len(parts) == 2 {
					result["used"] = strings.TrimSpace(parts[1])
				}
			}
			if strings.HasPrefix(line, "maxmemory_human:") {
				parts := strings.Split(line, ":")
				if len(parts) == 2 {
					result["max"] = strings.TrimSpace(parts[1])
				}
			}
		}
	}

	// 如果解析失败，使用原始值
	if result["used"] == "" {
		used, _ := rdb.Info(ctx, "memory").Result()
		if strings.Contains(used, "used_memory:") {
			for _, line := range strings.Split(used, "\n") {
				if strings.HasPrefix(line, "used_memory:") {
					parts := strings.Split(line, ":")
					if len(parts) == 2 {
						result["used"] = strings.TrimSpace(parts[1]) + " bytes"
					}
				}
			}
		}
	}
	if result["max"] == "" {
		max, _ := rdb.ConfigGet(ctx, "maxmemory").Result()
		if len(max) > 0 {
			result["max"] = fmt.Sprintf("%v", max)
		}
	}

	return result
}

func getStats(ctx context.Context, rdb *redis.Client) map[string]string {
	result := make(map[string]string)

	info, err := rdb.Info(ctx, "stats").Result()
	if err == nil {
		lines := strings.Split(info, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "evicted_keys:") {
				parts := strings.Split(line, ":")
				if len(parts) == 2 {
					result["evicted_keys"] = strings.TrimSpace(parts[1])
				}
			}
		}
	}

	if result["evicted_keys"] == "" {
		result["evicted_keys"] = "0"
	}

	return result
}
