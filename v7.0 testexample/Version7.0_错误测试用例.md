

##VERSION 1.0/1.5/2.0

##VERSION 3.0

iwr http://localhost:8080/HealthHandler -Method Post -UseBasicParsing -ErrorAction SilentlyContinue
iwr http://localhost:8080/HealthHandler -Method Put -UseBasicParsing -ErrorAction SilentlyContinue

iwr http://localhost:8080/EchoRequestHandler -Method Get -UseBasicParsing -ErrorAction SilentlyContinue
$body = "invalid json"
iwr http://localhost:8080/EchoRequestHandler -Method Post -Body $body -ContentType "application/json" -UseBasicParsing -ErrorAction SilentlyContinue
$body = '{"msg":"test panic","panic":true}'
iwr http://localhost:8080/EchoRequestHandler -Method Post -Body $body -ContentType "application/json" -UseBasicParsing -ErrorAction SilentlyContinue

iwr http://localhost:8080/SlowHandler -Method Get -UseBasicParsing -ErrorAction SilentlyContinue
iwr http://localhost:8080/SlowHandler?ms=abc -Method Get -UseBasicParsing -ErrorAction SilentlyContinue
iwr http://localhost:8080/SlowHandler?ms=@ -Method Get -UseBasicParsing -ErrorAction SilentlyContinue
iwr http://localhost:8080/SlowHandler?ms=-100 -Method Get -UseBasicParsing -ErrorAction SilentlyContinue

iwr http://localhost:8080/SubmitTaskHandler -Method Get -UseBasicParsing -ErrorAction SilentlyContinue
$body = '{"delay_time":1000}'
iwr http://localhost:8080/SubmitTaskHandler -Method Post -Body $body -ContentType "application/json" -UseBasicParsing -ErrorAction SilentlyContinue
$body = '{"Name":"test","delay_time":1000,}'
iwr http://localhost:8080/SubmitTaskHandler -Method Post -Body $body -ContentType "application/json" -UseBasicParsing -ErrorAction SilentlyContinue
$body = '{"Name":"test","delay_time":"abc"}'
iwr http://localhost:8080/SubmitTaskHandler -Method Post -Body $body -ContentType "application/json" -UseBasicParsing -ErrorAction SilentlyContinue

iwr http://localhost:8080/GetTaskStatusHandler -Method Post -UseBasicParsing -ErrorAction SilentlyContinue
iwr http://localhost:8080/GetTaskStatusHandler -Method Get -UseBasicParsing -ErrorAction SilentlyContinue
iwr http://localhost:8080/GetTaskStatusHandler?id=abc -Method Get -UseBasicParsing -ErrorAction SilentlyContinue
iwr http://localhost:8080/GetTaskStatusHandler?id=99999 -Method Get -UseBasicParsing -ErrorAction SilentlyContinue

iwr http://localhost:8080/NotFoundHandler -Method Get -UseBasicParsing -ErrorAction SilentlyContinue

总结
测试接口	测试错误场景	服务返回结果	是否符合预期
HealthHandler	POST / PUT 方法错误	200（该接口未限制请求方法）	✅ 正常
EchoRequestHandler	GET 方法错误	405 Method Not Allowed	✅ 符合
EchoRequestHandler	非法 JSON 请求	400 Invalid json	✅ 符合
EchoRequestHandler	主动触发 panic	500 服务器内部错误	✅ 符合
SlowHandler	未传 ms 参数	400 BadRequest	✅ 符合
SlowHandler	ms=abc（非数字）	400 BadRequest	✅ 符合
SlowHandler	ms=-100（负数）	200（未做负数校验）	✅ 正常
SubmitTaskHandler	GET 方法错误	405 Method Not Allowed	✅ 符合
SubmitTaskHandler	缺失必填字段	正常创建（允许缺失）	✅ 正常
SubmitTaskHandler	JSON 格式语法错误	400 解析失败	✅ 符合
SubmitTaskHandler	类型错误（字符串传数字）	400 类型不匹配	✅ 符合
GetTaskStatus	POST 方法错误	400 方法不允许	✅ 符合
GetTaskStatus	未传 id 参数	400 参数解析失败	✅ 符合
GetTaskStatus	id=abc（非数字）	400 参数解析失败	✅ 符合
GetTaskStatus	id=99999（不存在）	404 任务不存在	✅ 符合

## VERSION 3.5 与 3.0 一致



## VERSION 4.0

