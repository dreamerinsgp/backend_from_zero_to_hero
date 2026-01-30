# Redis学习路线图 - 基于DEX项目实战
查看web3转行实战项目：https://web3ite.fun
## 项目Redis使用概览

本项目使用了以下Redis核心功能：

- **基础操作**: GET/SET/SETEX/SETNX
- **数据结构**: Hash、Set、Sorted Set、String
- **高级功能**: 分布式锁、Pub/Sub、布隆过滤器
- **应用场景**: 缓存、计数、服务发现、WebSocket消息广播、时间序列数据

---

## 第一阶段：基础入门（1-2周）

### 1.1 Redis基础概念

**学习目标**: 理解Redis是什么，为什么使用Redis

**学习内容**:

- Redis简介：内存数据库、键值存储
- Redis vs MySQL：何时使用Redis
- Redis数据类型概览：String、Hash、List、Set、Sorted Set

**项目对应**:

- 查看 `docker-compose.yml` 了解Redis部署
- 理解为什么项目需要Redis（缓存、性能优化）

**实践任务**:

```bash
# 启动Redis（参考项目docker-compose.yml）
docker run -d -p 6379:6379 redis:7.0

# 连接Redis
redis-cli

# 基础命令练习
SET name "DEX"
GET name
EXISTS name
DEL name
```

---

### 1.2 String类型 - 最基础的数据类型

**学习目标**: 掌握String类型的所有操作

**学习内容**:

- SET/GET：设置和获取值
- SETEX：设置值并指定过期时间
- SETNX：仅当键不存在时设置（原子操作）
- INCR：递增操作
- 过期时间：EXPIRE、TTL

**项目对应代码**:

- `model/redismodel/token.go:16` - GET操作
- `model/redismodel/token.go:59` - SETNXEX操作（带过期时间）
- `pkg/xredis/helius_time.go:14` - INCR操作
- `pkg/oklink/client.go:186` - SETEX操作（缓存HTTP响应）

**实践任务**:

1. 实现一个简单的缓存函数（参考token.go的模式）
2. 实现一个计数器（参考helius_time.go）
3. 理解缓存穿透和缓存雪崩的概念

---

### 1.3 Hash类型 - 存储对象

**学习目标**: 掌握Hash类型，理解何时使用Hash vs String

**学习内容**:

- HSET/HGET：设置和获取字段
- HGETALL：获取所有字段
- HDEL：删除字段
- HKEYS/HVALS：获取所有键或值

**项目对应代码**:

- `apps/websocket/internal/svc/redischannel.go:60` - HSET（订阅管理）
- `apps/websocket/internal/svc/redischannel.go:239` - HGETALL（获取所有订阅）
- `apps/websocket/internal/svc/redischannel.go:249` - HDEL（取消订阅）

**实践任务**:

1. 用Hash存储用户信息（对比用String存储JSON的优劣）
2. 实现WebSocket订阅管理（参考redischannel.go）
3. 理解Hash适合存储对象的场景

---

## 第二阶段：进阶数据结构（2-3周）

### 2.1 Set类型 - 集合操作

**学习目标**: 掌握Set类型，理解集合运算

**学习内容**:

- SADD/SREM：添加/删除成员
- SMEMBERS：获取所有成员
- SISMEMBER：判断成员是否存在
- SINTER/SUNION/SDIFF：集合运算

**项目对应代码**:

- `apps/websocket/internal/svc/websocketcontext.go:241` - SADD（用户客户端关联）
- `apps/websocket/internal/svc/websocketcontext.go:267` - SREM（移除客户端）

**实践任务**:

1. 实现用户在线状态管理（参考websocketcontext.go）
2. 实现标签系统（用Set存储标签）
3. 实现共同关注功能（使用SINTER）

---

### 2.2 Sorted Set类型 - 有序集合

**学习目标**: 掌握Sorted Set，理解时间序列数据存储

**学习内容**:

- ZADD：添加成员和分数
- ZRANGE/ZREVRANGE：按排名范围查询
- ZRANGEBYSCORE：按分数范围查询（重要！）
- ZREM：删除成员
- ZSCORE：获取成员分数

**项目对应代码**:

- `model/redismodel/kline.go:59` - ZRANGEBYSCORE（K线数据查询）
- 用于存储时间序列的金融数据（K线图）

**实践任务**:

1. 实现排行榜功能（ZREVRANGE）
2. 实现时间范围查询（ZRANGEBYSCORE，参考kline.go）
3. 实现延迟队列（用分数表示执行时间）

---

## 第三阶段：高级功能（3-4周）

### 3.1 分布式锁

**学习目标**: 理解分布式锁的原理和实现

