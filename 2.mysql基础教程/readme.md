# MySQL应用场景分析与学习路线图

## 一、项目MySQL应用场景分析

### 1. 数据库连接与配置管理

**应用位置：**

- `trade/internal/svc/servicecontext.go` - Trade服务数据库连接
- `market/internal/svc/servicecontext.go` - Market服务数据库连接  
- `consumer/internal/svc/servicecontext.go` - Consumer服务数据库连接
- `dataflow/internal/svc/servicecontext.go` - Dataflow服务数据库连接

**技术要点：**

- GORM ORM框架 + MySQL驱动
- DSN连接字符串配置 (`user:password@tcp(host:port)/dbname?parseTime=true`)
- 连接池配置：MaxOpenConns(500), MaxIdleConns(200), ConnMaxLifetime(5分钟)
- 多服务独立数据库连接实例

### 2. 核心业务表结构设计

#### 2.1 交易订单表 (`trade_order`, `trade_order_log`)

**文件：** `model/sql/trade.sql`

- 订单状态管理（wait/proc/onchain/fail/suc/cancel/timeout）
- 高精度金额存储（decimal(32,18)）
- 复合索引优化查询（uid+status+trade_type, status+chain_id, created_at+uid+status）
- 唯一索引防止重复（order_id+status）

#### 2.2 区块链数据表

**文件：** `model/sql/sol.sql`

- `token` - 代币信息（唯一索引：chain_id+address）
- `pair` - 交易对信息（多字段索引：name, token_address, pump_point等）
- `trade` - 交易记录（唯一索引：hash_id，时间索引：block_time）
- `block` - 区块信息（唯一索引：slot）

#### 2.3 流动性池表

- `raydium_pool` - Raydium池信息
- `clmm_pool_info_v1/v2` - CLMM池信息（唯一索引：pool_state）
- `cpmm_pool_info` - CPMM池信息
- `pump_amm_info` - Pump AMM信息（唯一索引：pool_account）

#### 2.4 账户表

- `sol_account` - SOL账户（唯一索引：address）
- `sol_token_account` - Token账户（复合唯一索引：owner_address+token_account_address）

### 3. 高级MySQL特性应用

#### 3.1 表分片（Sharding）

**应用位置：**

- `dataflow/internal/data/kline_mysql.go` - K线数据按月分表
- `model/solmodel/trademodel.go` - 交易数据按时间分表

**实现方式：**

- 动态表名生成：`trade_kline_{interval}_{month}`
- 自动创建分表（createTableIfNotExists）
- 跨表查询处理（跨月数据合并）

#### 3.2 事务处理

**应用位置：**

- `model/trademodel/tradeordermodel.go` - InsertWithLog使用事务
- `model/trademodel/tradeorderlogmodel.go` - ON DUPLICATE KEY UPDATE

**技术要点：**

- GORM Transaction包装
- 事务内多表操作
- 错误回滚机制

#### 3.3 批量操作优化

**应用位置：**

- `dataflow/internal/data/kline_mysql.go` - SaveKline批量插入（batchSize=50）
- `model/solmodel/trademodel.go` - BatchInsertTrades（batchSize=1024）

**优化策略：**

- CreateInBatches批量插入
- 数据排序预处理
- 分批处理避免内存溢出

#### 3.4 冲突处理（Upsert）

**应用位置：**

- `dataflow/internal/data/kline_mysql.go` - clause.OnConflict实现UPSERT
- `model/trademodel/tradeorderlogmodel.go` - ON DUPLICATE KEY UPDATE

**实现方式：**

- 基于唯一索引的冲突检测
- 冲突时更新指定字段
- 多字段唯一约束

#### 3.5 死锁处理

**应用位置：**

- `dataflow/internal/data/kline_mysql.go` - SaveKlineWithRetry
- `model/solmodel/trademodel.go` - BatchInsertTrades死锁重试

**处理策略：**

- 检测Deadlock错误
- 指数退避重试（50ms, 100ms递增）
- 互斥锁保护关键操作

### 4. 查询优化技术

#### 4.1 索引设计

- **主键索引：** 自增ID
- **唯一索引：** 防止重复数据（chain_id+address, pool_state等）
- **复合索引：** 优化多条件查询（uid+status+trade_type）
- **时间索引：** 优化时间范围查询（block_time, created_at）

#### 4.2 查询模式

- WHERE条件过滤（多条件组合）
- ORDER BY排序（DESC/ASC）
- LIMIT/OFFSET分页
- GROUP BY聚合查询
- 跨表查询（时间分片表）

### 5. 数据迁移与维护

**文件：**

- `scripts/linux/setup-mysql.sh` - 数据库初始化脚本
- `scripts/linux/import-database.sh` - 数据导入脚本
- `model/migration/` - 数据库迁移代码

## 二、MySQL学习路线图

### 阶段一：MySQL基础（1-2周）

#### 1.1 数据库基础概念

- 数据库、表、字段、记录
- 数据类型（INT, VARCHAR, DECIMAL, TIMESTAMP等）
- 字符集与排序规则（UTF8MB4）

#### 1.2 SQL基础语法

- DDL：CREATE TABLE, ALTER TABLE, DROP TABLE
- DML：INSERT, UPDATE, DELETE, SELECT
- 约束：PRIMARY KEY, UNIQUE, NOT NULL, DEFAULT
- 参考项目文件：`model/sql/trade.sql`, `model/sql/sol.sql`

#### 1.3 数据类型深入

- 数值类型：INT, BIGINT, DECIMAL精度控制
- 字符串类型：VARCHAR长度限制，TEXT大文本
- 时间类型：TIMESTAMP, DATETIME区别
- 布尔类型：TINYINT(1)模拟布尔值

