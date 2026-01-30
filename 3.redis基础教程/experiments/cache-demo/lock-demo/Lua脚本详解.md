# go-zero RedisLock Lua 脚本详解

## 概述

go-zero 的 RedisLock 使用了两个 Lua 脚本来保证原子性操作：
1. **lockScript** - 获取锁（支持可重入）
2. **delScript** - 释放锁（安全释放）

## 脚本1：lockScript（获取锁）

### 源码

```lua
if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end
```

### 参数说明

- **KEYS[1]**：锁的key（如 `"lock:stock:product:1001"`）
- **ARGV[1]**：锁的唯一标识符（`rl.id`，16位随机字符串）
- **ARGV[2]**：过期时间（毫秒）= `seconds * 1000 + tolerance`

### 逻辑解析

#### 情况1：锁已存在且是当前进程持有（可重入）

```lua
if redis.call("GET", KEYS[1]) == ARGV[1] then
    -- 锁已存在，且值等于当前进程的id
    -- 说明：当前进程已经持有这个锁（可重入场景）
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    -- 更新过期时间（续期）
    return "OK"
```

**场景**：
- 进程A已经获取了锁
- 进程A的同一个goroutine或另一个goroutine再次尝试获取同一个锁
- 由于值匹配，直接续期，返回成功

**示例**：
```go
lock := redis.NewRedisLock(redis, "lock:test")
lock.SetExpire(10)

// 第一次获取
ok1, _ := lock.AcquireCtx(ctx)  // true

// 第二次获取（可重入）
ok2, _ := lock.AcquireCtx(ctx)  // true（续期，不阻塞）
```

#### 情况2：锁不存在或锁被其他进程持有

```lua
else
    -- 锁不存在，或者锁的值不等于当前进程的id
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
    -- 尝试设置锁（仅当key不存在时）
end
```

**子情况2.1：锁不存在**
- `GET` 返回 `nil`
- `SET NX` 成功，返回 `"OK"`
- 获取锁成功

**子情况2.2：锁被其他进程持有**
- `GET` 返回其他进程的id（如 `"abc123"`）
- `SET NX` 失败（因为key已存在），返回 `nil`
- 获取锁失败

### 执行流程图

```
开始
  ↓
GET lock:key
  ↓
值 == 当前进程id?
  ├─ 是 → SET lock:key value PX timeout (续期)
  │        ↓
  │      返回 "OK" (成功)
  │
  └─ 否 → SET lock:key value NX PX timeout
           ↓
          key不存在?
          ├─ 是 → 返回 "OK" (成功)
          └─ 否 → 返回 nil (失败，锁被其他进程持有)
```

### 关键特性

1. **可重入性**：
   - 同一个进程可以多次获取同一个锁
   - 通过检查值是否匹配来判断

2. **原子性**：
   - 整个操作在Lua脚本中原子执行
   - 避免了竞态条件

3. **自动续期**：
   - 如果锁已存在且是当前进程持有，自动续期
   - 防止锁在业务执行过程中过期

## 脚本2：delScript（释放锁）

### 源码

```lua
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end
```

### 参数说明

- **KEYS[1]**：锁的key
- **ARGV[1]**：锁的唯一标识符（`rl.id`）

### 逻辑解析

#### 情况1：锁存在且是当前进程持有

```lua
if redis.call("GET", KEYS[1]) == ARGV[1] then
    -- 锁存在，且值等于当前进程的id
    return redis.call("DEL", KEYS[1])
    -- 删除锁，返回 1（成功）
```

**场景**：
- 进程A持有锁
- 进程A释放锁
- 值匹配，删除成功

#### 情况2：锁不存在或锁被其他进程持有

```lua
else
    -- 锁不存在，或者锁的值不等于当前进程的id
    return 0
    -- 不删除，返回 0（失败）
end
```

**子情况2.1：锁不存在**
- `GET` 返回 `nil`
- 不删除，返回 `0`
- 可能锁已过期或被其他进程释放

**子情况2.2：锁被其他进程持有**
- `GET` 返回其他进程的id
- 不删除，返回 `0`
- **防止误删其他进程的锁！**

### 执行流程图

```
开始
  ↓
GET lock:key
  ↓
值 == 当前进程id?
  ├─ 是 → DEL lock:key
  │        ↓
  │      返回 1 (成功)
  │
  └─ 否 → 返回 0 (失败，不删除)
```

### 关键特性

1. **安全性**：
   - 只有锁的持有者才能释放锁
   - 防止误删其他进程的锁

2. **原子性**：
   - 检查和删除在Lua脚本中原子执行
   - 避免了竞态条件

3. **幂等性**：
   - 如果锁不存在，返回0（不报错）
   - 可以安全地多次调用

## 为什么需要 Lua 脚本？

### 问题：不使用 Lua 脚本的问题

#### 错误示例1：直接 SET

```go
// ❌ 错误：可能覆盖其他进程的锁
redis.Set("lock:key", "value", 10)
```

**问题**：
- 如果锁已存在，会覆盖其他进程的锁
- 没有使用 NX，无法保证互斥性

