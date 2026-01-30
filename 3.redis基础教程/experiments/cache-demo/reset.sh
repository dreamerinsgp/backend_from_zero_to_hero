#!/bin/bash

# 重置缓存脚本
# 功能：只清理 Redis 中的所有用户缓存，不重置数据库

echo "========== 重置缓存 =========="

# 配置
REDIS_PASSWORD=""  # 如果Redis有密码，填写密码

# 清理 Redis 缓存
echo ""
echo "清理 Redis 缓存..."
if [ -z "$REDIS_PASSWORD" ]; then
    KEYS=$(redis-cli KEYS "user:*" 2>/dev/null)
    if [ -n "$KEYS" ]; then
        echo "$KEYS" | xargs -r redis-cli DEL
        KEY_COUNT=$(echo "$KEYS" | grep -v "^$" | wc -l)
        echo "✓ 已清理 $KEY_COUNT 个缓存Key（无密码）"
    else
        echo "✓ 缓存已为空，无需清理"
    fi
else
    KEYS=$(redis-cli -a "$REDIS_PASSWORD" KEYS "user:*" 2>/dev/null)
    if [ -n "$KEYS" ]; then
        echo "$KEYS" | xargs -r redis-cli -a "$REDIS_PASSWORD" DEL
        KEY_COUNT=$(echo "$KEYS" | grep -v "^$" | wc -l)
        echo "✓ 已清理 $KEY_COUNT 个缓存Key（使用密码）"
    else
        echo "✓ 缓存已为空，无需清理"
    fi
fi

# 验证缓存是否已清理
if [ -z "$REDIS_PASSWORD" ]; then
    KEYS_OUTPUT=$(redis-cli KEYS "user:*" 2>/dev/null)
    if [ -z "$KEYS_OUTPUT" ]; then
        CACHE_COUNT=0
    else
        CACHE_COUNT=$(echo "$KEYS_OUTPUT" | grep -v "^$" | wc -l)
    fi
else
    KEYS_OUTPUT=$(redis-cli -a "$REDIS_PASSWORD" KEYS "user:*" 2>/dev/null)
    if [ -z "$KEYS_OUTPUT" ]; then
        CACHE_COUNT=0
    else
        CACHE_COUNT=$(echo "$KEYS_OUTPUT" | grep -v "^$" | wc -l)
    fi
fi
echo "  剩余缓存数量: $CACHE_COUNT"

echo ""
echo "========== 缓存重置完成 =========="
echo ""
echo "提示: 如需重置数据库，请使用: go run main.go reset-db"
echo "现在可以运行程序进行新的实验："
echo "  go run main.go"
echo ""
