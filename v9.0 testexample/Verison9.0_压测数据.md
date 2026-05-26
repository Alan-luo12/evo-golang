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
PS C:\Users\罗宇轩\Desktop\go> function Invoke-SignedRequest {
>>     param($Uri, $Method="POST", $Body, $ContentType="application/json")
>>     $ts = Get-UnixTimestamp
>>     $nonce = New-Nonce
>>     $sign = Sign-HMAC256 $Method "/Submit" $ts $nonce $Body
>>     $headers = @{"Content-Type"=$ContentType;"X-Sign"=$sign;"X-TimeStamp"="$ts";"X-Nonce"=$nonce}
>>     try { Invoke-WebRequest -Uri $Uri -Method $Method -Headers $headers -Body $Body -UseBasicParsing }
>>     catch { $_.Exception.Response }
>> }
PS C:\Users\罗宇轩\Desktop\go> $ts = Get-UnixTimestamp
PS C:\Users\罗宇轩\Desktop\go> $nonce = "fixedbenchmarknonce123"
PS C:\Users\罗宇轩\Desktop\go> $body = '{"name":"baseline","delay_time":1000}'
PS C:\Users\罗宇轩\Desktop\go> $sign = Sign-HMAC256 "POST" "/Submit" $ts $nonce $body
PS C:\Users\罗宇轩\Desktop\go> Write-Host "X-Sign: $sign"
X-Sign: c4973129e0aade6ef47eefdcedb32629f30c66f29017ba3da0b26db3ec18b9af
PS C:\Users\罗宇轩\Desktop\go> Write-Host "X-TimeStamp: $ts"
X-TimeStamp: 1778675018
PS C:\Users\罗宇轩\Desktop\go> Write-Host "X-Nonce: $nonce"
X-Nonce: fixedbenchmarknonce123
PS C:\Users\罗宇轩\Desktop\go> # 1. 健康检查
PS C:\Users\罗宇轩\Desktop\go> hey -n 10000 -c 100 http://localhost:8080/HealthHandler

Summary:
  Total:        0.5647 secs
  Slowest:      0.0375 secs
  Fastest:      0.0001 secs
  Average:      0.0056 secs
  Requests/sec: 17707.0785

  Total data:   710000 bytes
  Size/request: 71 bytes

Response time histogram:
  0.000 [1]     |
  0.004 [329]   |■
  0.008 [9357]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [194]   |■
  0.015 [18]    |
  0.019 [1]     |
  0.023 [0]     |
  0.026 [2]     |
  0.030 [41]    |
  0.034 [43]    |
  0.037 [14]    |


Latency distribution:
  10%% in 0.0048 secs
  25%% in 0.0050 secs
  50%% in 0.0052 secs
  75%% in 0.0056 secs
  90%% in 0.0063 secs
  95%% in 0.0069 secs
  99%% in 0.0262 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0271 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0262 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0044 secs
  resp wait:    0.0052 secs, 0.0001 secs, 0.0159 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0061 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 2. Echo JSON
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 5000 -c 50 -m POST -H "Content-Type: application/json" -d "{\"message\":\"test\",\"panic\":false}" http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.2725 secs
  Slowest:      0.0286 secs
  Fastest:      0.0001 secs
  Average:      0.0027 secs
  Requests/sec: 18350.7855

  Total data:   205000 bytes
  Size/request: 41 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4712]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [234]   |■■
  0.009 [3]     |
  0.011 [0]     |
  0.014 [0]     |
  0.017 [0]     |
  0.020 [0]     |
  0.023 [8]     |
  0.026 [25]    |
  0.029 [17]    |


Latency distribution:
  10%% in 0.0022 secs
  25%% in 0.0024 secs
  50%% in 0.0025 secs
  75%% in 0.0026 secs
  90%% in 0.0028 secs
  95%% in 0.0030 secs
  99%% in 0.0214 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0236 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0230 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0029 secs
  resp wait:    0.0024 secs, 0.0001 secs, 0.0057 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0046 secs

Status code distribution:
  [200] 5000 responses



PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> # 3. SlowHandler 10ms
PS C:\Users\罗宇轩\Desktop\go> hey -n 10000 -c 100 "http://localhost:8080/SlowHandler?ms=10"

