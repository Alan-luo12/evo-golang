## Verison 1.0

## version 1.5

## Version 2.0

## Version 3.0

1. 健康检查
powershell
iwr http://localhost:8080/HealthHandler -UseBasicParsing
2. 正常 Echo 接口
powershell
iwr http://localhost:8080/EchoRequestHandler -Method POST -ContentType "application/json" -Body '{"msg":"hello","panic":false}' -UseBasicParsing
3. 触发 panic
powershell
iwr http://localhost:8080/EchoRequestHandler -Method POST -ContentType "application/json" -Body '{"msg":"test","panic":true}' -UseBasicParsing
4. 慢接口
powershell
iwr "http://localhost:8080/SlowHandler?ms=200" -UseBasicParsing
5. 提交任务
iwr http://localhost:8080/SubmitTaskHandler -Method POST -ContentType "application/json" -Body '{"Name":"testTask","Delay_time":500}' -UseBasicParsing
6. 查询任务状态
powershell
iwr "http://localhost:8080/GetTaskStatusHandler?id=1" -UseBasicParsing

>一切正常

## Version 3.5

  和version3.0保持一致


## Verison 4.0

1. 检查健康状态 (HealthHandler)
(iwr -Uri http://localhost:8080/HealthHandler).Content

2. 提交 Echo 请求 (EchoRequestHandler)
(iwr -Uri http://localhost:8080/EchoRequestHandler -Method Post -Body '{"message": "Hello Server", "panic": false}' -ContentType "application/json").Content


3. 测试慢响应接口 (SlowHandler - 模拟 500ms 延迟)
(iwr -Uri "http://localhost:8080/SlowHandler?ms=500").Content


4. 提交异步任务 (Submit)
(iwr -Uri http://localhost:8080/Submit -Method Post -Body '{"name": "测试任务", "delay_time": 2000}' -ContentType "application/json").Content

5. 查询任务状态 (Getstatus)
(iwr -Uri "http://localhost:8080/Getstatus?id=1").Content

>>>全部通过


## Version5.0

 (iwr -Uri http://localhost:8080/Submit -Method Post -Body '{"name": "测试任务", "delay_time": 2000}' -ContentType "application/json").Content

  {"code":0,"msg":"Success","data":{"task_id":335954282765352961,"status":"submitted"}}                                       


(iwr -Uri "http://localhost:8080/Getstatus?id=335954282765352961").Content

{"code":0,"msg":"Success","data":{"task_id":335954282765352961,"status":"done"}}


curl.exe -X GET http://localhost:8080/HealthHandler

{"code":0,"msg":"Success","data":{"time":"2026-05-06T23:28:17+08:00"}}



 curl.exe -X GET "http://localhost:8080/SlowHandler?ms=150"

{"code":0,"msg":"Success","data":{"delay_time":150}}


时序测试
使用脚本
go\测试用例\时序测试.ps1
>>全部通过

## Version6.0

链路完整测试通过

1. 提交任务，保存 task_id
$resp = iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"name":"smoke","delay_time":1000}' -UseBasicParsing
$taskId = ($resp.Content | ConvertFrom-Json).data.task_id
Write-Host "Task ID: $taskId"

2. 立即查询（应返回 queued 或 running）
iwr "http://localhost:8080/Getstatus?id=$taskId" -UseBasicParsing

3. 等待 2 秒后查询（应返回 done）
Start-Sleep -Seconds 2
iwr "http://localhost:8080/Getstatus?id=$taskId" -UseBasicParsing


## Verison7.0

  6.0中的链路完整通过


## Verison 8.0
链路完整测试通过

1. 提交任务，保存 task_id
$resp = iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"name":"smoke","delay_time":1000}' -UseBasicParsing
$taskId = ($resp.Content | ConvertFrom-Json).data.task_id
Write-Host "Task ID: $taskId"

2. 立即查询（应返回 queued 或 running）
iwr "http://localhost:8080/Getstatus?id=$taskId" -UseBasicParsing

3. 等待 2 秒后查询（应返回 done）
Start-Sleep -Seconds 2
iwr "http://localhost:8080/Getstatus?id=$taskId" -UseBasicParsing

通过

# 健康检查
curl.exe -s http://localhost:8080/HealthHandler
# 预期: {"code":0,"msg":"Success","data":{"time":"..."}}

# Echo JSON
iwr -Uri http://localhost:8080/EchoRequestHandler -Method Post -ContentType "application/json" -Body '{"message":"hello","panic":false}'
# 200: {"code":0,"msg":"Success","data":"hello"}

# 慢接口
curl.exe "http://localhost:8080/SlowHandler?ms=50"
# 200, 延迟约 50ms

# 404
curl.exe http://localhost:8080/xxxapi
# 404

# Panic 恢复
iwr -Uri http://localhost:8080/EchoRequestHandler -Method Post -ContentType "application/json" -Body '{"message":"crash","panic":true}'
# 500, 服务不挂


通过


## Verison 9.0
iwr有bug传输过程会丢失双引号

PS C:\Users\罗宇轩\Desktop\go> # ==================== 定义函数 ====================
PS C:\Users\罗宇轩\Desktop\go> function Get-UnixTimestamp { 
>>     $epoch = Get-Date -Year 1970 -Month 1 -Day 1 -Hour 0 -Minute 0 -Second 0
>>     [int64]([DateTime]::UtcNow - $epoch).TotalSeconds
>> }
PS C:\Users\罗宇轩\Desktop\go> function New-Nonce {
>>     $bytes = [byte[]]::new(16)
>>     [Security.Cryptography.RNGCryptoServiceProvider]::Create().GetBytes($bytes)
>>     -join ($bytes | ForEach-Object { $_.ToString("x2") })
>> }
PS C:\Users\罗宇轩\Desktop\go> function Sign-HMAC256 {
>>     param($method, $path, $timestamp, $nonce, $body, $secret="test-secret")
>>     $payload = "$method$path$timestamp$nonce$body"
>>     $hmac = New-Object System.Security.Cryptography.HMACSHA256
>>     $hmac.Key = [Text.Encoding]::UTF8.GetBytes($secret)
>>     $hash = $hmac.ComputeHash([Text.Encoding]::UTF8.GetBytes($payload))
>>     -join ($hash | ForEach-Object { $_.ToString("x2") })
>> }
PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # ==================== 生成签名 ====================
PS C:\Users\罗宇轩\Desktop\go> $ts = Get-UnixTimestamp
PS C:\Users\罗宇轩\Desktop\go> $nonce = New-Nonce
PS C:\Users\罗宇轩\Desktop\go> $body = '{"name":"t","delay_time":100}'
PS C:\Users\罗宇轩\Desktop\go> $sign = Sign-HMAC256 "POST" "/Submit" $ts $nonce $body
PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> Write-Host "Timestamp: $ts"
Timestamp: 1778673487
PS C:\Users\罗宇轩\Desktop\go> Write-Host "Nonce: $nonce"
Nonce: 613ad7b37507fd887919347fa27afac9
PS C:\Users\罗宇轩\Desktop\go> Write-Host "Sign: $sign"                                                    Sign: 2928294129da27a107a1110d4b312a21e045f40ae59697f26f3c8b3fef094d6e                                     PS C:\Users\罗宇轩\Desktop\go>                                                                             PS C:\Users\罗宇轩\Desktop\go> # ==================== 使用 Invoke-WebRequest 发送（保留引号）====================
PS C:\Users\罗宇轩\Desktop\go> $headers = @{
>>     "Content-Type" = "application/json"
>>     "X-Sign" = $sign
>>     "X-TimeStamp" = "$ts"
>>     "X-Nonce" = $nonce
>> }
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> try {
>>     $resp = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers $headers -Body $body -UseBasicParsing
>>     Write-Host "StatusCode: $($resp.StatusCode)" -ForegroundColor Green
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
PS C:\Users\罗宇轩\Desktop\go>                                                                             PS C:\Users\罗宇轩\Desktop\go> # ==================== 使用 Invoke-WebRequest 发送（保留引号）====================
PS C:\Users\罗宇轩\Desktop\go> $headers = @{
>>     "Content-Type" = "application/json"
>>     "X-Sign" = $sign
>>     "X-TimeStamp" = "$ts"
>>     "X-Nonce" = $nonce
>> }
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> try {
>>     $resp = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers $headers -Body $body -UseBasicParsing
>>     Write-Host "StatusCode: $($resp.StatusCode)" -ForegroundColor Green
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
PS C:\Users\罗宇轩\Desktop\go> # ==================== 使用 Invoke-WebRequest 发送（保留引号）====================
PS C:\Users\罗宇轩\Desktop\go> $headers = @{
>>     "Content-Type" = "application/json"
>>     "X-Sign" = $sign
>>     "X-TimeStamp" = "$ts"
>>     "X-Nonce" = $nonce
>> }
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> try {
>>     $resp = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers $headers -Body $body -UseBasicParsing
>>     Write-Host "StatusCode: $($resp.StatusCode)" -ForegroundColor Green
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
======
PS C:\Users\罗宇轩\Desktop\go> $headers = @{
>>     "Content-Type" = "application/json"
>>     "X-Sign" = $sign
>>     "X-TimeStamp" = "$ts"
>>     "X-Nonce" = $nonce
>> }
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> try {
>>     $resp = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers $headers -Body $body -UseBasicParsing
>>     Write-Host "StatusCode: $($resp.StatusCode)" -ForegroundColor Green
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
>>     "Content-Type" = "application/json"
>>     "X-Sign" = $sign
>>     "X-TimeStamp" = "$ts"
>>     "X-Nonce" = $nonce
>> }
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> try {
>>     $resp = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers $headers -Body $body -UseBasicParsing
>>     Write-Host "StatusCode: $($resp.StatusCode)" -ForegroundColor Green
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
>>     "X-Nonce" = $nonce
>> }
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> try {
>>     $resp = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers $headers -Body $body -UseBasicParsing
>>     Write-Host "StatusCode: $($resp.StatusCode)" -ForegroundColor Green
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
PS C:\Users\罗宇轩\Desktop\go> try {
>>     $resp = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers $headers -Body $body -UseBasicParsing
>>     Write-Host "StatusCode: $($resp.StatusCode)" -ForegroundColor Green
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
>>     Write-Host "StatusCode: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
>>     $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
>>     $errorBody = $reader.ReadToEnd()
>>     Write-Host "Response: $($resp.Content)" -ForegroundColor Green
>> } catch {
>>     Write-Host "StatusCode: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
>>     $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
>>     $errorBody = $reader.ReadToEnd()
>>     Write-Host "StatusCode: $($_.Exception.Response.StatusCode.value__)" -ForegroundColor Red
>>     $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
>>     $errorBody = $reader.ReadToEnd()
>>     $errorBody = $reader.ReadToEnd()
>>     Write-Host "Error Response: $errorBody" -ForegroundColor Red
>> }
StatusCode: 200
>>     Write-Host "Error Response: $errorBody" -ForegroundColor Red
>> }
StatusCode: 200
Response: {"code":0,"msg":"Success","data":{"task_id":336948049681121281,"status":"submitted"}}

StatusCode: 200
Response: {"code":0,"msg":"Success","data":{"task_id":336948049681121281,"status":"submitted"}}

Response: {"code":0,"msg":"Success","data":{"task_id":336948049681121281,"status":"submitted"}}

PS C:\Users\罗宇轩\Desktop\go> Invoke-WebRequest -Uri http://localhost:8080/Submit -Method POST -Headers @{"Content-Type"="application/json";"X-TimeStamp"="123";"X-Nonce"="abc"} -Body '{}' -UseBasicParsing -ErrorAction SilentlyContinue | Select-Object -ExpandProperty StatusCode                                           Invoke-WebRequest : {"code":4091,"msg":"Sign, TimeStamp, or Nonce is empty","data":null}                   所在位置 行:1 字符: 1                                                                                      + Invoke-WebRequest -Uri http://localhost:8080/Submit -Method POST -Hea ...                                
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebReq 
   uest]，WebException                                                                                         + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebReques     tCommand                                                                                                PS C:\Users\罗宇轩\Desktop\go> Invoke-WebRequest -Uri http://localhost:8080/Submit -Method POST -Headers @{"Content-Type"="application/json";"X-Sign"="wrong";"X-TimeStamp"="123";"X-Nonce"="abc"} -Body '{}' -UseBasicParsing -ErrorAction SilentlyContinue | Select-Object -ExpandProperty StatusCode
Invoke-WebRequest : {"code":4093,"msg":"Nonce is invalid","data":null}
所在位置 行:1 字符: 1
+ Invoke-WebRequest -Uri http://localhost:8080/Submit -Method POST -Hea ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebReq 
   uest]，WebException                                                                                         + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebReques     tCommand                                                                                                PS C:\Users\罗宇轩\Desktop\go> (Invoke-WebRequest -Uri http://localhost:8080/HealthHandler -UseBasicParsing).StatusCode
200
PS C:\Users\罗宇轩\Desktop\go> # 第一次 - 应成功
PS C:\Users\罗宇轩\Desktop\go> $ts=Get-UnixTimestamp; $n=New-Nonce; $b='{"name":"replay","delay_time":100}'; $s=Sign-HMAC256 "POST" "/Submit" $ts $n $b                                                               PS C:\Users\罗宇轩\Desktop\go> Invoke-WebRequest -Uri http://localhost:8080/Submit -Method POST -Headers @{"Content-Type"="application/json";"X-Sign"=$s;"X-TimeStamp"="$ts";"X-Nonce"=$n} -Body $b -UseBasicParsing > $null                                                                                                     
PS C:\Users\罗宇轩\Desktop\go> # 第二次（相同 nonce） - 应 409
PS C:\Users\罗宇轩\Desktop\go> try { Invoke-WebRequest -Uri http://localhost:8080/Submit -Method POST -Headers @{"Content-Type"="application/json";"X-Sign"=$s;"X-TimeStamp"="$ts";"X-Nonce"=$n} -Body $b -UseBasicParsing } catch { $_.Exception.Response.StatusCode.value__ }                                                  409                                                                                                        PS C:\Users\罗宇轩\Desktop\go> # 提交任务                                                                  PS C:\Users\罗宇轩\Desktop\go> $ts=Get-UnixTimestamp; $n=New-Nonce; $b='{"name":"e2e","delay_time":2000}'; $s=Sign-HMAC256 "POST" "/Submit" $ts $n $b
PS C:\Users\罗宇轩\Desktop\go> $resp = Invoke-WebRequest -Uri http://localhost:8080/Submit -Method POST -Headers @{"Content-Type"="application/json";"X-Sign"=$s;"X-TimeStamp"="$ts";"X-Nonce"=$n} -Body $b -UseBasicParsing
PS C:\Users\罗宇轩\Desktop\go> $taskId = ($resp.Content | ConvertFrom-Json).data.task_id
PS C:\Users\罗宇轩\Desktop\go> Write-Host "Task ID: $taskId"
Task ID: 336948265671000065
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> # 立即查询（应 queued 或 running）
PS C:\Users\罗宇轩\Desktop\go> Start-Sleep -Seconds 1
PS C:\Users\罗宇轩\Desktop\go> Invoke-WebRequest -Uri "http://localhost:8080/Getstatus?id=$taskId" -UseBasicParsing | Select-Object -ExpandProperty Content
{"code":0,"msg":"Success","data":{"task_id":336948265671000065,"status":"running"}}

PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> # 等待任务完成（delay_time=2000ms）
PS C:\Users\罗宇轩\Desktop\go> Start-Sleep -Seconds 3
PS C:\Users\罗宇轩\Desktop\go> Invoke-WebRequest -Uri "http://localhost:8080/Getstatus?id=$taskId" -UseBasicParsing | Select-Object -ExpandProperty Content
{"code":0,"msg":"Success","data":{"task_id":336948265671000065,"status":"done"}}

PS C:\Users\罗宇轩\Desktop\go> # 注意：固定 nonce 会导致大量 409，仅用于测试系统处理能力
PS C:\Users\罗宇轩\Desktop\go> $ts=Get-UnixTimestamp; $n="fixedbench"; $b='{"name":"bench","delay_time":100}'; $s=Sign-HMAC256 "POST" "/Submit" $ts $n $b
PS C:\Users\罗宇轩\Desktop\go> hey -n 10000 -c 100 -m POST -H "Content-Type: application/json" -H "X-Sign: $s" -H "X-TimeStamp: $ts" -H "X-Nonce: $n" -d "$b" http://localhost:8080/Submit

Summary:
  Total:        1.1135 secs
  Slowest:      0.0884 secs
  Fastest:      0.0003 secs
  Average:      0.0110 secs
  Requests/sec: 8980.5931

  Total data:   548000 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.009 [1361]  |■■■■■■■
  0.018 [8082]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.027 [312]   |■■
  0.036 [145]   |■
  0.044 [27]    |
  0.053 [6]     |
  0.062 [3]     |
  0.071 [15]    |
  0.080 [22]    |
  0.088 [26]    |


Latency distribution:
  10%% in 0.0089 secs
  25%% in 0.0093 secs
  50%% in 0.0097 secs
  75%% in 0.0104 secs
  90%% in 0.0121 secs
  95%% in 0.0198 secs
  99%% in 0.0354 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0299 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0273 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0046 secs
  resp wait:    0.0107 secs, 0.0002 secs, 0.0622 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [409] 400 responses
  [429] 9600 responses



PS C:\Users\罗宇轩\Desktop\go>

全部通过


测试项	预期	实际	状态
正常签名请求	200	{"code":0,...}	✅
缺少 X-Sign	409	409	✅
错误签名	409	409	✅
健康检查（无签名）	200	200	✅
Nonce 重用（第二次）	409	409	✅
提交任务 → 查询状态	200 → running → done	全部成功	✅
压力测试（10000 请求）	系统稳定	QPS ~8980	✅

## pprof 测试入口
http://localhost:6060/debug/pprof/

正常访问pprof服务