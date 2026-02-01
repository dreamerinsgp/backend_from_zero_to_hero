# Kafka 学习路线图 - 基于实际项目实践

> 本路线图基于 `fun_dex_v2` 项目中的实际 Kafka 使用场景，确保所学知识可直接应用于生产环境。

## 📋 项目中的 Kafka 使用概览

### 使用的 Kafka 客户端库
- **Go 语言**: 
  - `github.com/IBM/sarama` (v1.x) - 用于 Producer（同步和异步）
  - `github.com/segmentio/kafka-go` (v0.4.47) - 用于 Consumer

### 实际应用场景
1. **Producer（生产者）**: 发送交易数据到 Kafka Topic
   - 同步 Producer (`sarama.SyncProducer`)
   - 异步 Producer (`sarama.AsyncProducer`)
   
2. **Consumer（消费者）**: 消费交易数据并处理
   - 使用 `kafka-go` 库进行消息消费
   - 批量处理交易数据
   - 生成 K 线数据并存储到数据库

3. **安全配置**:
   - SASL/PLAIN 认证
   - TLS 加密连接
   - 支持 Google Cloud Managed Kafka

---

## 🎯 学习路线图（从基础到高级）

### 第一阶段：Kafka 基础概念（1-2周）

#### 1.1 核心概念理解
- [ ] **Topic（主题）**: 消息的分类
  - 项目中使用的 Topic: `sol-trades`, `eth-trades`, `sol-pair-price-change`
  - 理解 Topic 的分区和副本机制

- [ ] **Partition（分区）**: 消息的物理存储单元
  - 查看: `key.Partition` 用于追踪消息来源
  - 理解分区如何实现并行处理

- [ ] **Producer（生产者）**: 发送消息的客户端
  - 项目中实现位置: `apps/consumer/internal/logic/mq/producer.go`
  - 项目中实现位置: `apps/market/internal/mqs/producer/producer.go`

- [ ] **Consumer（消费者）**: 消费消息的客户端
  - 项目中实现位置: `apps/market/internal/mqs/consumers/trade_consumer.go`
  - 项目中实现位置: `apps/dataflow/internal/mqs/consumers/trade_consumer.go`

- [ ] **Consumer Group（消费者组）**: 实现负载均衡和容错
  - 项目中配置: `Group: data-flow-default-group10`

- [ ] **Offset（偏移量）**: 消息在分区中的位置
  - 项目中追踪: `key.Offset` 用于日志记录

#### 1.2 实践任务
```bash
# 1. 安装 Kafka（使用 Docker）
docker run -p 9092:9092 apache/kafka:latest

# 2. 创建 Topic（参考项目中的 Topic）
kafka-topics.sh --create --topic sol-trades --bootstrap-server localhost:9092

# 3. 发送测试消息
kafka-console-producer.sh --topic sol-trades --bootstrap-server localhost:9092

# 4. 消费测试消息
kafka-console-consumer.sh --topic sol-trades --bootstrap-server localhost:9092 --from-beginning
```

---

### 第二阶段：Go 语言 Kafka 客户端（2-3周）

#### 2.1 Sarama 库学习（Producer）

**学习重点**（参考 `apps/consumer/internal/logic/mq/producer.go`）:

- [ ] **配置 Producer**
  ```go
  config := sarama.NewConfig()
  config.Producer.Timeout = time.Second
  config.Producer.MaxMessageBytes = 1024 * 1024 * 10 // 10MB
  config.Producer.Partitioner = sarama.NewHashPartitioner
  config.Producer.Retry.Max = 3
  ```

- [ ] **同步 Producer**
  ```go
  producer, err := sarama.NewSyncProducer(brokers, config)
  message := &sarama.ProducerMessage{
      Topic: topic,
      Key:   sarama.StringEncoder(key),
      Value: sarama.ByteEncoder(data),
  }
  partition, offset, err := producer.SendMessage(message)
  ```

