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