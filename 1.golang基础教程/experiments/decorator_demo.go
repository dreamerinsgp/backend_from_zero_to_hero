package experiments

import (
	"fmt"
	"log"
	"time"
)

// ==================== 示例1：基本装饰器 ====================

// 原函数
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 日志装饰器
func logDecorator(fn func(string)) func(string) {
	return func(name string) {
		fmt.Println("【日志】函数调用开始")
		fn(name)
		fmt.Println("【日志】函数调用结束")
	}
}

// ==================== 示例2：带返回值的装饰器 ====================

// 原函数：计算两数之和
func add(a, b int) int {
	return a + b
}

// 日志装饰器（带返回值）
func logDecoratorWithReturn(fn func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Printf("【日志】调用函数: add(%d, %d)\n", a, b)
		result := fn(a, b)
		fmt.Printf("【日志】函数返回: %d\n", result)
		return result
	}
}

// ==================== 示例3：计时装饰器 ====================

// 计时装饰器
func timingDecorator(fn func()) func() {
	return func() {
		start := time.Now()
		fn()
		duration := time.Since(start)
		fmt.Printf("【性能】执行耗时: %v\n", duration)
	}
}

// ==================== 示例4：重试装饰器 ====================

// 重试装饰器
func retryDecorator(maxRetries int, delay time.Duration, fn func() error) func() error {
	return func() error {
		var lastErr error
		for i := 0; i < maxRetries; i++ {
			err := fn()
			if err == nil {
				if i > 0 {
					fmt.Printf("【重试】第 %d 次重试成功\n", i)
				}
				return nil
			}
			lastErr = err
			if i < maxRetries-1 {
				fmt.Printf("【重试】第 %d 次失败，%v 后重试...\n", i+1, delay)
				time.Sleep(delay)
			}
		}
		return fmt.Errorf("重试 %d 次后仍然失败: %w", maxRetries, lastErr)
	}
}

// 模拟可能失败的操作
func unstableOperation() error {
	if time.Now().Unix()%3 == 0 {
		return fmt.Errorf("操作失败")
	}
	fmt.Println("操作成功")
	return nil
}

// ==================== 示例5：错误处理装饰器 ====================

// 错误处理装饰器
func errorHandlerDecorator(fn func() error) func() error {
	return func() error {
		err := fn()
		if err != nil {
			log.Printf("【错误处理】捕获错误: %v\n", err)
			// 可以在这里添加错误恢复逻辑
		}
		return err
	}
}

// ==================== 示例6：缓存装饰器 ====================

// 缓存装饰器
func cacheDecorator(fn func(string) string) func(string) string {
	cache := make(map[string]string)
	return func(key string) string {
		if val, ok := cache[key]; ok {
			fmt.Printf("【缓存】命中缓存: %s -> %s\n", key, val)
			return val
		}
		fmt.Printf("【缓存】未命中，调用原函数\n")
		val := fn(key)
		cache[key] = val
		return val
	}
}

// 原函数：查询数据
func queryData(key string) string {
	fmt.Printf("查询数据库: %s\n", key)
	return fmt.Sprintf("data-for-%s", key)
}

// ==================== 示例7：链式组合装饰器 ====================

// 多个装饰器组合
func chainDecorators() {
	fmt.Println("\n=== 示例7：链式组合装饰器 ===")

	// 原函数
	originalFn := func() {
		fmt.Println("执行业务逻辑...")
		time.Sleep(100 * time.Millisecond)
	}

	// 组合：日志 -> 计时
	decoratedFn := logDecorator(
		func(name string) {
			timingDecorator(originalFn)()
		},
	)

	decoratedFn("test")
}

// ==================== 示例8：HTTP中间件模式（装饰器） ====================

type HandlerFunc func()

// 认证装饰器
func authMiddleware(next HandlerFunc) HandlerFunc {
	return func() {
		fmt.Println("【认证】验证token...")
		// 模拟认证逻辑
		if true { // 假设认证成功
			fmt.Println("【认证】认证成功")
			next()
		} else {
			fmt.Println("【认证】认证失败")
		}
	}
}

// 日志装饰器
func logMiddleware(next HandlerFunc) HandlerFunc {
	return func() {
		fmt.Println("【日志】请求开始")
		next()
		fmt.Println("【日志】请求结束")
	}
}

// 业务处理函数
func handleRequest() {
	fmt.Println("处理业务逻辑...")
}

// ==================== 主函数 ====================

func main() {
	fmt.Println("==========================================")
	fmt.Println("Go 装饰器模式示例")
	fmt.Println("==========================================")

	// 示例1：基本装饰器
	fmt.Println("\n=== 示例1：基本装饰器 ===")
	decoratedGreet := logDecorator(greet)
	decoratedGreet("Alice")

	// 示例2：带返回值的装饰器
	fmt.Println("\n=== 示例2：带返回值的装饰器 ===")
	decoratedAdd := logDecoratorWithReturn(add)
	result := decoratedAdd(3, 5)
	fmt.Printf("结果: %d\n", result)

	// 示例3：计时装饰器
	fmt.Println("\n=== 示例3：计时装饰器 ===")
	timedFn := timingDecorator(func() {
		fmt.Println("执行耗时操作...")
		time.Sleep(200 * time.Millisecond)
	})
	timedFn()

	// 示例4：重试装饰器
	fmt.Println("\n=== 示例4：重试装饰器 ===")
	retryFn := retryDecorator(3, 100*time.Millisecond, unstableOperation)
	err := retryFn()
	if err != nil {
		fmt.Printf("最终失败: %v\n", err)
	}

	// 示例5：错误处理装饰器
	fmt.Println("\n=== 示例5：错误处理装饰器 ===")
	errorHandledFn := errorHandlerDecorator(func() error {
		return fmt.Errorf("模拟错误")
	})
	errorHandledFn()

	// 示例6：缓存装饰器
	fmt.Println("\n=== 示例6：缓存装饰器 ===")
	cachedQuery := cacheDecorator(queryData)
	fmt.Println("第一次查询:")
	result1 := cachedQuery("user:123")
	fmt.Printf("结果: %s\n", result1)

	fmt.Println("\n第二次查询（应该命中缓存）:")
	result2 := cachedQuery("user:123")
	fmt.Printf("结果: %s\n", result2)

	// 示例7：链式组合
	fmt.Println("\n=== 示例7：链式组合装饰器 ===")
	chainDecorators()

	// 示例8：HTTP中间件模式
	fmt.Println("\n=== 示例8：HTTP中间件模式（装饰器） ===")
	handler := logMiddleware(authMiddleware(handleRequest))
	handler()

	fmt.Println("\n==========================================")
	fmt.Println("所有示例完成！")
	fmt.Println("==========================================")
}