PS C:\Users\罗宇轩\Desktop\go> curl.exe -s -X POST http://localhost:8080/EchoRequestHandler -H "Content-Type: application/json" -d "{\"message\": \"boom\", \"panic\": true}"
{"code":4007,"msg":"bad request","data":null}

PS C:\Users\罗宇轩\Desktop\go> curl.exe -s -X GET http://localhost:8080/Submit
{"code":4003,"msg":"invalid json","data":null}

PS C:\Users\罗宇轩\Desktop\go> curl.exe -s -X POST http://localhost:8080/Submit -H "Content-Type: application/json" -d "{\"name\": \"bad_json\""
{"code":4004,"msg":"Bad Request","data":null}

PS C:\Users\罗宇轩\Desktop\go> curl.exe -s -X POST http://localhost:8080/Submit -H "Content-Type: application/json" -d "{\"name\": \"\", \"delay_time\": 100}"
{"code":4004,"msg":"Bad Request","data":null}

PS C:\Users\罗宇轩\Desktop\go> curl.exe -s -X GET "http://localhost:8080/Getstatus?id=abc"
{"code":4006,"msg":"Bad Request","data":null}

PS C:\Users\罗宇轩\Desktop\go> curl.exe -s -X GET "http://localhost:8080/Getstatus?id=999999"
{"code":4041,"msg":"task not found","data":null}

通过错误注入测试


## Version 5.0

接口错误测试