**学习内容**:

- 为什么需要分布式锁
- SETNX + EXPIRE实现锁
- 锁的超时和续期
- 锁的释放（Lua脚本保证原子性）

**项目对应代码**:

- `pkg/xredis/lock.go` - 完整的分布式锁实现
  - `Lock()`: 获取锁
  - `MustLock()`: 带超时的获取锁
  - `ReleaseLock()`: 释放锁

**实践任务**:

1. 阅读并理解lock.go的实现
2. 实现一个简单的分布式锁
3. 理解死锁和锁超时的问题
4. 实现锁续期机制

**关键知识点**:

- SETNX的原子性
- 过期时间的重要性
- Lua脚本保证释放锁的原子性

---

### 3.2 Pub/Sub - 发布订阅

**学习目标**: 掌握Redis的发布订阅模式

**学习内容**:

- PUBLISH：发布消息
- SUBSCRIBE/UNSUBSCRIBE：订阅/取消订阅
- PSUBSCRIBE：模式订阅
- 消息队列 vs Pub/Sub的区别

**项目对应代码**:

- `apps/websocket/internal/svc/websocket_cluster.go:26` - SUBSCRIBE（节点间消息）
- `apps/websocket/internal/svc/websocket_cluster.go:118` - PUBLISH（消息转发）

**实践任务**:

1. 实现简单的聊天室（Pub/Sub）
2. 实现WebSocket集群消息广播（参考websocket_cluster.go）
3. 理解Pub/Sub的优缺点（消息丢失、持久化问题）

**关键知识点**:

- Pub/Sub是"fire and forget"，不保证消息送达
- 适合实时通知，不适合重要消息传递
- 与消息队列（如Kafka）的区别

---

### 3.3 服务发现模式

**学习目标**: 理解如何用Redis实现服务发现

**学习内容**:

- 服务注册：SET + EXPIRE（心跳机制）
- 服务发现：GET/KEYS/SCAN
- 心跳续期：定期更新过期时间
- 服务下线：DEL操作

**项目对应代码**:

- `pkg/discovery/discovery.go` - 完整的服务发现实现
  - `Register()`: 注册节点
  - `updateNode()`: 更新节点信息（带TTL）
  - `GetActiveNodes()`: 获取活跃节点
  - `StartHeartbeat()`: 心跳机制

**实践任务**:

1. 阅读discovery.go，理解服务发现的实现
2. 实现一个简单的服务注册中心
3. 理解TTL在服务发现中的作用
4. 实现服务健康检查

**关键知识点**:

- TTL实现自动下线
- 心跳机制保证服务可用性
- 服务发现的CAP理论权衡

---

### 3.4 布隆过滤器

**学习目标**: 理解布隆过滤器的原理和应用

**学习内容**:

- 布隆过滤器原理：位数组 + 多个哈希函数
- 优点：空间效率高，查询快
- 缺点：有误判率，不能删除
- Redis实现布隆过滤器

**项目对应代码**:

- `pkg/filter/filter.go` - 布隆过滤器封装
  - `AddAddress()`: 添加地址
  - `ContainsAddress()`: 检查地址是否存在

**实践任务**:

1. 理解布隆过滤器的数学原理
2. 实现一个简单的布隆过滤器
3. 用布隆过滤器实现缓存穿透防护
4. 理解误判率和位数组大小的关系

---

## 第四阶段：实战应用（4-5周）

### 4.1 缓存模式

**学习目标**: 掌握常见的缓存模式

**学习内容**:

- Cache-Aside模式（项目使用）
- Read-Through/Write-Through
- Write-Behind
- 缓存更新策略

**项目对应代码**:

- `model/redismodel/token.go:14-70` - Cache-Aside模式

  1. 先查Redis缓存
  2. 缓存未命中，查数据库
  3. 将结果写入缓存

**实践任务**:

1. 实现Cache-Aside模式（参考token.go）
2. 处理缓存穿透（布隆过滤器或空值缓存）
3. 处理缓存雪崩（随机过期时间）
4. 处理缓存击穿（分布式锁）

---

### 4.2 时间序列数据存储

**学习目标**: 用Sorted Set存储和查询时间序列数据

**学习内容**:

- 时间戳作为分数
- ZRANGEBYSCORE查询时间范围
- 数据压缩和存储优化
- 数据过期策略

**项目对应代码**:

- `model/redismodel/kline.go` - K线数据存储
  - 使用ZRANGEBYSCORE查询24小时数据
  - 自定义压缩格式存储数据

**实践任务**:

1. 实现K线数据存储（参考kline.go）
2. 实现时间范围聚合查询
3. 实现数据压缩（减少存储空间）
4. 实现数据自动过期清理