- [ ] **异步 Producer**
  ```go
  producer, err := sarama.NewAsyncProducer(brokers, config)
  producer.Input() <- &sarama.ProducerMessage{
      Topic: topic,
      Key:   sarama.StringEncoder(key),
      Value: sarama.ByteEncoder(data),
  }
  // 处理错误
  go func() {
      for err := range producer.Errors() {
          log.Errorf("Kafka producer error: %v", err)
      }
  }()
  ```

**实践任务**:
- [ ] 实现一个简单的同步 Producer，发送 JSON 格式的交易数据
- [ ] 实现一个异步 Producer，处理错误和成功回调
- [ ] 对比同步和异步 Producer 的性能差异

#### 2.2 Kafka-Go 库学习（Consumer）

**学习重点**（参考 `apps/market/internal/mqs/consumers/trade_consumer.go`）:

- [ ] **Consumer 配置**
  ```go
  reader := kafka.NewReader(kafka.ReaderConfig{
      Brokers:  brokers,
      Topic:    topic,
      GroupID:  groupID,
      MinBytes: 10e3, // 10KB
      MaxBytes: 10e6, // 10MB
  })
  ```

- [ ] **消息消费**
  ```go
  message, err := reader.ReadMessage(ctx)
  // 处理消息
  err := json.Unmarshal(message.Value, &tradeMsg)
  ```

- [ ] **批量消费和并发处理**
  ```go
  // 参考项目中的实现
  var wg sync.WaitGroup
  workerPool.Submit(func() {
      defer wg.Done()
      // 处理消息
  })
  ```

**实践任务**:
- [ ] 实现一个 Consumer，消费交易数据并打印
- [ ] 实现批量消费（参考项目中的批量处理逻辑）
- [ ] 实现 Worker Pool 并发处理消息（使用 `ants` 库）

---

### 第三阶段：项目中的高级特性（2-3周）

#### 3.1 安全配置（SASL + TLS）

**学习重点**（参考 `apps/consumer/internal/logic/mq/producer.go`）:

- [ ] **SASL/PLAIN 认证**
  ```go
  config.Net.SASL.Enable = true
  config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
  config.Net.SASL.User = username
  config.Net.SASL.Password = password
  ```

- [ ] **TLS 加密**
  ```go
  config.Net.TLS.Enable = true
  // 可选: 加载 CA 证书
  config.Net.TLS.Config = &tls.Config{
      RootCAs: caCertPool,
  }
  ```

**实践任务**:
- [ ] 配置本地 Kafka 使用 SASL/PLAIN
- [ ] 配置 TLS 连接（使用自签名证书）

#### 3.2 消息序列化与反序列化

**学习重点**（参考项目中的实现）:

- [ ] **JSON 序列化**
  ```go
  // Producer 端
  data, _ := json.Marshal(tradeList)
  SendEventLogKafkaInfoMessage(topic, key, data)
  
  // Consumer 端
  var tradeMsg []*model.TradeWithPair
  json.Unmarshal(key.Value, &tradeMsg)
  ```

- [ ] **消息 Key 的使用**
  - 项目中使用: `fmt.Sprintf("%v", slot)` 作为 Key
  - Key 用于分区路由（HashPartitioner）

**实践任务**:
- [ ] 实现自定义消息格式的序列化/反序列化
- [ ] 理解 Key 如何影响消息分区

#### 3.3 错误处理和重试机制

**学习重点**:

- [ ] **Producer 重试**
  ```go
  config.Producer.Retry.Max = 3
  config.Producer.Retry.Backoff = 100 * time.Millisecond
  config.Producer.Return.Errors = true
  ```

- [ ] **Consumer 错误处理**
  ```go
  if err := json.Unmarshal(key.Value, &tradeMsg); err != nil {
      logc.Errorf(ctx, "failed to unmarshal: %+v", err)
      return err
  }
  ```

**实践任务**:
- [ ] 实现 Producer 错误重试逻辑
- [ ] 实现 Consumer 死信队列（处理失败的消息）

