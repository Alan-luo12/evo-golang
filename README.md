# Go 基于net/http原生库上层封装的web框架 (V9.0)

一个基于 Go 构建的高性能、高并发分布式异步任务处理 API 服务。采用典型的 Clean Architecture，内置完整的安全防御机制（防重放签名、分布式限流）、异步 Worker 池调度，以及平滑重启能力。

## 🚀 核心特性

- **自定义高性能路由**：支持洋葱模型中间件链路，优先级严格控制（`Panic 恢复 → 链路追踪 TraceID → 接口日志 → 限流 → 签名校验 → 鉴权`）。
- **高可用安全防御**：
  - 基于 `HMAC-SHA256` 的严格接口签名验证 (`X-Sign`, `X-TimeStamp`, `X-Nonce`)。
  - 基于 `Redis SETNX` 的分布式 Nonce 防重放攻击拦截。
- **双模限流机制**：支持本地令牌桶（Local）和基于 Redis Lua 脚本的分布式限流（Dist），支持故障降级 (`FailOpen`)。
- **异步任务调度**：使用 Worker Pool 消费模型，请求立即返回雪花 ID，任务通过 Redis 队列异步下发，MySQL 最终落盘，解耦高并发写入瓶颈。
- **分布式追溯**：使用 Sonyflake 生成全局唯一的 Task ID 与 Trace ID。
- **优雅启停**：监听系统信号，支持平滑退出 (`Graceful Shutdown`)，确保 Worker 池中的任务安全执行完毕。
- **性能剖析**：内置 pprof，支持 CPU、内存、Goroutine 实时分析。

## 🏗️ 系统架构

```text
┌─────────────────────────────────────────────────────────────┐
│                      API Gateway / Server                    │
│  ┌─────────────────────────────────────────────────────┐    │
│  │  1. Panic Recover: 异常拦截防崩溃                   │    │
│  │  2. Trace / Log: 注入 TraceID 及耗时记录            │    │
│  │  3. Limiter: Redis 分布式限流 / 本地令牌桶          │    │
│  │  4. Security: HMAC-SHA256 签名 + Redis 防重放       │    │
│  └─────────────────────────────────────────────────────┘    │
│                          │                                   │
│                          ▼                                   │
│  ┌──────────┐      ┌──────────┐      ┌──────────┐          │
│  │ Handler  │ ───► │ Service  │ ───► │   Repo   │          │
│  └──────────┘      └──────────┘      └──────────┘          │
│                          │                                   │
│                          ▼                                   │
│              [Redis Queue] + [MySQL DB]                      │
│                          │                                   │
│                          ▼                                   │
│             [Async Worker Pool Processing]                   │
└─────────────────────────────────────────────────────────────┘
```

## 🛠️ 核心 API

| 接口 | 方法 | 说明 | 挂载中间件 |
| --- | --- | --- | --- |
| `/Submit` | POST | 提交异步任务 | 限流、签名防重放、鉴权 |
| `/Getstatus` | GET | 查询任务状态（支持缓存防击穿） | 限流、鉴权 |
| `/HealthHandler` | GET | 服务健康检查 | 基础中间件 |
| `/EchoRequestHandler`| POST | 测试序列化吞吐及 Panic 捕获 | 基础中间件 |
| `/SlowHandler` | GET | 测试慢请求与高延迟阻塞处理 | 基础中间件 |
| `/debug/pprof/*` | GET | 性能分析（CPU/内存/Goroutine） | 无 |

## ⚙️ 快速启动

配置全量基于环境变量驱动，以下为 PowerShell 测试环境启动示例：

```powershell
# 1. 基础及组件配置
$env:PORT='8080'
$env:MACHINEID='1'
$env:REDISADDR='localhost:6379'
$env:MYSQLPATH='root:123456@tcp(localhost:3306)/data?charset=utf8mb4&parseTime=True&loc=Local'

# 2. 安全与防重放配置
$env:SIGNSECRET='test-secret'
$env:SIGNWINDOWSEC='60'

# 3. 限流配置
$env:MODEL='dist'                 # local(令牌桶) 或 dist(分布式)
$env:DISTLIMITMAX='100'           # 窗口期内最大请求数
$env:DISTLIMITWINDOWMS='1000'     # 窗口时间 (毫秒)
$env:DISTLIMITFAILOPEN='false'    # Redis 不可用时是否放行

# 4. Worker 协程池与并发度配置
$env:WORKERPOOLSIZE='10'
$env:JOBQUEUESIZE='100'
$env:PROCESSCONCURRENCY='5'

# 5. pprof 性能分析配置
$env:PPROF_ENABLED='true'         # 启用 pprof
$env:PPROF_ADDR='localhost:6060'  # pprof 监听地址

# 5. 启动服务
go run main.go
```

## 📊 性能剖析 (pprof)

服务内置 pprof，用于实时性能观测。

### 访问地址

启动服务后访问：`http://localhost:6060/debug/pprof/`

### 常用命令

```bash
# CPU 分析（采样 30 秒）
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# 内存分析
go tool pprof http://localhost:6060/debug/pprof/heap

# Web UI 火焰图
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile?seconds=30
```
```

## 🔐 客户端调用规范 (以 `/Submit` 为例)

访问受安全保护的接口时，客户端需在 Headers 中提供正确的签名：

1. 生成当前 Unix 时间戳 `X-TimeStamp` 和随机防重放字符串 `X-Nonce`。
2. 构造待签名 Payload: `Method + Path + TimeStamp + Nonce + Body`。
3. 使用约定的 Secret 计算 HMAC-SHA256 哈希，转为小写 Hex 字符串，作为 `X-Sign` 传入。

**正确请求示例：**
```http
POST /Submit HTTP/1.1
Host: localhost:8080
Content-Type: application/json
X-TimeStamp: 1778673487
X-Nonce: 613ad7b37507fd887919347fa27afac9
X-Sign: 2928294129da27a107a1110d4b312a21e045f40ae59697f26f3c8b3fef094d6e

{"name":"real_task","delay_time":5000}
```