---

### 4.3 WebSocket订阅管理

**学习目标**: 用Redis管理WebSocket订阅关系

**学习内容**:

- Hash存储订阅关系（channel -> clients）
- Set存储用户客户端关系
- Pub/Sub实现跨节点消息广播
- 订阅关系的清理和过期

**项目对应代码**:

- `apps/websocket/internal/svc/redischannel.go` - 完整的订阅管理
- `apps/websocket/internal/svc/websocketcontext.go` - 客户端管理

**实践任务**:

1. 实现WebSocket订阅系统（参考redischannel.go）
2. 实现跨节点消息广播
3. 实现订阅关系的持久化
4. 处理客户端断线重连

---

## 第五阶段：性能优化和最佳实践（5-6周）

### 5.1 Redis性能优化

**学习内容**:

- Pipeline批量操作
- Lua脚本减少网络往返
- 连接池配置
- 内存优化

**项目实践**:

- 查看项目中的Redis连接配置
- 分析哪些操作可以用Pipeline优化
- 理解go-zero框架的Redis封装

---

### 5.2 Redis高可用

**学习内容**:

- Redis主从复制
- Redis Sentinel
- Redis Cluster
- 数据持久化（RDB/AOF）

**项目实践**:

- 查看docker-compose.yml的Redis配置
- 理解单机Redis的局限性
- 规划Redis集群方案

---

### 5.3 监控和运维

**学习内容**:

- Redis监控指标（内存、连接数、命令统计）
- 慢查询分析
- 内存分析
- 故障排查

**实践任务**:

1. 使用redis-cli监控Redis
2. 分析慢查询日志
3. 内存使用分析
4. 性能压测

---

## 学习资源推荐

### 官方文档

- Redis官方文档：https://redis.io/docs/
- Redis命令参考：https://redis.io/commands/

### 项目代码学习路径

1. **基础操作**: `model/redismodel/token.go` → `pkg/oklink/client.go`
2. **数据结构**: `apps/websocket/internal/svc/redischannel.go` → `model/redismodel/kline.go`
3. **高级功能**: `pkg/xredis/lock.go` → `pkg/discovery/discovery.go` → `pkg/filter/filter.go`
4. **实战应用**: `apps/websocket/internal/svc/websocket_cluster.go`

### 实践建议

1. **边学边做**: 每学一个功能，就在项目中找到对应的代码
2. **修改测试**: 尝试修改项目代码，理解每个参数的作用
3. **独立实现**: 不看代码，自己实现类似功能，然后对比
4. **问题驱动**: 遇到问题先查文档，再看项目代码如何解决

---

## 学习检查清单

### 基础阶段

- [ ] 能够独立部署Redis
- [ ] 掌握String的所有操作（SET/GET/SETEX/SETNX/INCR）
- [ ] 理解过期时间的作用
- [ ] 掌握Hash的基本操作（HSET/HGET/HGETALL/HDEL）

### 进阶阶段

- [ ] 掌握Set操作（SADD/SREM/SMEMBERS）
- [ ] 掌握Sorted Set操作（ZADD/ZRANGE/ZRANGEBYSCORE）
- [ ] 能够用Sorted Set实现排行榜
- [ ] 能够用Sorted Set存储时间序列数据

### 高级阶段

- [ ] 理解并实现分布式锁
- [ ] 理解Pub/Sub的使用场景和限制
- [ ] 能够实现简单的服务发现
- [ ] 理解布隆过滤器的原理

### 实战阶段

- [ ] 能够实现Cache-Aside缓存模式
- [ ] 能够处理缓存穿透、雪崩、击穿
- [ ] 能够用Redis管理WebSocket订阅
- [ ] 能够优化Redis性能

---

## 常见问题解答

**Q: 为什么项目用Hash而不是String存储JSON？**

A: Hash可以单独更新字段，不需要序列化整个对象，更节省内存。

**Q: SETNX和SET NX有什么区别？**

A: SETNX是旧命令，SET NX是新语法，功能相同，但SET更灵活。

**Q: Pub/Sub会丢失消息吗？**

A: 会。如果订阅者不在线，消息会丢失。需要可靠消息传递时用消息队列。

**Q: 分布式锁的过期时间设置多少合适？**

A: 根据业务操作时间设置，通常设置为操作时间的2-3倍，并实现续期机制。

---

## 下一步学习方向

完成本路线图后，可以深入学习：

1. **Redis Cluster**: 集群模式和分片
2. **Redis Streams**: 消息队列功能
3. **Redis Modules**: 扩展功能（如RediSearch、RedisJSON）
4. **Redis源码**: 深入理解Redis实现原理
