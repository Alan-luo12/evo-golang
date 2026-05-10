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