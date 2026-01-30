#!/bin/bash

# 内存淘汰策略实验脚本

echo "=========================================="
echo "Redis 内存淘汰策略实验"
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

# 支持的策略
POLICIES=("allkeys-lru" "volatile-lru" "allkeys-lfu" "volatile-lfu" "allkeys-random" "volatile-random" "volatile-ttl" "noeviction")

# 检查参数
if [ -z "$1" ]; then
    echo "请选择要测试的淘汰策略："
    echo ""
    for i in "${!POLICIES[@]}"; do
        echo "  $((i+1)). ${POLICIES[$i]}"
    done
    echo ""
    echo "使用方法: $0 <策略名称>"
    echo "示例: $0 allkeys-lru"
    echo ""
    echo "或者运行所有策略对比: $0 all"
    exit 1
fi

POLICY=$1

if [ "$POLICY" = "all" ]; then
    echo "=========================================="
    echo "运行所有策略对比实验"
    echo "=========================================="
    echo ""
    
    for policy in "${POLICIES[@]}"; do
        echo ""
        echo "=========================================="
        echo "测试策略: $policy"
        echo "=========================================="
        echo ""
        go run test_eviction_policy.go "$policy"
        echo ""
        echo "按Enter继续下一个策略..."
        read
    done
    
    echo ""
    echo "=========================================="
    echo "所有策略测试完成！"
    echo "=========================================="
else
    # 验证策略是否有效
    VALID=false
    for p in "${POLICIES[@]}"; do
        if [ "$p" = "$POLICY" ]; then
            VALID=true
            break
        fi
    done
    
    if [ "$VALID" = false ]; then
        echo "❌ 无效的策略: $POLICY"
        echo ""
        echo "支持的策略："
        for p in "${POLICIES[@]}"; do
            echo "  - $p"
        done
        exit 1
    fi
    
    # 运行单个策略测试
    go run test_eviction_policy.go "$POLICY"
fi
