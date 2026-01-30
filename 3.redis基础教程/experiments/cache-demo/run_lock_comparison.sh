#!/bin/bash

# 分布式锁对比实验脚本
# 用于对比普通锁（mutex）和分布式锁的效果

echo "=========================================="
echo "分布式锁对比实验"
echo "=========================================="
echo ""

# 检查Redis是否运行
if ! redis-cli ping > /dev/null 2>&1; then
    echo "❌ 错误: Redis未运行，请先启动Redis"
    echo "   启动命令: redis-server"
    exit 1
fi

echo "✅ Redis运行正常"
echo ""

# 重置库存
echo "📦 重置库存为100..."
redis-cli SET stock:product:1001 100 > /dev/null 2>&1
echo "✅ 库存已重置为100"
echo ""

# 选择实验
if [ -z "$1" ]; then
    echo "请选择实验："
    echo "  1. 实验1: 多进程普通锁（mutex）- 展示问题"
    echo "  2. 实验2: 多进程分布式锁 - 正确解决方案"
    echo "  3. 对比实验: 同时运行两个实验进行对比"
    echo ""
    echo "使用方法: $0 [1|2|3]"
    exit 1
fi

EXPERIMENT=$1

case $EXPERIMENT in
    1)
        echo "=========================================="
        echo "【实验1】多进程普通锁（mutex）- 展示问题"
        echo "=========================================="
        echo ""
        echo "⚠️  注意：将启动3个进程，观察超卖问题"
        echo ""
        read -p "按Enter键开始..."
        echo ""
        
        # 重置库存
        redis-cli SET stock:product:1001 100 > /dev/null 2>&1
        
        # 启动3个进程
        cd lock-demo
        echo "启动进程1..."
        go run test_local_lock.go 进程A &
        PID1=$!
        
        sleep 0.5
        
        echo "启动进程2..."
        go run test_local_lock.go 进程B &
        PID2=$!
        
        sleep 0.5
        
        echo "启动进程3..."
        go run test_local_lock.go 进程C &
        PID3=$!
        
        echo ""
        echo "等待所有进程完成..."
        wait $PID1 $PID2 $PID3
        
        echo ""
        echo "=========================================="
        echo "实验1完成"
        echo "=========================================="
        echo ""
        echo "📊 最终库存:"
        redis-cli GET stock:product:1001
        echo ""
        echo "⚠️  问题：多个进程同时操作，导致超卖！"
        ;;
        
    2)
        echo "=========================================="
        echo "【实验2】多进程分布式锁 - 正确解决方案"
        echo "=========================================="
        echo ""
        echo "✅ 将启动3个进程，观察分布式锁如何防止超卖"
        echo ""
        read -p "按Enter键开始..."
        echo ""
        
        # 重置库存
        redis-cli SET stock:product:1001 100 > /dev/null 2>&1
        
        # 启动3个进程
        cd lock-demo
        echo "启动进程1..."
        go run test_distributed_lock.go 进程A &
        PID1=$!
        
        sleep 0.5
        
        echo "启动进程2..."
        go run test_distributed_lock.go 进程B &
        PID2=$!
        
        sleep 0.5
        
        echo "启动进程3..."
        go run test_distributed_lock.go 进程C &
        PID3=$!
        
        echo ""
        echo "等待所有进程完成..."
        wait $PID1 $PID2 $PID3
        
        echo ""
        echo "=========================================="
        echo "实验2完成"
        echo "=========================================="
        echo ""
        echo "📊 最终库存:"
        redis-cli GET stock:product:1001
        echo ""
        echo "✅ 优势：分布式锁确保同一时刻只有一个进程操作库存！"
        ;;
        
    3)
        echo "=========================================="
        echo "【对比实验】普通锁 vs 分布式锁"
        echo "=========================================="
        echo ""
        
        # 实验1：普通锁
        echo ""
        echo "=========================================="
        echo "第一步：运行实验1（普通锁）"
        echo "=========================================="
        echo ""
        read -p "按Enter键开始实验1..."
        
        redis-cli SET stock:product:1001 100 > /dev/null 2>&1
        
        cd lock-demo
        go run test_local_lock.go 进程A &
        PID1=$!
        sleep 0.5
        go run test_local_lock.go 进程B &
        PID2=$!
        sleep 0.5
        go run test_local_lock.go 进程C &
        PID3=$!
        
        wait $PID1 $PID2 $PID3
        cd ..
        
        FINAL_STOCK1=$(redis-cli GET stock:product:1001)
        echo ""
        echo "实验1完成，最终库存: $FINAL_STOCK1"
        echo ""
        
        # 等待
        sleep 3
        
        # 实验2：分布式锁
        echo ""
        echo "=========================================="
        echo "第二步：运行实验2（分布式锁）"
        echo "=========================================="
        echo ""
        read -p "按Enter键开始实验2..."
        
        redis-cli SET stock:product:1001 100 > /dev/null 2>&1
        
        cd lock-demo
        go run test_distributed_lock.go 进程A &
        PID1=$!
        sleep 0.5
        go run test_distributed_lock.go 进程B &
        PID2=$!
        sleep 0.5
        go run test_distributed_lock.go 进程C &
        PID3=$!
        
        wait $PID1 $PID2 $PID3
        cd ..
        
        wait $PID1 $PID2 $PID3
        
        FINAL_STOCK2=$(redis-cli GET stock:product:1001)
        echo ""
        echo "实验2完成，最终库存: $FINAL_STOCK2"
        echo ""
        
        # 对比结果
        echo "=========================================="
        echo "对比结果"
        echo "=========================================="
        echo ""
        echo "实验1（普通锁）最终库存: $FINAL_STOCK1"
        echo "实验2（分布式锁）最终库存: $FINAL_STOCK2"
        echo ""
        echo "分析："
        echo "  - 实验1：每个进程都有独立的mutex，可以同时操作，导致超卖"
        echo "  - 实验2：所有进程共享同一个分布式锁，顺序执行，防止超卖"
        echo ""
        ;;
        
    *)
        echo "❌ 无效的实验编号: $EXPERIMENT"
        echo "   请使用: 1, 2, 或 3"
        exit 1
        ;;
esac