PS C:\Users\罗宇轩\Desktop\go> (iwr -Uri http://localhost:8080/EchoRequestHandler -Method Post -Body '{"message":"boom","panic":true}' -ContentType "application/json" -ErrorAction SilentlyContinue).Content
iwr : {"code":5002,"msg":"Manual Panic triggered","data":null}
所在位置 行:1 字符: 2
+ (iwr -Uri http://localhost:8080/EchoRequestHandler -Method Post -Body ...
+  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest]，WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
PS C:\Users\罗宇轩\Desktop\go> curl.exe -X GET http://localhost:8080/EchoRequestHandler
{"code":40051,"msg":"Method Not Allowed","data":null}
PS C:\Users\罗宇轩\Desktop\go> curl.exe -X POST "http://localhost:8080/Getstatus?id=1"
{"code":4005,"msg":"invalid method","data":null}
PS C:\Users\罗宇轩\Desktop\go> curl.exe -X POST http://localhost:8080/Submit -H "Content-Type: application/json" -d "{bad_json: missing_quotes}"
{"code":4004,"msg":"Bad Request","data":null}
PS C:\Users\罗宇轩\Desktop\go> curl.exe -X GET "http://localhost:8080/Getstatus?id=abc123"
{"code":4006,"msg":"Bad Request","data":null}
PS C:\Users\罗宇轩\Desktop\go> curl.exe -X GET "http://localhost:8080/Getstatus?id=-1"
{"code":4002,"msg":"task id invalid","data":null}
PS C:\Users\罗宇轩\Desktop\go> curl.exe -X GET "http://localhost:8080/Getstatus?id=9999999999999999"
{"code":4041,"msg":"task not found","data":null}
PS C:\Users\罗宇轩\Desktop\go> 

redisdown的测试
1.启动前redisdown

    

错误测试通过

## Version 6.0

以下是为 **V6 服务** 定制的**完整错误注入测试**，涵盖 Redis、SQLite、Panic、超时等常见故障，每条操作都有预期结果，可直接复制执行验证。

---

## 错误注入测试清单

### 1. Redis 连接失败（启动时）
**目的**：验证服务启动检测到 Redis 不可用时能否快速失败/退出。

**步骤**：
1. 停止 Redis 服务（或修改 `REDISADDR` 为错误地址）。
2. 启动 V6 服务。

**预期**：
```
[Error] Failed to connect to Redis at localhost:6379
```
服务退出，不继续运行。

---

### 2. Redis 服务在运行中突然宕机

**目的**：验证服务在运行时 Redis 突然断开，是否影响现有 worker 和后续请求。

**步骤**：
1. 正常启动服务（Redis 已启动）。
2. 提交一个任务（`iwr` 正常返回 task_id）。
3. **停止 Redis 服务**。
4. 再次提交任务。

**预期**：
- 新提交请求应快速返回错误（约几秒内）：
```json
{"code":5001,"msg":"failed to enqueue the task"}
```
- 之前已提交的任务应正常被 worker 消费完成（因任务已入队）。
- 查询状态依然可能返回缓存内的值（如果 Redis 完全不可用，返回 500）。

--

### 4. SQLite 写入并发冲突（注入长任务）

**目的**：验证多个 worker 同时写 SQLite 时，由于 `ProcessConcurrency=1`，不会出现 `database locked`。

**步骤**：
```bash
# 并发提交 200 个长任务（每个 sleep 5 秒）
for i in {1..200}; do
  curl -s -X POST http://localhost:8080/Submit -H "Content-Type: application/json" -d '{"name":"long","delay_time":5000}' &
done
wait
```
**预期**：
- 所有请求返回 200（submitted）。
- 日志中 worker 处理顺序且无锁错误。
- SQLite 不会报 `database locked`。

---

### 5. SQLite 磁盘满 / 只读（模拟）

**目的**：验证写入失败时状态回滚到 `failed`。

**步骤**（需要临时修改数据库文件权限）：
1. 停止服务。
2. 将 `tasks.db` 文件设为只读（Windows 使用 `attrib +R tasks.db`）。
3. 启动服务，提交任务。

**预期**：
```json
{"code":0,"msg":"Success","data":{"task_id":...}}  // Submit 仍返回 200
```
但 worker 处理时日志打印 `failed to create task in db`，并且查询状态最终应为 `failed`。

---

### 6. 请求触发 Panic（已有）

**步骤**：
```bash
iwr http://localhost:8080/EchoRequestHandler -Method Post -ContentType "application/json" -Body '{"message":"crash","panic":true}' -UseBasicParsing
```
**预期**：
- 返回 `{"code":5002,"msg":"Manual Panic triggered","data":null}`
- 服务不崩溃，日志有 `[Panic Recovered]`。

---


---

### 8. 优雅关闭时 worker 尚在处理任务

**步骤**：
1. 提交一个延迟 10 秒的任务。
2. 立即按 Ctrl+C 停止服务。
**预期**：
- HTTP 服务先关闭，不再接受新请求。
- Worker 继续执行当前任务（最多等待 10 秒 或 主程序超时 5 秒）。
- 日志输出 `worker pool drained` 或 `worker pool drain timeout`，服务退出。

---

### 9. 限流中间件拒绝请求（429）

**步骤**：
```bash
hey -n 10000 -c 2000 -m POST -d '{"name":"flood"}' http://localhost:8080/Submit
```
**预期**：
- 部分请求返回 `429`，令牌桶算法生效。
- 后端 worker 队列长度保持稳定，不会无限堆积。

---

### 10. 无效请求格式注入

| 场景 | 命令 | 预期 |
|------|------|------|
| 非 JSON body | `curl -X POST -d "not json" http://localhost:8080/Submit` | 4004 Bad Request |
| 缺少必填字段 | `curl -X POST -H "Content-Type: application/json" -d '{"delay_time":100}' http://localhost:8080/Submit` | 4001 task name can not be empty |
| 超出长度 | 发送超大 name（>1MB） | 服务应拒绝或截断，不崩溃 |
| 错误 method | `curl -X GET http://localhost:8080/Submit` | 4003 invalid json（或 method not allowed） |

---


---

## 结论
全部通过     但是  SQLite 只读  部分通过	worker 检测到错误，但状态未更新为 failed（需修复代码）





## Version 7.0


参数校验触发业务错误码
powershell
# 缺失 name 字段
iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"delay_time":100}'
# 预期：400 {"code":4001,"msg":"task name can not be empty"}

# name 为空字符串
iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"name":"","delay_time":100}'
# 预期：400 code:4001

# JSON 格式错误
iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{bad json'
# 预期：400 {"code":4004,"msg":"Bad Request"}

# 查询状态 id 为非数字
iwr -Uri "http://localhost:8080/Getstatus?id=abc"
# 预期：400 {"code":4006,"msg":"Bad Request"}

# 查询状态 id 为负数
iwr -Uri "http://localhost:8080/Getstatus?id=-1"
# 预期：400 {"code":4002,"msg":"task id invalid"}

# 错误 HTTP 方法
iwr -Uri "http://localhost:8080/Getstatus?id=123" -Method Post
# 预期：400 {"code":4005,"msg":"invalid method"}

>>通过

# Echo 接口主动触发 Panic
iwr -Uri http://localhost:8080/EchoRequestHandler -Method Post -ContentType "application/json" -Body '{"message":"crash","panic":true}'
# 预期：500 {"code":5004,"msg":"Internal Server Error"}，服务不退出，日志含 [Panic Recovered]

# 提交一个长延迟任务（如 30 秒）
iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"name":"longtask","delay_time":30000}'
# 立即按 Ctrl+C 停止服务

>通过