# 内存淘汰策略实验

## 简介

这个实验演示了Redis的8种内存淘汰策略，通过实际触发内存淘汰来观察不同策略的行为差异。

## 文件说明

- `test_eviction_policy.go` - 主实验程序
- `run_eviction_test.sh` - 实验脚本（支持单个策略或全部策略对比）
- `测试说明_内存淘汰策略.md` - 详细的使用说明和预期结果

## 快速开始

### 1. 测试单个策略

```bash
cd /home/ubuntu/dex_full/web3fun-Dex/1.基础篇/3.redis基础教程/experiments/cache-demo

# 使用脚本
./run_eviction_test.sh allkeys-lru

# 或直接运行
go run test_eviction_policy.go allkeys-lru
```

### 2. 测试所有策略（对比）

```bash
./run_eviction_test.sh all
```

### 3. 查看支持的策略

```bash
go run test_eviction_policy.go
```

## 实验流程

1. **清理数据** - 清空Redis
2. **设置配置** - 设置maxmemory=5MB和淘汰策略
3. **填充数据** - 写入约100个key（每个100KB），直到内存满
4. **模拟访问** - 对于LRU/LFU策略，模拟访问模式
5. **触发淘汰** - 继续写入新数据，触发淘汰
6. **分析结果** - 统计被淘汰的key，分析策略行为

## 支持的策略

| 策略 | 说明 | 推荐度 |
|------|------|--------|
| `allkeys-lru` | 所有key，LRU算法 | ⭐⭐⭐⭐⭐ 推荐 |
| `volatile-lru` | 设置了过期时间的key，LRU算法 | ⭐⭐⭐⭐ |
| `allkeys-lfu` | 所有key，LFU算法 | ⭐⭐⭐⭐ |
| `volatile-lfu` | 设置了过期时间的key，LFU算法 | ⭐⭐⭐ |
| `volatile-ttl` | 设置了过期时间的key，TTL最小 | ⭐⭐⭐ |
| `allkeys-random` | 所有key，随机 | ⭐（测试用） |
| `volatile-random` | 设置了过期时间的key，随机 | ⭐（测试用） |
| `noeviction` | 不淘汰，返回错误 | ⭐⭐（特殊场景） |

## 预期结果示例

### allkeys-lru

```
写入新数据前存在的key数量: 50
写入新数据后存在的key数量: 45
被淘汰的key数量: 5

被淘汰的key（前10个）：
  - test:key:0  ← 最久未访问
  - test:key:1
  - test:key:2
  ...
```

### volatile-ttl

```
被淘汰的key（前10个）：
  - test:key:0  ← TTL最小（即将过期）
  - test:key:1
  ...
```

### noeviction

```
❌ 写入失败 (noeviction策略): key=test:new:key:100, error=OOM command not allowed when used memory > 'maxmemory'
```

## 注意事项

1. **内存限制**：实验设置maxmemory=5MB，便于快速触发淘汰
2. **数据大小**：每个key约100KB，写入约50个key即可填满5MB
3. **volatile策略**：需要设置过期时间，实验会自动设置
4. **访问模式**：LRU/LFU策略会模拟访问模式，便于观察差异

## 详细说明

请参考 `测试说明_内存淘汰策略.md` 获取更详细的使用说明和预期结果。