Summary:
  Total:        1.0877 secs
  Slowest:      0.0403 secs
  Fastest:      0.0095 secs
  Average:      0.0108 secs
  Requests/sec: 9193.5678

  Total data:   520000 bytes
  Size/request: 52 bytes

Response time histogram:
  0.010 [1]     |
  0.013 [9876]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.016 [23]    |
  0.019 [0]     |
  0.022 [0]     |
  0.025 [0]     |
  0.028 [0]     |
  0.031 [0]     |
  0.034 [21]    |
  0.037 [32]    |
  0.040 [47]    |


Latency distribution:
  10%% in 0.0100 secs
  25%% in 0.0102 secs
  50%% in 0.0104 secs
  75%% in 0.0107 secs
  90%% in 0.0110 secs
  95%% in 0.0113 secs
  99%% in 0.0311 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0274 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0261 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0032 secs
  resp wait:    0.0105 secs, 0.0093 secs, 0.0145 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 4. SlowHandler 200ms
PS C:\Users\罗宇轩\Desktop\go> hey -n 1000 -c 50 "http://localhost:8080/SlowHandler?ms=200"

Summary:
  Total:        4.0385 secs
  Slowest:      0.2267 secs
  Fastest:      0.1998 secs
  Average:      0.2018 secs
  Requests/sec: 247.6191

  Total data:   53000 bytes
  Size/request: 53 bytes

Response time histogram:
  0.200 [1]     |
  0.202 [949]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.205 [0]     |
  0.208 [0]     |
  0.211 [0]     |
  0.213 [0]     |
  0.216 [0]     |
  0.219 [0]     |
  0.221 [0]     |
  0.224 [13]    |■
  0.227 [37]    |■■


Latency distribution:
  10%% in 0.2002 secs
  25%% in 0.2003 secs
  50%% in 0.2006 secs
  75%% in 0.2008 secs
  90%% in 0.2012 secs
  95%% in 0.2228 secs
  99%% in 0.2262 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0012 secs, 0.0000 secs, 0.0255 secs
  DNS-lookup:   0.0011 secs, 0.0000 secs, 0.0234 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:    0.2006 secs, 0.1997 secs, 0.2020 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200] 1000 responses



PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> # 5. Submit (固定签名 - 测试系统吞吐)
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 2000 -c 50 -m POST -H "Content-Type: application/json" -H "X-Sign: $sign" -H "X-TimeStamp: $ts" -H "X-Nonce: $nonce" -d "{\"name\":\"baseline\",\"delay_time\":1000}" http://localhost:8080/Submit

Summary:
  Total:        0.2170 secs
  Slowest:      0.0312 secs
  Fastest:      0.0003 secs
  Average:      0.0053 secs
  Requests/sec: 9218.6333

  Total data:   110000 bytes
  Size/request: 55 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [43]    |■
  0.006 [1888]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [18]    |
  0.013 [0]     |
  0.016 [0]     |
  0.019 [0]     |
  0.022 [0]     |
  0.025 [12]    |
  0.028 [20]    |
  0.031 [18]    |


Latency distribution:
  10%% in 0.0045 secs
  25%% in 0.0046 secs
  50%% in 0.0048 secs
  75%% in 0.0050 secs
  90%% in 0.0053 secs
  95%% in 0.0055 secs
  99%% in 0.0278 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0005 secs, 0.0000 secs, 0.0235 secs
  DNS-lookup:   0.0005 secs, 0.0000 secs, 0.0222 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0036 secs
  resp wait:    0.0047 secs, 0.0002 secs, 0.0077 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [409] 200 responses
  [429] 1800 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 6. 404 路由
PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 http://localhost:8080/xxxapi

Summary:
  Total:        0.2714 secs
  Slowest:      0.0304 secs
  Fastest:      0.0001 secs
  Average:      0.0027 secs
  Requests/sec: 18425.9316

  Total data:   95000 bytes
  Size/request: 19 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4797]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [151]   |■
  0.009 [1]     |
  0.012 [0]     |
  0.015 [0]     |
  0.018 [0]     |
  0.021 [0]     |
  0.024 [26]    |
  0.027 [17]    |
  0.030 [7]     |


