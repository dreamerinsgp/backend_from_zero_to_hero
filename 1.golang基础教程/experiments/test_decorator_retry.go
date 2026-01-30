package experiments

// import (
// 	"fmt"
// 	"math/rand"
// 	"strings"
// 	"sync/atomic"
// 	"time"
// )

// // ==================== 重试装饰器实现 ====================

// // RetryConfig 重试配置
// type RetryConfig struct {
// 	MaxRetries int           // 最大重试次数
// 	Delay      time.Duration // 重试延迟
// }

// // RetryDecorator 重试装饰器
// // fn: 需要重试的函数，返回error表示失败，nil表示成功
// func RetryDecorator(config RetryConfig, fn func() error) func() error {
// 	return func() error {
// 		var lastErr error

// 		for attempt := 0; attempt <= config.MaxRetries; attempt++ {
// 			// 执行函数
// 			err := fn()

// 			// 成功，返回
// 			if err == nil {
// 				if attempt > 0 {
// 					fmt.Printf("✅ 【重试成功】第 %d 次重试成功\n", attempt)
// 				}
// 				return nil
// 			}

// 			// 记录最后一次错误
// 			lastErr = err

// 			// 如果还有重试机会
// 			if attempt < config.MaxRetries {
// 				fmt.Printf("⚠️  【重试】第 %d 次失败: %v，%v 后重试...\n",
// 					attempt+1, err, config.Delay)
// 				time.Sleep(config.Delay)
// 			}
// 		}

// 		// 所有重试都失败
// 		return fmt.Errorf("❌ 【重试失败】重试 %d 次后仍然失败，最后一次错误: %w",
// 			config.MaxRetries, lastErr)
// 	}
// }

// // ==================== 测试场景 ====================

// // 模拟不稳定的API调用（随机失败）
// type UnstableAPI struct {
// 	successRate float64 // 成功率（0.0 - 1.0）
// 	callCount   int64   // 调用次数（原子操作）
// }

// func NewUnstableAPI(successRate float64) *UnstableAPI {
// 	return &UnstableAPI{
// 		successRate: successRate,
// 		callCount:   0,
// 	}
// }

// func (api *UnstableAPI) Call() error {
// 	atomic.AddInt64(&api.callCount, 1)
// 	count := atomic.LoadInt64(&api.callCount)

// 	// 模拟随机失败
// 	if rand.Float64() < api.successRate {
// 		fmt.Printf("  📞 API调用成功（第 %d 次调用）\n", count)
// 		return nil
// 	}

// 	fmt.Printf("  📞 API调用失败（第 %d 次调用）\n", count)
// 	return fmt.Errorf("API调用失败: 网络错误")
// }

// func (api *UnstableAPI) GetCallCount() int64 {
// 	return atomic.LoadInt64(&api.callCount)
// }

// func (api *UnstableAPI) Reset() {
// 	atomic.StoreInt64(&api.callCount, 0)
// }

// // ==================== 主函数 ====================

// func main() {
// 	// 设置随机种子
// 	rand.Seed(time.Now().UnixNano())

// 	fmt.Println(strings.Repeat("=", 80))
// 	fmt.Println("装饰器在重试机制中的应用演示")
// 	fmt.Println(strings.Repeat("=", 80))

// 	api := NewUnstableAPI(0.3) // 30%成功率

// 	// ==================== 不使用装饰器 ====================
// 	fmt.Println("\n【场景1】不使用装饰器 - 直接调用")
// 	fmt.Println(strings.Repeat("-", 80))
// 	fmt.Println("代码：")
// 	fmt.Println("  err := api.Call()")
// 	fmt.Println("  if err != nil {")
// 	fmt.Println("      // 需要手动处理错误，无法自动重试")
// 	fmt.Println("  }")
// 	fmt.Println("\n执行结果：")
// 	err1 := api.Call()
// 	if err1 != nil {
// 		fmt.Printf("  ❌ 失败 - %v\n", err1)
// 		fmt.Println("  💡 问题：失败后无法自动重试，需要手动编写重试逻辑")
// 	} else {
// 		fmt.Printf("  ✅ 成功\n")
// 	}

// 	api.Reset()

// 	// ==================== 使用装饰器 ====================
// 	fmt.Println("\n【场景2】使用装饰器 - 自动重试")
// 	fmt.Println(strings.Repeat("-", 80))
// 	fmt.Println("代码：")
// 	fmt.Println("  config := RetryConfig{MaxRetries: 3, Delay: 200ms}")
// 	fmt.Println("  decoratedCall := RetryDecorator(config, api.Call)")
// 	fmt.Println("  err := decoratedCall()  // 自动重试")
// 	fmt.Println("\n配置：")
// 	config := RetryConfig{
// 		MaxRetries: 3,
// 		Delay:      200 * time.Millisecond,
// 	}
// 	fmt.Printf("  - 最大重试次数: %d\n", config.MaxRetries)
// 	fmt.Printf("  - 重试延迟: %v\n", config.Delay)
// 	fmt.Printf("  - API成功率: %.0f%%\n", api.successRate*100)

// 	fmt.Println("\n执行结果：")
// 	decoratedCall := RetryDecorator(config, api.Call)
// 	start := time.Now()
// 	err2 := decoratedCall()
// 	duration := time.Since(start)

// 	fmt.Printf("\n最终结果：\n")
// 	fmt.Printf("  - 总调用次数: %d\n", api.GetCallCount())
// 	fmt.Printf("  - 总耗时: %v\n", duration)
// 	if err2 != nil {
// 		fmt.Printf("  - 状态: ❌ 失败 - %v\n", err2)
// 	} else {
// 		fmt.Printf("  - 状态: ✅ 成功\n")
// 	}

// 	// ==================== 总结 ====================
// 	fmt.Println("\n" + strings.Repeat("=", 80))
// 	fmt.Println("装饰器的优势")
// 	fmt.Println(strings.Repeat("=", 80))
// 	fmt.Println("✅ 不修改原函数代码（api.Call 保持不变）")
// 	fmt.Println("✅ 自动重试逻辑，代码更简洁")
// 	fmt.Println("✅ 重试逻辑可复用，可以应用到其他函数")
// 	fmt.Println("✅ 易于维护，重试策略集中管理")
// }
