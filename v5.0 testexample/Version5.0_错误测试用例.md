

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