Latency distribution:
  10%% in 0.0022 secs
  25%% in 0.0023 secs
  50%% in 0.0024 secs
  75%% in 0.0026 secs
  90%% in 0.0028 secs
  95%% in 0.0030 secs
  99%% in 0.0218 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0232 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0221 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0035 secs
  resp wait:    0.0023 secs, 0.0001 secs, 0.0061 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0044 secs

Status code distribution:
  [404] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 7. Panic 恢复
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 1000 -c 20 -m POST -H "Content-Type: application/json" -d "{\"message\":\"crash\",\"panic\":true}" http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.1176 secs
  Slowest:      0.0252 secs
  Fastest:      0.0001 secs
  Average:      0.0023 secs
  Requests/sec: 8500.9219

  Total data:   56000 bytes
  Size/request: 56 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [887]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [92]    |■■■■
  0.008 [0]     |
  0.010 [0]     |
  0.013 [0]     |
  0.015 [0]     |
  0.018 [0]     |
  0.020 [0]     |
  0.023 [7]     |
  0.025 [13]    |■


Latency distribution:
  10%% in 0.0005 secs
  25%% in 0.0012 secs
  50%% in 0.0022 secs
  75%% in 0.0024 secs
  90%% in 0.0027 secs
  95%% in 0.0031 secs
  99%% in 0.0235 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0004 secs, 0.0000 secs, 0.0219 secs
  DNS-lookup:   0.0004 secs, 0.0000 secs, 0.0201 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0013 secs
  resp wait:    0.0018 secs, 0.0001 secs, 0.0045 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [500] 1000 responses



PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> # 8. 提交真实任务（需要签名）
PS C:\Users\罗宇轩\Desktop\go> $resp = Invoke-SignedRequest -Uri "http://localhost:8080/Submit" -Body '{"name":"real","delay_time":5000}'
PS C:\Users\罗宇轩\Desktop\go> $taskId = ($resp.Content | ConvertFrom-Json).data.task_id
ConvertFrom-Json : 无法将参数绑定到参数“InputObject”，因为该参数是空值。
所在位置 行:1 字符: 28
+ $taskId = ($resp.Content | ConvertFrom-Json).data.task_id
+                            ~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidData: (:) [ConvertFrom-Json]，ParameterBindingValidationException     
    + FullyQualifiedErrorId : ParameterArgumentValidationErrorNullNotAllowed,Microsoft.PowerShell.Command  
   s.ConvertFromJsonCommand

PS C:\Users\罗宇轩\Desktop\go> Write-Host "Task ID: $taskId"
Task ID:
PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> # 9. 查询任务状态
PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=$taskId"

Summary:
  Total:        0.5534 secs
  Slowest:      0.0325 secs
  Fastest:      0.0002 secs
  Average:      0.0055 secs
  Requests/sec: 9035.2226

  Total data:   273200 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [108]   |■
  0.007 [4640]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.010 [128]   |■
  0.013 [22]    |
  0.016 [28]    |
  0.020 [23]    |
  0.023 [0]     |
  0.026 [16]    |
  0.029 [12]    |
  0.033 [22]    |


Latency distribution:
  10%% in 0.0047 secs
  25%% in 0.0049 secs
  50%% in 0.0050 secs
  75%% in 0.0053 secs
  90%% in 0.0059 secs
  95%% in 0.0067 secs
  99%% in 0.0231 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0251 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0233 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0028 secs
  resp wait:    0.0052 secs, 0.0002 secs, 0.0196 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0034 secs

Status code distribution:
  [400] 200 responses
  [429] 4800 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 10. 查询不存在任务
PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=999999999"

Summary:
  Total:        0.5305 secs
  Slowest:      0.0523 secs
  Fastest:      0.0002 secs
  Average:      0.0052 secs
  Requests/sec: 9425.9058

  Total data:   273800 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.005 [4598]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [346]   |■■■
  0.016 [1]     |
  0.021 [0]     |
  0.026 [7]     |
  0.031 [3]     |
  0.037 [0]     |
  0.042 [7]     |
  0.047 [2]     |
  0.052 [35]    |


