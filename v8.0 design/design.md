# 分布式限流系统设计

## 核心架构
- 多实例部署，共享Redis实现全局限流
- 雪花算法生成全局唯一ID（MachineID区分实例）
- Worker Pool + Channel 异步处理任务

## 限流算法
滑动窗口计数器（Redis Lua原子操作）
- Key: `ratelimit:{path}:{window_bucket}`
- 窗口内超过阈值则返回429

## 数据流
Submit → Redis队列 → Worker消费 → MySQL持久化 → 状态缓存