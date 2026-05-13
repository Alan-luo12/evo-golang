## Version 8.0 分布式任务限制
//这个是分布式任务限制的示例，采用了会话级别的环境变量，每个会话都有自己的配置，互不干扰。
//例如，我们有两个实例1和实例2，每个实例都有自己的配置。
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

## Version 9.0 分布式任务限制
步骤：

方案一：两个 PowerShell 窗口（推荐，可同时运行）
窗口 1 - 实例 1 (端口 8080)
powershell
cd C:\Users\罗宇轩\Desktop\go\cmd\server
$env:PORT='8080'
$env:MACHINEID='1'
$env:SIGNSECRET='test-secret'
$env:SIGNWINDOWSEC='60'
$env:REDISADDR='localhost:6379'
$env:LimitModel='dist'
$env:DISTLIMITMAX='100'
$env:DISTLIMITWINDOWMS='1000'
go run main.go
窗口 2 - 实例 2 (端口 8081)
powershell
cd C:\Users\罗宇轩\Desktop\go\cmd\server
$env:PORT='8081'
$env:MACHINEID='2'
$env:SIGNSECRET='test-secret'
$env:SIGNWINDOWSEC='60'
$env:REDISADDR='localhost:6379'
$env:LimitModel='dist'
$env:DISTLIMITMAX='100'
$env:DISTLIMITWINDOWMS='1000'
go run main.go

测试命令（在第三个 PowerShell 窗口执行）
1. 定义签名函数
powershell
function Get-UnixTimestamp { 
    $epoch = Get-Date -Year 1970 -Month 1 -Day 1 -Hour 0 -Minute 0 -Second 0
    [int64]([DateTime]::UtcNow - $epoch).TotalSeconds
}
function New-Nonce {
    $bytes = [byte[]]::new(16)
    [Security.Cryptography.RNGCryptoServiceProvider]::Create().GetBytes($bytes)
    -join ($bytes | ForEach-Object { $_.ToString("x2") })
}
function Sign-HMAC256 {
    param($method, $path, $timestamp, $nonce, $body, $secret="test-secret")
    $payload = "$method$path$timestamp$nonce$body"
    $hmac = New-Object System.Security.Cryptography.HMACSHA256
    $hmac.Key = [Text.Encoding]::UTF8.GetBytes($secret)
    $hash = $hmac.ComputeHash([Text.Encoding]::UTF8.GetBytes($payload))
    -join ($hash | ForEach-Object { $_.ToString("x2") })
}
2. 测试：跨实例 Nonce 防重放
powershell
# 生成签名
$ts = Get-UnixTimestamp
$nonce = New-Nonce
$body = '{"name":"dist_test","delay_time":100}'
$sign = Sign-HMAC256 "POST" "/Submit" $ts $nonce $body

# 实例1 请求
Write-Host "实例1 (8080):" -ForegroundColor Cyan
$resp1 = Invoke-WebRequest -Uri "http://localhost:8080/Submit" -Method POST -Headers @{
    "Content-Type"="application/json"
    "X-Sign"=$sign
    "X-TimeStamp"="$ts"
    "X-Nonce"=$nonce
} -Body $body -UseBasicParsing -ErrorAction SilentlyContinue
Write-Host "状态码: $($resp1.StatusCode)" -ForegroundColor Green

# 实例2 使用相同 nonce（应失败）
Write-Host "`n实例2 (8081) 相同 nonce:" -ForegroundColor Yellow
try {
    $resp2 = Invoke-WebRequest -Uri "http://localhost:8081/Submit" -Method POST -Headers @{
        "Content-Type"="application/json"
        "X-Sign"=$sign
        "X-TimeStamp"="$ts"
        "X-Nonce"=$nonce
    } -Body $body -UseBasicParsing -ErrorAction Stop
    Write-Host "状态码: $($resp2.StatusCode)" -ForegroundColor Red
} catch {
    Write-Host "状态码: $($_.Exception.Response.StatusCode.value__) ✅ 被拒绝" -ForegroundColor Green
}
3. 测试：不同实例独立 Nonce（应都成功）
powershell
for ($i=1; $i -le 5; $i++) {
    $ts = Get-UnixTimestamp
    $nonce = New-Nonce
    $sign = Sign-HMAC256 "POST" "/Submit" $ts $nonce $body
    $port = if ($i % 2 -eq 0) { "8081" } else { "8080" }
    $resp = Invoke-WebRequest -Uri "http://localhost:$port/Submit" -Method POST -Headers @{
        "Content-Type"="application/json"
        "X-Sign"=$sign
        "X-TimeStamp"="$ts"
        "X-Nonce"=$nonce
    } -Body $body -UseBasicParsing -ErrorAction SilentlyContinue
    Write-Host "实例$port 请求 $i : $($resp.StatusCode)"
    Start-Sleep -Milliseconds 100
}


测试场景	结果	验证点
实例1 (8080) 使用 nonce	✅ 200	首次请求成功
实例2 (8081) 相同 nonce	✅ 409	跨实例 Nonce 防重放生效
交替请求不同 nonce	✅ 全部 200	各自独立正常