Latency distribution:
  10%% in 0.0045 secs
  25%% in 0.0046 secs
  50%% in 0.0048 secs
  75%% in 0.0050 secs
  90%% in 0.0053 secs
  95%% in 0.0057 secs
  99%% in 0.0253 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0251 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0232 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0017 secs
  resp wait:    0.0049 secs, 0.0002 secs, 0.0295 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [404] 200 responses
  [429] 4800 responses



PS C:\Users\罗宇轩\Desktop\go> # 并发 100
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 100 -m POST -H "Content-Type: application/json" -H "X-Sign: $sign" -H "X-TimeStamp: $ts" -H "X-Nonce: $nonce" -d "{\"name\":\"grad100\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0068 secs
  Slowest:      0.0458 secs
  Fastest:      0.0002 secs
  Average:      0.0102 secs
  Requests/sec: 9769.6115

  Total data:   16123525 bytes
  Size/request: 55 bytes

Response time histogram:
  0.000 [1]     |
  0.005 [2135]  |
  0.009 [66568] |■■■■■■■■■■■■
  0.014 [218557]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.018 [4334]  |■
  0.023 [775]   |
  0.028 [407]   |
  0.032 [229]   |
  0.037 [107]   |
  0.041 [32]    |
  0.046 [10]    |


Latency distribution:
  10%% in 0.0091 secs
  25%% in 0.0094 secs
  50%% in 0.0098 secs
  75%% in 0.0110 secs
  90%% in 0.0119 secs
  95%% in 0.0124 secs
  99%% in 0.0164 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0259 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0228 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0004 secs
  resp wait:    0.0102 secs, 0.0002 secs, 0.0440 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0012 secs

Status code distribution:
  [409] 3100 responses
  [429] 290055 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 200
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 200 -m POST -H "Content-Type: application/json" -H "X-Sign: $sign" -H "X-TimeStamp: $ts" -H "X-Nonce: $nonce" -d "{\"name\":\"grad200\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0146 secs
  Slowest:      0.1864 secs
  Fastest:      0.0002 secs
  Average:      0.0197 secs
  Requests/sec: 10138.1487

  Total data:   16736060 bytes
  Size/request: 55 bytes

Response time histogram:
  0.000 [1]     |
  0.019 [104775]        |■■■■■■■■■■■■■■■■■■■■■
  0.037 [198472]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.056 [676]   |
  0.075 [157]   |
  0.093 [89]    |
  0.112 [47]    |
  0.131 [26]    |
  0.149 [12]    |
  0.168 [11]    |
  0.186 [26]    |


Latency distribution:
  10%% in 0.0182 secs
  25%% in 0.0186 secs
  50%% in 0.0192 secs
  50%% in 0.0192 secs
  75%% in 0.0200 secs
  90%% in 0.0217 secs
  95%% in 0.0241 secs
  99%% in 0.0310 secs

Details (average, fastest, slowest):
  99%% in 0.0310 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0268 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0251 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0011 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0251 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0011 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:    0.0197 secs, 0.0002 secs, 0.1863 secs
  resp wait:    0.0197 secs, 0.0002 secs, 0.1863 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [409] 3000 responses
  [429] 301292 responses



PS C:\Users\罗宇轩\Desktop\go>
PS C:\Users\罗宇轩\Desktop\go> # 并发 500
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 500 -m POST -H "Content-Type: application/json" -H "X-Sign: $sign" -H "X-TimeStamp: $ts" -H "X-Nonce: $nonce" -d "{\"name\":\"grad500\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0315 secs
  Slowest:      1.4567 secs
  Fastest:      0.0002 secs
  Average:      0.0513 secs
  Requests/sec: 9800.8420
  
  Total data:   15988610 bytes
  Size/request: 55 bytes

Response time histogram:
  0.000 [1]     |
  0.146 [290293]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.292 [25]    |
  0.437 [82]    |
  0.583 [52]    |
  0.728 [7]     |
  0.874 [66]    |
  1.020 [20]    |
  1.165 [36]    |
  1.311 [77]    |
  1.457 [43]    |


