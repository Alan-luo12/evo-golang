## Version 8.0 分布式任务限制
PS C:\Users\罗宇轩\Desktop\go> cd C:\Users\罗宇轩\Desktop\go\cmd\server
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:PORT='8080'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:MACHINEID='1'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:MODEL='dist'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:DISTLIMITMAX='5'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:DISTLIMITWINDOWMS='2000'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:REDISADDR='localhost:6379'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> go run main.go

PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:PORT='8081'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:MACHINEID='2'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:MODEL='dist'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:DISTLIMITMAX='5'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:DISTLIMITWINDOWMS='2000'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> $env:REDISADDR='localhost:6379'
PS C:\Users\罗宇轩\Desktop\go\cmd\server> go run main.go


PS C:\Users\罗宇轩\Desktop\go> 
>> for ($i=1; $i -le 10; $i++) {
>>     $result = curl.exe -s -X POST "http://localhost:8080/Submit" -H "Content-Type: application/json" -d '{\"name\":\"test\",\"delay_time\":10}'
>>     Write-Host $result
>>     Start-Sleep -Milliseconds 10
>> }
{"code":0,"msg":"Success","data":{"task_id":336792365001932801,"status":"submitted"}}
{"code":0,"msg":"Success","data":{"task_id":336792365052264449,"status":"submitted"}}
{"code":0,"msg":"Success","data":{"task_id":336792365136150529,"status":"submitted"}}
{"code":0,"msg":"Success","data":{"task_id":336792365220036609,"status":"submitted"}}
{"code":0,"msg":"Success","data":{"task_id":336792365287145473,"status":"submitted"}}
{"code":50011,"msg":"Dist Limit Exceeded","data":null}
{"code":50011,"msg":"Dist Limit Exceeded","data":null}
{"code":50011,"msg":"Dist Limit Exceeded","data":null}
{"code":50011,"msg":"Dist Limit Exceeded","data":null}
{"code":50011,"msg":"Dist Limit Exceeded","data":null}
PS C:\Users\罗宇轩\Desktop\go> 

通过