#### 3.4 性能优化

**学习重点**（参考项目中的优化）:

- [ ] **批量处理**
  - 项目中: 批量处理交易数据生成 K 线
  - 使用 `sync.Pool` 减少内存分配

- [ ] **并发处理**
  - 使用 `ants` Worker Pool（项目中: `ants.NewPool(30)`）
  - 使用 `sync.WaitGroup` 等待所有任务完成

- [ ] **消息大小限制**
  ```go
  config.Producer.MaxMessageBytes = 1024 * 1024 * 10 // 10MB
  ```

**实践任务**:
- [ ] 对比单线程 vs 并发处理的性能
- [ ] 实现消息压缩（如果消息很大）

---

### 第四阶段：生产环境最佳实践（2-3周）

#### 4.1 监控和日志

**学习重点**（参考项目中的实现）:

- [ ] **日志记录**
  ```go
  logx.Infof("[kafka] send event log to kafka success: %v:%v:%v", 
      topic, partition, offset)
  logx.Infof("kafka key offset: %d, partition: %d, time: %v", 
      key.Offset, key.Partition, msg.CreateTime)
  ```

- [ ] **Prometheus 指标**
  ```go
  // 项目中使用的指标
  prometheus.NewCounter(prometheus.CounterOpts{
      Name: "dataflow_kafka_consumer_fetch",
      Help: "dataflow kafka consumer fetch number",
  })
  ```

**实践任务**:
- [ ] 添加 Producer 发送速率指标
- [ ] 添加 Consumer 消费延迟指标
- [ ] 集成 Grafana 可视化

#### 4.2 消息顺序和幂等性

**学习重点**:

- [ ] **消息顺序保证**
  - 使用相同的 Key 确保消息发送到同一分区
  - Consumer 按分区顺序消费

- [ ] **幂等性处理**
  - 项目中: 检查 `TokenPriceUSD == 0` 跳过无效消息
  - 实现去重逻辑（基于交易 Hash）

**实践任务**:
- [ ] 实现基于消息 Key 的顺序保证
- [ ] 实现消息去重机制

#### 4.3 容错和高可用

**学习重点**:

- [ ] **Consumer Group 的 Rebalance**
  - 理解 Consumer Group 如何实现负载均衡
  - 处理 Consumer 故障时的 Rebalance

- [ ] **Offset 管理**
  - 项目中配置: `Offset: last`（从最新消息开始）
  - 理解 `earliest` vs `last` 的区别

**实践任务**:
- [ ] 实现手动提交 Offset
- [ ] 实现 Consumer 优雅关闭
- [ ] 测试 Consumer 故障恢复

---

### 第五阶段：项目实战（2-3周）

#### 5.1 理解项目架构

**任务清单**:
- [ ] 阅读 `apps/consumer/internal/logic/mq/producer.go`
- [ ] 阅读 `apps/market/internal/mqs/consumers/trade_consumer.go`
- [ ] 阅读 `apps/dataflow/internal/mqs/consumers/trade_consumer.go`
- [ ] 理解消息流转: Consumer → Kafka → Market/Dataflow

#### 5.2 实现新功能

**实战任务**:
- [ ] **任务 1**: 添加一个新的 Topic `sol-blocks`，发送区块数据
  - 实现 Producer 发送区块数据
  - 实现 Consumer 消费并存储到数据库

- [ ] **任务 2**: 优化现有 Consumer 性能
  - 增加批量大小
  - 优化 Worker Pool 大小
  - 添加性能指标

- [ ] **任务 3**: 实现消息重试机制
  - 失败消息发送到重试 Topic
  - 实现指数退避重试
  - 达到最大重试次数后发送到死信队列

#### 5.3 测试和调试

**任务清单**:
- [ ] 编写单元测试（参考 `producer_test.go`, `consumer_test.go`）
- [ ] 使用 Docker Compose 搭建本地 Kafka 环境
- [ ] 使用 `kafka-console-consumer.sh` 调试消息
- [ ] 使用 `kafka-consumer-groups.sh` 查看 Consumer Group 状态

