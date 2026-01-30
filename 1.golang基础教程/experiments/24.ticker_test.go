package experiments

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestTickerBasicUsage(t *testing.T) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	var count int
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				count++
				if count >= 5 {
					done <- true
					return
				}
			}
		}
	}()

	<-done

	if count != 5 {
		t.Errorf("期望触发5此，实际触发:%d次", count)
	}
}

// TestTickerWithSelect 测试 Ticker 配合 select 语句使用
func TestTickerWithSelect(t *testing.T) {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	var count int
	for {
		select {
		case <-ctx.Done():
			// 超时退出
			if count < 3 {
				t.Errorf("期望至少触发3次，实际触发%d次", count)
			}
			return
		case <-ticker.C:
			count++
			t.Logf("Ticker 触发第 %d 次", count)
		}
	}
}

// TestTickerStop 测试停止 Ticker 后不再触发
func TestTickerStop(t *testing.T) {
	ticker := time.NewTicker(50 * time.Millisecond)

	var count int
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				count++
				if count == 3 {
					ticker.Stop() // 停止 ticker
					done <- true
					return
				}
			}
		}
	}()

	<-done

	// 等待一段时间，确认 ticker 已经停止
	time.Sleep(200 * time.Millisecond)

	if count != 3 {
		t.Errorf("期望触发3次后停止，实际触发%d次", count)
	}
}

// TestTickerRateLimiting 测试使用 Ticker 实现限流控制
func TestTickerRateLimiting(t *testing.T) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	var processedCount int
	var totalTasks = 10

	// 模拟需要处理的任务
	tasks := make(chan int, totalTasks)
	for i := 0; i < totalTasks; i++ {
		tasks <- i
	}
	close(tasks)

	// 使用 ticker 控制处理频率
	done := make(chan bool)
	go func() {
		for task := range tasks {
			<-ticker.C // 等待 ticker 触发，实现限流
			processedCount++
			t.Logf("处理任务 %d", task)
		}
		done <- true
	}()

	<-done

	if processedCount != totalTasks {
		t.Errorf("期望处理%d个任务，实际处理%d个", totalTasks, processedCount)
	}
}

// TestTickerPeriodicCheck 测试周期性检查场景（类似项目中的用法）
func TestTickerPeriodicCheck(t *testing.T) {
	checkTicker := time.NewTicker(100 * time.Millisecond)
	sendTicker := time.NewTicker(50 * time.Millisecond)
	defer checkTicker.Stop()
	defer sendTicker.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	var checkCount int
	var sendCount int
	items := []int{1, 2, 3, 4, 5}

	for {
		select {
		case <-ctx.Done():
			// 验证周期性检查
			if checkCount == 0 {
				t.Error("期望至少执行一次检查")
			}
			return
		case <-checkTicker.C:
			checkCount++
			t.Logf("执行第 %d 次检查，发现 %d 个待处理项", checkCount, len(items))

			// 模拟处理每个项目，使用 sendTicker 控制发送频率
			for _, item := range items {
				select {
				case <-ctx.Done():
					return
				case <-sendTicker.C:
					sendCount++
					t.Logf("发送项目 %d", item)
				}
			}
		}
	}
}

// TestTickerWithGoroutine 测试多个 goroutine 使用同一个 Ticker
func TestTickerWithGoroutine(t *testing.T) {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	var wg sync.WaitGroup
	var mu sync.Mutex
	var totalCount int

	// 启动3个 goroutine 共享同一个 ticker
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			count := 0
			for count < 3 {
				<-ticker.C
				mu.Lock()
				totalCount++
				mu.Unlock()
				count++
				t.Logf("Goroutine %d 收到第 %d 次触发", id, count)
			}
		}(i)
	}

	wg.Wait()

	// 3个 goroutine，每个触发3次，总共9次
	if totalCount != 9 {
		t.Errorf("期望总共触发9次，实际触发%d次", totalCount)
	}
}

// TestTickerVsTimer 对比 Ticker 和 Timer 的区别
func TestTickerVsTimer(t *testing.T) {
	// Timer 只触发一次
	timer := time.NewTimer(100 * time.Millisecond)
	defer timer.Stop()

	timerCount := 0
	<-timer.C
	timerCount++

	// 等待一段时间，确认 timer 不会再次触发
	time.Sleep(200 * time.Millisecond)

	if timerCount != 1 {
		t.Errorf("Timer 应该只触发1次，实际触发%d次", timerCount)
	}

	// Ticker 会周期性触发
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	tickerCount := 0
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				tickerCount++
				if tickerCount >= 3 {
					done <- true
					return
				}
			}
		}
	}()

	<-done

	if tickerCount != 3 {
		t.Errorf("Ticker 应该触发3次，实际触发%d次", tickerCount)
	}
}

// TestTickerReset 测试 Ticker 的 Reset 方法（注意：Ticker 没有 Reset，但可以重新创建）
func TestTickerRecreate(t *testing.T) {
	// Ticker 没有 Reset 方法，如果需要改变间隔，需要停止旧的并创建新的
	ticker := time.NewTicker(100 * time.Millisecond)

	var count int
	done := make(chan bool)

	go func() {
		for i := 0; i < 2; i++ {
			<-ticker.C
			count++
		}
		ticker.Stop() // 停止旧的 ticker

		// 创建新的 ticker，使用不同的间隔
		ticker = time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for i := 0; i < 2; i++ {
			<-ticker.C
			count++
		}
		done <- true
	}()

	<-done

	if count != 4 {
		t.Errorf("期望触发4次，实际触发%d次", count)
	}
}

// TestTickerMemoryLeak 演示如果不停止 Ticker 可能导致的问题
func TestTickerMemoryLeak(t *testing.T) {
	// 正确做法：使用 defer 确保停止
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop() // 确保资源释放

	// 错误做法示例（注释掉，避免实际泄漏）
	// ticker2 := time.NewTicker(100 * time.Millisecond)
	// 忘记调用 ticker2.Stop() 会导致内存泄漏

	var count int
	done := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			<-ticker.C
			count++
		}
		done <- true
	}()

	<-done

	if count != 3 {
		t.Errorf("期望触发3次，实际触发%d次", count)
	}
	// ticker 会在 defer 中自动停止
}