#### 错误示例2：GET + SET（非原子）

```go
// ❌ 错误：非原子操作，存在竞态条件
value := redis.Get("lock:key")
if value == myId {
    redis.Del("lock:key")  // 可能在这之间，锁被其他进程获取了！
}
```

**问题**：
- GET 和 DEL 之间有时间间隔
- 在这期间，锁可能被其他进程获取
- 导致误删其他进程的锁

**时序问题**：
```
T1: 进程A GET lock:key → "abc123" (进程A的id)
T2: 进程A检查：value == myId → true
T3: 进程A的锁过期了！
T4: 进程B获取锁：SET lock:key "def456" NX → 成功
T5: 进程A执行 DEL lock:key → 删除了进程B的锁！❌
```

### 解决方案：使用 Lua 脚本

Lua 脚本在 Redis 中**原子执行**：
- 所有操作在一个事务中完成
- 不会被其他命令打断
- 保证了操作的原子性

## 实际使用示例

### 示例1：基本使用

```go
lock := redis.NewRedisLock(redis, "lock:order:123")
lock.SetExpire(30)

// 获取锁
ok, err := lock.AcquireCtx(ctx)
if !ok {
    return fmt.Errorf("获取锁失败")
}
defer lock.ReleaseCtx(ctx)  // 确保释放锁

// 执行业务逻辑
doSomething()
```

**执行过程**：
1. `AcquireCtx()` 调用 `lockScript`
2. 如果锁不存在，`SET NX` 成功，返回 `"OK"`
3. 业务逻辑执行
4. `ReleaseCtx()` 调用 `delScript`
5. 检查值匹配，`DEL` 成功，返回 `1`

### 示例2：可重入场景

```go
lock := redis.NewRedisLock(redis, "lock:test")
lock.SetExpire(10)

// 第一次获取
ok1, _ := lock.AcquireCtx(ctx)  // true

// 第二次获取（同一个锁对象）
ok2, _ := lock.AcquireCtx(ctx)  // true（续期）

// 释放
lock.ReleaseCtx(ctx)  // 删除锁
lock.ReleaseCtx(ctx)  // 返回 false（锁已不存在）
```

**执行过程**：
1. 第一次 `AcquireCtx()`：
   - `GET lock:test` → `nil`
   - `SET lock:test "abc123" NX PX 10000` → `"OK"`
   - 返回 `true`

2. 第二次 `AcquireCtx()`：
   - `GET lock:test` → `"abc123"`（当前进程的id）
   - 值匹配，执行 `SET lock:test "abc123" PX 10000`（续期）
   - 返回 `"OK"`，返回 `true`

3. 第一次 `ReleaseCtx()`：
   - `GET lock:test` → `"abc123"`
   - 值匹配，`DEL lock:test` → `1`
   - 返回 `true`

4. 第二次 `ReleaseCtx()`：
   - `GET lock:test` → `nil`（已删除）
   - 值不匹配，返回 `0`
   - 返回 `false`

### 示例3：多进程竞争

```go
// 进程A
lockA := redis.NewRedisLock(redis, "lock:test")
lockA.SetExpire(10)
okA, _ := lockA.AcquireCtx(ctx)  // true（获取成功）

// 进程B（同时运行）
lockB := redis.NewRedisLock(redis, "lock:test")
lockB.SetExpire(10)
okB, _ := lockB.AcquireCtx(ctx)  // false（获取失败，锁被进程A持有）
```

**执行过程**：
1. 进程A：`SET lock:test "abc123" NX PX 10000` → `"OK"`
2. 进程B：`SET lock:test "def456" NX PX 10000` → `nil`（失败，key已存在）

## 总结

### lockScript（获取锁）

| 情况 | GET结果 | 值匹配？ | 操作 | 返回值 |
|------|---------|----------|------|--------|
| 锁不存在 | `nil` | - | `SET NX` | `"OK"`（成功） |
| 锁存在，当前进程持有 | `"abc123"` | ✅ 是 | `SET PX`（续期） | `"OK"`（成功） |
| 锁存在，其他进程持有 | `"def456"` | ❌ 否 | `SET NX`（失败） | `nil`（失败） |

### delScript（释放锁）

| 情况 | GET结果 | 值匹配？ | 操作 | 返回值 |
|------|---------|----------|------|--------|
| 锁不存在 | `nil` | - | 不删除 | `0`（失败） |
| 锁存在，当前进程持有 | `"abc123"` | ✅ 是 | `DEL` | `1`（成功） |
| 锁存在，其他进程持有 | `"def456"` | ❌ 否 | 不删除 | `0`（失败） |

### 关键要点

1. ✅ **原子性**：Lua脚本保证操作的原子性
2. ✅ **安全性**：只有锁的持有者才能释放锁
3. ✅ **可重入性**：同一个进程可以多次获取同一个锁
4. ✅ **自动续期**：可重入时自动续期，防止过期
5. ✅ **幂等性**：释放锁可以安全地多次调用
