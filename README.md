# Lin - Go Web 框架 (V10.0)

基于 `net/http` 原生库封装的轻量级 Web 框架，提供洋葱模型中间件、安全防御、限流控制及异步任务调度能力。

## 🚀 核心特性

- **洋葱模型中间件**：支持优先级控制的中间件链路（Panic 恢复 → 链路追踪 → 日志 → 限流 → 签名校验 → 鉴权）
- **安全防御**：基于 HMAC-SHA256 的接口签名验证 + Nonce 内存存储防重放攻击
- **限流机制**：本地令牌桶限流
- **异步任务调度**：Worker Pool 消费模型，Dispatcher 分发，内存队列缓冲
- **链路追溯**：基于 Sonyflake 生成全局唯一 TraceID
- **优雅启停**：支持 Context 超时控制与服务平滑退出

## 🏗️ 架构概览

```
┌─────────────────────────────────────────────────┐
│                    HTTP Server                   │
│  ┌───────────────────────────────────────────┐  │
│  │ Recover → Trace → Log → Limit → Sign → Auth │  │
│  └───────────────────────────────────────────┘  │
│                       │                          │
│                       ▼                          │
│              ┌──────────────┐                   │
│              │   Handler    │                   │
│              └──────────────┘                   │
│                       │                          │
│                       ▼                          │
│  ┌─────────────────────────────────────────┐    │
│  │         Dispatcher + Worker Pool         │    │
│  │   (MemoryQueue → Jobs Channel → Worker)  │    │
│  └─────────────────────────────────────────┘    │
└─────────────────────────────────────────────────┘
```

## 🛠️ 核心组件

| 模块 | 说明 |
|------|------|
| `Lin.App` | 路由管理 + 优先级中间件链 + HTTP Server |
| `Lin.Chain` | 洋葱模型中间件组装函数 |
| `middlewares.Recover` | Panic 恢复，避免服务崩溃 |
| `middlewares.Trace` | 生成/传递 Sonyflake TraceID |
| `middlewares.Log` | 请求方法、路径、状态码、耗时、TraceID 日志 |
| `middlewares.Auth` | Bearer Token / X-AUTH-TOKEN 认证 |
| `middlewares.LocalLimit` | 令牌桶本地限流 |
| `middlewares.Sign` | HMAC-SHA256 签名验证 + Nonce 防重放 |
| `ratelimit.TokenBucket` | 线程安全令牌桶实现 |
| `security` | 签名验证、Nonce 内存存储、时间窗口校验 |
| `queue.MemoryQueue` | 基于 Channel 的内存队列 |
| `works.Dispatcher` | 从队列拉取任务分发到工作通道 |
| `works.WorkerPool` | 并发 Worker 处理任务 |
| `pkg/errors` | 统一错误类型（6 种分类）及 HTTP 状态映射 |
| `pkg/response` | 标准化 JSON 响应 `{code, msg, data}` |
| `snowid` | Sonyflake 分布式 ID 生成 |

## ⚙️ 优先级配置

| 中间件 | 优先级 | 说明 |
|--------|--------|------|
| RecoverMiddleware | 0 | Panic 恢复（最高优先级） |
| TraceMiddleware | 10 | 链路追踪 ID 注入 |
| LogMiddleware | 20 | 请求日志记录 |
| SignMiddleware | 用户自定义 | 签名验证（建议 5） |
| LocalLimitMiddleware | 用户自定义 | 限流（建议 15） |
| AuthMiddleware | 用户自定义 | 鉴权（建议 25） |

## 🔐 签名规范

受保护接口需在 Header 中携带：

| Header | 说明 | 校验规则 |
|--------|------|----------|
| `X-TimeStamp` | Unix 时间戳（秒） | 必须在时间窗口内 |
| `X-Nonce` | 随机字符串 | 长度 8-128 位，防重放 |
| `X-Sign` | HMAC-SHA256 签名 | 与 Payload 校验一致 |

**签名 Payload 构造顺序**：`Method + Path + TimeStamp + Nonce + Body`

## 📦 依赖

| 依赖 | 用途 |
|------|------|
| `github.com/sony/sonyflake` | 全局唯一 ID（TraceID / TaskID） |