**实践任务：**

- 创建trade_order表结构
- 理解decimal(32,18)的精度含义
- 练习基本CRUD操作

### 阶段二：索引与性能优化（2-3周）

#### 2.1 索引基础

- 主键索引（PRIMARY KEY）
- 唯一索引（UNIQUE KEY）
- 普通索引（KEY/INDEX）
- 复合索引设计原则

#### 2.2 索引优化

- EXPLAIN分析查询计划
- 索引选择原则（最左前缀匹配）
- 覆盖索引优化
- 索引对INSERT/UPDATE的影响

**参考项目案例：**

- `trade_order`表的复合索引设计
- `pair`表的多字段索引
- `sol_token_account`的复合唯一索引

**实践任务：**

- 分析项目中的索引设计
- 使用EXPLAIN优化慢查询
- 设计适合业务场景的索引

### 阶段三：高级特性（3-4周）

#### 3.1 事务与并发控制

- ACID特性
- 事务隔离级别
- 锁机制（表锁、行锁）
- 死锁检测与处理

**参考项目案例：**

- `tradeordermodel.go`的事务使用
- `kline_mysql.go`的死锁重试机制

**实践任务：**

- 实现事务包装的批量操作
- 处理并发场景下的死锁

#### 3.2 批量操作优化

- 批量INSERT优化
- 批量UPDATE策略
- 批量DELETE注意事项
- 分批处理大数据量

**参考项目案例：**

- `SaveKline`的批量插入（batchSize=50）
- `BatchInsertTrades`的批量处理（batchSize=1024）

**实践任务：**

- 实现高效的批量插入
- 优化大批量数据导入

#### 3.3 UPSERT操作

- INSERT ... ON DUPLICATE KEY UPDATE
- REPLACE INTO
- GORM的Clause.OnConflict

**参考项目案例：**

- `kline_mysql.go`的OnConflict实现
- `tradeorderlogmodel.go`的ON DUPLICATE KEY UPDATE

**实践任务：**

- 实现基于唯一索引的UPSERT
- 处理冲突更新逻辑

### 阶段四：表设计与分片（2-3周）

#### 4.1 表设计原则

- 范式化设计
- 反范式化优化
- 软删除设计（deleted_at）
- 时间戳字段（created_at, updated_at）

#### 4.2 表分片策略

- 水平分片（按时间、按ID范围）
- 垂直分片（按业务模块）
- 动态表创建
- 跨分片查询处理

**参考项目案例：**

- K线数据按月分表（`trade_kline_{interval}_{month}`）
- 交易数据按时间分表

**实践任务：**

- 设计时间分片表结构
- 实现动态表创建逻辑
- 处理跨分片查询

### 阶段五：连接池与性能调优（1-2周）

#### 5.1 连接池管理

- 连接池参数配置
- MaxOpenConns vs MaxIdleConns
- 连接生命周期管理
- 连接泄漏检测

**参考项目配置：**

- `servicecontext.go`中的连接池设置
- MaxOpenConns(500), MaxIdleConns(200)

#### 5.2 查询性能优化

- 慢查询日志分析
- 查询缓存使用
- 分区表优化
- 读写分离

**实践任务：**

- 配置MySQL慢查询日志
- 分析并优化慢查询
- 测试连接池性能

### 阶段六：ORM框架深入（2周）

#### 6.1 GORM框架

- 模型定义与标签
- 关联关系（HasOne, HasMany, BelongsTo）
- 查询构建器
- 原生SQL执行

**参考项目代码：**

- `model/solmodel/` - 模型定义示例
- `kline_mysql.go` - GORM查询示例

#### 6.2 迁移管理

- 数据库迁移工具
- 版本控制
- 回滚策略

**参考项目：**

- `model/migration/` - 迁移代码示例

**实践任务：**

- 使用GORM实现CRUD操作
- 实现数据库迁移脚本

### 阶段七：生产环境实践（持续）

#### 7.1 监控与运维

- 性能监控指标
- 慢查询分析
- 备份与恢复
- 主从复制

#### 7.2 故障处理

- 死锁问题排查
- 连接数过多处理
- 数据一致性保证
- 灾难恢复

## 三、学习资源推荐

### 官方文档

- MySQL 8.0官方文档
- GORM官方文档

### 项目实战

- 深入研究fun_dex_v2项目的MySQL使用
- 阅读并理解每个表的索引设计
- 分析查询性能优化点

### 实践建议

1. 搭建本地MySQL环境，导入项目SQL文件
2. 阅读项目代码，理解每个MySQL操作场景
3. 尝试优化现有查询，使用EXPLAIN分析
4. 实现类似的分片表逻辑
5. 处理并发场景下的数据一致性

## 四、关键文件清单

### SQL定义文件

- `model/sql/trade.sql` - 交易订单表结构
- `model/sql/sol.sql` - 区块链相关表结构
- `model/sql/pump_amm.sql` - Pump AMM表结构

### 数据库操作代码

- `trade/internal/svc/servicecontext.go` - Trade服务数据库连接
- `market/internal/svc/servicecontext.go` - Market服务数据库连接
- `dataflow/internal/data/kline_mysql.go` - K线数据操作（分片、批量、UPSERT）
- `model/trademodel/tradeordermodel.go` - 订单模型（事务）
- `model/solmodel/trademodel.go` - 交易模型（批量插入、分片）

### 配置与脚本

- `scripts/linux/setup-mysql.sh` - 数据库初始化
- `scripts/linux/import-database.sh` - 数据导入
- `market/etc/market.yaml` - 数据库配置示例