Latency distribution:
  10%% in 0.0462 secs
  25%% in 0.0470 secs
  50%% in 0.0483 secs
  75%% in 0.0507 secs
  90%% in 0.0575 secs
  95%% in 0.0668 secs
  99%% in 0.0907 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0686 secs
  DNS-lookup:   0.0004 secs, 0.0000 secs, 0.0652 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0127 secs
  resp wait:    0.0504 secs, 0.0002 secs, 1.4238 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0126 secs

Status code distribution:
  [409] 3000 responses
  [429] 287702 responses

Error distribution:
  [3632]        Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 1000
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 1000 -m POST -H "Content-Type: application/json" -H "X-Sign: $sign" -H "X-TimeStamp: $ts" -H "X-Nonce: $nonce" -d "{\"name\":\"grad1000\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0836 secs
  Slowest:      3.4142 secs
  Fastest:      0.0002 secs
  Average:      0.1150 secs
  Requests/sec: 8993.0286

  Total data:   13937000 bytes
  Size/request: 55 bytes

Response time histogram:
  0.000 [1]     |
  0.342 [252585]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.683 [69]    |
  1.024 [129]   |
  1.366 [102]   |
  1.707 [121]   |
  2.049 [172]   |
  2.390 [71]    |
  2.731 [68]    |
  3.073 [19]    |
  3.414 [63]    |


Latency distribution:
  10%% in 0.0791 secs
  25%% in 0.1005 secs
  50%% in 0.1125 secs
  75%% in 0.1226 secs
  90%% in 0.1346 secs
  95%% in 0.1451 secs
  99%% in 0.1768 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0582 secs
  DNS-lookup:   0.0053 secs, 0.0000 secs, 0.0699 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0190 secs
  resp wait:    0.1040 secs, 0.0002 secs, 3.3741 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0196 secs

Status code distribution:
  [409] 3017 responses
  [429] 250383 responses

Error distribution:
  [17143]       Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> hey --% -n 50000 -c 2000 -m POST -H "Content-Type: application/json" -H "X-Sign: $sign" -H "X-TimeStamp: $ts" -H "X-Nonce: $nonce" -d "{\"name\":\"brutal\",\"delay_time\":1000}" http://localhost:8080/Submit

Summary:
  Total:        6.2056 secs
  Slowest:      4.5172 secs
  Fastest:      0.0003 secs
  Average:      0.2378 secs
  Requests/sec: 8057.2763

  Total data:   2314400 bytes
  Size/request: 55 bytes

Response time histogram:
  0.000 [1]     |
  0.452 [41625] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.904 [90]    |
  1.355 [6]     |
  1.807 [38]    |
  2.259 [109]   |
  2.710 [4]     |
  3.162 [55]    |
  3.614 [18]    |
  4.065 [72]    |
  4.065 [72]    |
  4.517 [62]    |


  4.517 [62]    |



Latency distribution:
  10%% in 0.0978 secs
Latency distribution:
  10%% in 0.0978 secs
  25%% in 0.1853 secs
  50%% in 0.2139 secs
  75%% in 0.2641 secs
  90%% in 0.3212 secs
  95%% in 0.3764 secs
  99%% in 0.4599 secs
  25%% in 0.1853 secs
  50%% in 0.2139 secs
  75%% in 0.2641 secs
  90%% in 0.3212 secs
  95%% in 0.3764 secs
  99%% in 0.4599 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0013 secs, 0.0000 secs, 0.2116 secs
  90%% in 0.3212 secs
  95%% in 0.3764 secs
  99%% in 0.4599 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0013 secs, 0.0000 secs, 0.2116 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0013 secs, 0.0000 secs, 0.2116 secs
  DNS-lookup:   0.0675 secs, 0.0000 secs, 0.2619 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0230 secs
  resp wait:    0.0662 secs, 0.0002 secs, 4.4166 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0410 secs

  resp read:    0.0001 secs, 0.0000 secs, 0.0410 secs

Status code distribution:
  [409] 700 responses
Status code distribution:
  [409] 700 responses
  [409] 700 responses
  [429] 41380 responses
  [429] 41380 responses

Error distribution:
Error distribution:
  [7920]        Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.
  [7920]        Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.
 made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go>