---

## 📚 推荐学习资源

### 官方文档
- [Apache Kafka 官方文档](https://kafka.apache.org/documentation/)
- [Sarama Go 客户端文档](https://github.com/IBM/sarama)
- [Kafka-Go 文档](https://github.com/segmentio/kafka-go)

### 项目中的关键文件
1. **Producer 实现**:
   - `/root/fun_dex_from_zero_to_hero/dex_full/apps/consumer/internal/logic/mq/producer.go`
   - `/root/fun_dex_from_zero_to_hero/dex_full/apps/market/internal/mqs/producer/producer.go`

2. **Consumer 实现**:
   - `/root/fun_dex_from_zero_to_hero/dex_full/apps/market/internal/mqs/consumers/trade_consumer.go`
   - `/root/fun_dex_from_zero_to_hero/dex_full/apps/dataflow/internal/mqs/consumers/trade_consumer.go`

3. **配置文件**:
   - `/root/fun_dex_from_zero_to_hero/dex_full/apps/consumer/etc/consumer.yaml`
   - `/root/fun_dex_from_zero_to_hero/dex_full/apps/market/etc/market.yaml`

### 实践环境搭建

```yaml
# docker-compose.yml (参考项目中的 docker-compose.yml)
version: '3.8'
services:
  zookeeper:
    image: confluentinc/cp-zookeeper:latest
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      
  kafka:
    image: confluentinc/cp-kafka:latest
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://localhost:9092
```

---

## ✅ 学习检查清单

完成以下任务后，你已经掌握了项目中使用的 Kafka 知识：

- [ ] 能够独立搭建 Kafka 环境
- [ ] 理解 Topic、Partition、Consumer Group 等核心概念
- [ ] 能够使用 Sarama 实现同步和异步 Producer
- [ ] 能够使用 Kafka-Go 实现 Consumer
- [ ] 理解并实现 SASL/TLS 安全配置
- [ ] 能够处理消息序列化/反序列化
- [ ] 理解错误处理和重试机制
- [ ] 能够优化 Producer/Consumer 性能
- [ ] 能够监控和调试 Kafka 应用
- [ ] 理解消息顺序和幂等性
- [ ] 能够阅读并理解项目中的 Kafka 代码

---

## 🎓 进阶学习方向

完成基础学习后，可以深入学习：

1. **Kafka Streams**: 流式处理框架
2. **Kafka Connect**: 数据集成工具
3. **Schema Registry**: 消息 Schema 管理（Avro/Protobuf）
4. **Kafka 集群管理**: 多 Broker 配置、副本机制
5. **性能调优**: 吞吐量优化、延迟优化
6. **云服务**: Google Cloud Managed Kafka、Confluent Cloud

---

## 💡 常见问题解答

### Q: 项目中为什么使用两个不同的 Kafka 客户端库？
**A**: 
- `sarama` 用于 Producer，功能更全面，支持同步和异步模式
- `kafka-go` 用于 Consumer，API 更简洁，适合简单的消费场景

### Q: 如何选择同步还是异步 Producer？
**A**: 
- **同步 Producer**: 需要确认消息发送成功，适合关键业务数据（项目中用于交易数据）
- **异步 Producer**: 追求高吞吐量，可以容忍少量消息丢失（项目中用于 K 线数据）

### Q: Consumer Group 的作用是什么？
**A**: 
- 实现负载均衡：多个 Consumer 实例共同消费一个 Topic
- 实现容错：Consumer 故障时，其他 Consumer 接管其分区
- 项目中: `Group: data-flow-default-group10` 确保多个服务实例不会重复消费

---

**最后更新**: 基于 `/root/fun_dex_from_zero_to_hero/dex_full` 项目分析
**预计学习时间**: 8-12 周（根据个人基础调整）

