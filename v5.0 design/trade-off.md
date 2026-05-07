## Version 5.0 Trade-off

### 选择
V5 选择 Submit 不写 DB pending，而是仅写：
- Redis queue
- Redis status cache

worker 真正消费任务时，再写 DB running，并最终更新为 done / failed。

### 获得的收益
1. Submit 请求路径更轻，不再受 SQLite 写锁限制。
2. 高并发提交吞吐显著提升。
3. 状态查询命中缓存时性能极高。
4. 从“同步写库”切换为“异步队列消费”，系统更接近真实任务系统。

### 付出的代价
1. Submit 成功不代表任务已落库。
2. worker 消费前，Redis 成为短期状态真源。
3. Redis 丢失或 worker 崩溃时，任务可能查不到或丢失。
4. LPUSH + BRPOP 没有 ACK，消费失败后无法自动重试。
5. 队列与状态缓存分离，存在双写不一致窗口。

### 结论
V5 不是为了追求强一致，而是为了优先解决：
- 入口高并发写冲突
- 状态共享
- 提交吞吐量

该设计适合作为单机任务系统向受控异步系统过渡的中间版本。
