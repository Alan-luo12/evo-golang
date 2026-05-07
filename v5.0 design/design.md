## V5 设计决策

### 决策1：Submit 不再写 DB pending，而是只写 Redis queue + Redis status
- 背景：
  V4 中 Submit 在高并发下直接写 SQLite，出现严重 database lock / 500 错误，成功率极低。
- 选项：
  1. Submit 继续写 DB pending
  2. Submit 写 DB + Redis 双写
  3. Submit 只写 Redis，消费时再落 DB
- 选择：
  选择方案 3：Submit 只写 Redis queue + Redis status。
- 理由：
  1. 当前主矛盾是入口并发写冲突，而不是强一致性。
  2. Redis 写入成本低，能显著降低 Submit 延迟。
  3. 通过队列把高并发前台写转换为后台受控消费。
- 风险：
  1. Submit 成功时任务尚未落 DB。
  2. Redis 在 worker 消费前成为短期状态真源。
  3. 若 worker 崩溃或 Redis 故障，可能出现状态丢失或查询不一致。

### 决策2：使用 Redis List（LPUSH / BRPOP）实现最小任务队列
- 背景：
  需要快速引入一个可运行的异步任务队列，替代 V4 的野生 goroutine。
- 选项：
  1. 内存 channel
  2. Redis List
  3. Redis Stream
- 选择：
  Redis List（LPUSH + BRPOP）
- 理由：
  1. 实现简单，学习成本低。
  2. 支持跨进程共享，天然优于单机内存队列。
  3. 便于后续演进到多 worker / 分布式实例。
- 风险：
  1. BRPOP 后消息立即出队，没有 ACK 机制。
  2. worker 在消费后崩溃时，任务可能丢失。

### 决策3：状态缓存与任务队列分离
- 背景：
  任务排队与任务状态查询属于两种不同职责。
- 选择：
  队列使用 `task:queue:normal`，状态使用独立 key。
- 理由：
  1. queue 用于消费，不适合作为查询接口状态源。
  2. cache 用于高频读，能显著降低 DB 查询压力。
- 风险：
  队列和状态缓存分离后，存在双写一致性问题。
