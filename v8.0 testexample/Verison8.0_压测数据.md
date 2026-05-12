下面单机测试都是在单机上测试的，没有分布式部署

PS C:\Users\罗宇轩\Desktop\go> # 1. 健康检查
配置
		RateLimitCapacity:   Getenvint64("RATELIMITCAPACITY", 500),
		RateLimitRefillRate: Getenvint64("RATELIMITREFILLRATE", 100),
		WorkerPool:         Getenvint64("WORKERPOOLSIZE", 10),
		JobQueue:           Getenvint64("JOBQUEUESIZE", 100),
		ProcessConcurrency: Getenvint64("PROCESSCONCURRENCY", 1),
    mysql 默认


PS C:\Users\罗宇轩\Desktop\go> hey -n 10000 -c 100 http://localhost:8080/HealthHandler

Summary:
  Total:        0.4167 secs
  Slowest:      0.0299 secs
  Fastest:      0.0001 secs
  Average:      0.0041 secs
  Requests/sec: 23995.3622
  
  Total data:   710000 bytes
  Size/request: 71 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [550]   |■■
  0.006 [9194]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [147]   |■
  0.012 [6]     |
  0.015 [2]     |
  0.018 [0]     |
  0.021 [19]    |
  0.024 [47]    |
  0.027 [22]    |
  0.030 [12]    |


Latency distribution:
  10%% in 0.0036 secs
  25%% in 0.0038 secs
  50%% in 0.0039 secs
  75%% in 0.0042 secs
  90%% in 0.0045 secs
  95%% in 0.0048 secs
  99%% in 0.0185 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0210 secs
  DNS-lookup:   0.0001 secs, 0.0000 secs, 0.0152 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0026 secs
  resp wait:    0.0038 secs, 0.0001 secs, 0.0126 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0036 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 2. Echo JSON（使用 --% 避免 PowerShell 解析花括号）
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 5000 -c 50 -m POST -H "Content-Type: application/json" -d "{\"message\":\"test\",\"panic\":false}" http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.2254 secs
  Slowest:      0.0244 secs
  Fastest:      0.0001 secs
  Average:      0.0022 secs
  Requests/sec: 22183.1503
  
  Total data:   205000 bytes
  Size/request: 41 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4474]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [472]   |■■■■
  0.007 [3]     |
  0.010 [0]     |
  0.012 [0]     |
  0.015 [0]     |
  0.017 [0]     |
  0.020 [8]     |
  0.022 [23]    |
  0.024 [19]    |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0019 secs
  50%% in 0.0020 secs
  75%% in 0.0022 secs
  90%% in 0.0025 secs
  95%% in 0.0029 secs
  99%% in 0.0188 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0194 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0191 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0041 secs
  resp wait:    0.0019 secs, 0.0001 secs, 0.0063 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0029 secs

Status code distribution:
  [200] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 3. SlowHandler 10ms
PS C:\Users\罗宇轩\Desktop\go> hey -n 10000 -c 100 "http://localhost:8080/SlowHandler?ms=10"

Summary:
  Total:        1.0729 secs
  Slowest:      0.0322 secs
  Fastest:      0.0096 secs
  Average:      0.0106 secs
  Requests/sec: 9320.1405
  
  Total data:   520000 bytes
  Size/request: 52 bytes

Response time histogram:
  0.010 [1]     |
  0.012 [9896]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.014 [3]     |
  0.016 [0]     |
  0.019 [0]     |
  0.021 [0]     |
  0.023 [0]     |
  0.025 [0]     |
  0.028 [22]    |
  0.030 [26]    |
  0.032 [52]    |


Latency distribution:
  10%% in 0.0101 secs
  25%% in 0.0103 secs
  50%% in 0.0104 secs
  75%% in 0.0106 secs
  90%% in 0.0108 secs
  95%% in 0.0110 secs
  99%% in 0.0256 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0194 secs
  DNS-lookup:   0.0001 secs, 0.0000 secs, 0.0168 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0038 secs
  resp wait:    0.0104 secs, 0.0095 secs, 0.0140 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 4. SlowHandler 200ms
PS C:\Users\罗宇轩\Desktop\go> hey -n 1000 -c 50 "http://localhost:8080/SlowHandler?ms=200"

Summary:
  Total:        4.0301 secs
  Slowest:      0.2193 secs
  Fastest:      0.1997 secs
  Average:      0.2014 secs
  Requests/sec: 248.1357
  
  Total data:   53000 bytes
  Size/request: 53 bytes

Response time histogram:
  0.200 [1]     |
  0.202 [933]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.204 [16]    |■
  0.206 [0]     |
  0.208 [0]     |
  0.210 [0]     |
  0.211 [0]     |
  0.213 [0]     |
  0.215 [0]     |
  0.217 [9]     |
  0.219 [41]    |■■


Latency distribution:
  10%% in 0.2001 secs
  25%% in 0.2003 secs
  50%% in 0.2005 secs
  75%% in 0.2008 secs
  90%% in 0.2013 secs
  95%% in 0.2165 secs
  99%% in 0.2189 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0008 secs, 0.0000 secs, 0.0178 secs
  DNS-lookup:   0.0007 secs, 0.0000 secs, 0.0166 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:    0.2005 secs, 0.1997 secs, 0.2028 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200] 1000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 5. Submit 限流内（无认证，用于对比）
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 2000 -c 50 -m POST -H "Content-Type: application/json" -d "{\"name\":\"baseline\",\"delay_time\":1000}" http://localhost:8080/Submit

Summary:
  Total:        0.2148 secs
  Slowest:      0.0219 secs
  Fastest:      0.0002 secs
  Average:      0.0053 secs
  Requests/sec: 9310.6835
  
  Total data:   124640 bytes
  Size/request: 62 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [109]   |■■■
  0.005 [415]   |■■■■■■■■■■■■
  0.007 [1359]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [15]    |
  0.011 [0]     |
  0.013 [1]     |
  0.015 [22]    |■
  0.018 [41]    |■
  0.020 [28]    |■
  0.022 [9]     |


Latency distribution:
  10%% in 0.0028 secs
  25%% in 0.0044 secs
  50%% in 0.0051 secs
  75%% in 0.0054 secs
  90%% in 0.0060 secs
  95%% in 0.0147 secs
  99%% in 0.0189 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0003 secs, 0.0000 secs, 0.0170 secs
  DNS-lookup:   0.0003 secs, 0.0000 secs, 0.0149 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0023 secs
  resp wait:    0.0048 secs, 0.0002 secs, 0.0179 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0023 secs

Status code distribution:
  [200] 520 responses
  [429] 1480 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 6. 404 路由
PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 http://localhost:8080/xxxapi

Summary:
  Total:        0.2353 secs
  Slowest:      0.0212 secs
  Fastest:      0.0001 secs
  Average:      0.0023 secs
  Requests/sec: 21251.0312
  
  Total data:   95000 bytes
  Size/request: 19 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [3110]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [1804]  |■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [34]    |
  0.009 [1]     |
  0.011 [0]     |
  0.013 [0]     |
  0.015 [0]     |
  0.017 [9]     |
  0.019 [23]    |
  0.021 [18]    |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0019 secs
  50%% in 0.0021 secs
  75%% in 0.0024 secs
  90%% in 0.0029 secs
  95%% in 0.0034 secs
  99%% in 0.0156 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0170 secs
  DNS-lookup:   0.0001 secs, 0.0000 secs, 0.0153 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0030 secs
  resp wait:    0.0020 secs, 0.0001 secs, 0.0067 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0040 secs

Status code distribution:
  [404] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 7. Panic 恢复
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 1000 -c 20 -m POST -H "Content-Type: application/json" -d "{\"message\":\"crash\",\"panic\":true}" http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.1190 secs
  Slowest:      0.0190 secs
  Fastest:      0.0002 secs
  Average:      0.0023 secs
  Requests/sec: 8405.6570
  
  Total data:   56000 bytes
  Size/request: 56 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [516]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [444]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [19]    |■
  0.008 [0]     |
  0.010 [0]     |
  0.011 [0]     |
  0.013 [0]     |
  0.015 [0]     |
  0.017 [13]    |■
  0.019 [7]     |■


Latency distribution:                                                                                      
  10%% in 0.0004 secs                                                                                      
  25%% in 0.0011 secs                                                                                      
  50%% in 0.0020 secs                                                                                      
  75%% in 0.0030 secs
  90%% in 0.0033 secs
  95%% in 0.0037 secs
  99%% in 0.0164 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0003 secs, 0.0000 secs, 0.0156 secs
  DNS-lookup:   0.0003 secs, 0.0000 secs, 0.0147 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0014 secs
  resp wait:    0.0019 secs, 0.0001 secs, 0.0045 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [500] 1000 responses



PS C:\Users\罗宇轩\Desktop\go> $resp = iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"name":"cache_hit","delay_time":10000}' -UseBasicParsing
PS C:\Users\罗宇轩\Desktop\go> $taskId = ($resp.Content | ConvertFrom-Json).data.task_id
PS C:\Users\罗宇轩\Desktop\go> Write-Host "Task ID: $taskId"
Task ID: 336498383868985345
PS C:\Users\罗宇轩\Desktop\go> # 用刚拿到的 $taskId 替换下面的 12345
PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=$taskId"

Summary:
  Total:        0.2242 secs
  Slowest:      0.0214 secs
  Fastest:      0.0002 secs
  Average:      0.0022 secs
  Requests/sec: 22302.4217
  
  Total data:   415000 bytes
  Size/request: 83 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [4142]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [796]   |■■■■■■■■
  0.007 [11]    |
  0.009 [0]     |
  0.011 [0]     |
  0.013 [0]     |
  0.015 [0]     |
  0.017 [0]     |
  0.019 [18]    |
  0.021 [32]    |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0019 secs
  50%% in 0.0020 secs
  75%% in 0.0022 secs
  90%% in 0.0025 secs
  95%% in 0.0029 secs
  99%% in 0.0172 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0178 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0174 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0041 secs
  resp wait:    0.0019 secs, 0.0001 secs, 0.0053 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0021 secs

Status code distribution:
  [200] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=999999999"

Summary:
  Total:        1.5012 secs
  Slowest:      0.0511 secs
  Fastest:      0.0011 secs
  Average:      0.0141 secs
  Requests/sec: 3330.6621
  
  Total data:   265075 bytes
  Size/request: 53 bytes

Response time histogram:
  0.001 [1]     |
  0.006 [678]   |■■■■■■■■■■
  0.011 [2657]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.016 [117]   |■■
  0.021 [161]   |■■
  0.026 [207]   |■■■
  0.031 [704]   |■■■■■■■■■■■
  0.036 [395]   |■■■■■■
  0.041 [49]    |■
  0.046 [22]    |
  0.051 [9]     |


Latency distribution:
  10%% in 0.0041 secs
  25%% in 0.0077 secs
  50%% in 0.0087 secs
  75%% in 0.0250 secs
  90%% in 0.0310 secs
  95%% in 0.0325 secs
  99%% in 0.0380 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0163 secs
  DNS-lookup:   0.0001 secs, 0.0000 secs, 0.0150 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:    0.0139 secs, 0.0008 secs, 0.0441 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [400] 4197 responses
  [500] 803 responses



PS C:\Users\罗宇轩\Desktop\go> # 并发 100
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 100 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad100\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0060 secs
  Slowest:      0.1404 secs
  Fastest:      0.0002 secs
  Average:      0.0098 secs
  Requests/sec: 10198.7789
  
  Total data:   16637286 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.014 [305080]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.028 [844]   |
  0.042 [0]     |
  0.056 [0]     |
  0.070 [2]     |
  0.084 [19]    |
  0.098 [11]    |
  0.112 [22]    |
  0.126 [15]    |
  0.140 [31]    |


Latency distribution:
  10%% in 0.0092 secs
  25%% in 0.0094 secs
  50%% in 0.0097 secs
  75%% in 0.0102 secs
  90%% in 0.0108 secs
  95%% in 0.0112 secs
  99%% in 0.0130 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0170 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0161 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:    0.0098 secs, 0.0001 secs, 0.1227 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200] 3498 responses
  [429] 302527 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 200
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 200 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad200\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0130 secs
  Slowest:      0.0558 secs
  Fastest:      0.0001 secs
  Average:      0.0195 secs
  Requests/sec: 10234.4358
  
  Total data:   16683476 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.006 [583]   |
  0.011 [5883]  |■
  0.017 [1990]  |
  0.022 [285425]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.028 [12522] |■■
  0.034 [593]   |
  0.039 [45]    |
  0.045 [27]    |
  0.050 [36]    |
  0.056 [61]    |


Latency distribution:
  10%% in 0.0186 secs
  25%% in 0.0190 secs
  50%% in 0.0194 secs
  75%% in 0.0203 secs
  90%% in 0.0215 secs
  95%% in 0.0223 secs
  99%% in 0.0255 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0207 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0194 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0006 secs
  resp wait:    0.0195 secs, 0.0001 secs, 0.0373 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200] 3016 responses
  [429] 304150 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 500
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 500 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad500\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0316 secs
  Slowest:      0.7449 secs
  Fastest:      0.0001 secs
  Average:      0.0499 secs
  Requests/sec: 10105.9366
  
  Total data:   16260472 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.075 [298217]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.149 [774]   |
  0.224 [50]    |
  0.298 [22]    |
  0.373 [56]    |
  0.447 [55]    |
  0.521 [44]    |
  0.596 [20]    |
  0.670 [68]    |
  0.745 [25]    |


Latency distribution:
  10%% in 0.0476 secs
  25%% in 0.0484 secs
  50%% in 0.0495 secs
  75%% in 0.0518 secs
  90%% in 0.0544 secs
  95%% in 0.0562 secs
  99%% in 0.0666 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0374 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0312 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0084 secs
  resp wait:    0.0495 secs, 0.0001 secs, 0.7216 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0080 secs

Status code distribution:
  [200] 3017 responses
  [429] 296315 responses

Error distribution:
  [4165]        Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 1000
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 1000 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad1000\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0658 secs
  Slowest:      3.1771 secs
  Fastest:      0.0001 secs
  Average:      0.1039 secs
  Requests/sec: 10172.6674
  
  Total data:   14968844 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.318 [274722]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.636 [68]    |
  0.953 [61]    |
  1.271 [62]    |
  1.589 [60]    |
  1.906 [50]    |
  2.224 [99]    |
  2.542 [117]   |
  2.859 [106]   |
  3.177 [64]    |


Latency distribution:
  10%% in 0.0729 secs
  25%% in 0.0970 secs
  50%% in 0.1015 secs
  75%% in 0.1070 secs
  90%% in 0.1147 secs
  95%% in 0.1251 secs
  99%% in 0.1758 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0537 secs
  DNS-lookup:   0.0043 secs, 0.0000 secs, 0.0625 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0322 secs
  resp wait:    0.0947 secs, 0.0001 secs, 3.1394 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0323 secs

Status code distribution:
  [200] 3022 responses
  [429] 272388 responses

Error distribution:
  [30439]       Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> # 超高并发短时爆破
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 50000 -c 2000 -m POST -H "Content-Type: application/json" -d "{\"name\":\"brutal\",\"delay_time\":1000}" http://localhost:8080/Submit

Summary:
  Total:        4.5599 secs
  Slowest:      4.1453 secs
  Fastest:      0.0002 secs
  Average:      0.1699 secs
  Requests/sec: 10965.1198
  
  Total data:   1751542 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.415 [31610] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.829 [10]    |
  1.244 [49]    |
  1.658 [9]     |
  2.073 [9]     |
  2.487 [8]     |
  2.902 [10]    |
  3.316 [50]    |
  3.731 [10]    |
  4.145 [107]   |


Latency distribution:
  10%% in 0.0244 secs
  25%% in 0.1332 secs
  50%% in 0.1588 secs
  75%% in 0.1893 secs
  90%% in 0.2065 secs
  95%% in 0.2165 secs
  99%% in 0.2447 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0009 secs, 0.0000 secs, 0.1670 secs
  DNS-lookup:   0.0515 secs, 0.0000 secs, 0.1638 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0803 secs
  resp wait:    0.0420 secs, 0.0001 secs, 4.0772 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0806 secs

Status code distribution:
  [200] 950 responses
  [429] 30923 responses

Error distribution:
  [18127]       Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 长时间稳定性（10 分钟）
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 10m -c 200 -m POST -H "Content-Type: application/json" -d "{\"name\":\"longhaul\",\"delay_time\":500}" http://localhost:8080/Submit



# 1. 健康检查
hey -n 10000 -c 100 http://localhost:8080/HealthHandler

# 2. Echo JSON
hey --% -n 5000 -c 50 -m POST -H "Content-Type: application/json" -d "{\"message\":\"test\",\"panic\":false}" http://localhost:8080/EchoRequestHandler

# 3. SlowHandler 10ms
hey -n 10000 -c 100 "http://localhost:8080/SlowHandler?ms=10"

# 4. SlowHandler 200ms
hey -n 1000 -c 50 "http://localhost:8080/SlowHandler?ms=200"

# 5. Submit 限流内基准
hey --% -n 2000 -c 50 -m POST -H "Content-Type: application/json" -d "{\"name\":\"baseline\",\"delay_time\":1000}" http://localhost:8080/Submit

# 6. 404 路由
hey -n 5000 -c 50 http://localhost:8080/xxxapi

# 7. Panic 恢复
hey --% -n 1000 -c 20 -m POST -H "Content-Type: application/json" -d "{\"message\":\"crash\",\"panic\":true}" http://localhost:8080/EchoRequestHandler

# 获取任务ID
$resp = iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"name":"cache_hit","delay_time":10000}' -UseBasicParsing
$taskId = ($resp.Content | ConvertFrom-Json).data.task_id
Write-Host "Task ID: $taskId"

# 查询真实任务状态
hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=$taskId"

# 查询不存在任务ID
hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=999999999"

# 并发 100 持续30秒
hey --% -z 30s -c 100 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad100\",\"delay_time\":100}" http://localhost:8080/Submit

# 并发 200 持续30秒
hey --% -z 30s -c 200 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad200\",\"delay_time\":100}" http://localhost:8080/Submit

# 并发 500 持续30秒
hey --% -z 30s -c 500 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad500\",\"delay_time\":100}" http://localhost:8080/Submit

# 并发 1000 持续30秒
hey --% -z 30s -c 1000 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad1000\",\"delay_time\":100}" http://localhost:8080/Submit

# 超高并发短时爆破 2000并发
hey --% -n 50000 -c 2000 -m POST -H "Content-Type: application/json" -d "{\"name\":\"brutal\",\"delay_time\":1000}" http://localhost:8080/Submit

# 长时间稳定性 10分钟 200并发
hey --% -z 10m -c 200 -m POST -H "Content-Type: application/json" -d "{\"name\":\"longhaul\",\"delay_time\":500}" http://localhost:8080/Submit


>>接下来就是换了一个配置，从1个进程到5个进程，看下效果
配置
  RateLimitCapacity:   Getenvint64("RATELIMITCAPACITY", 500),
  RateLimitRefillRate: Getenvint64("RATELIMITREFILLRATE", 100), 
  WorkerPool:         Getenvint64("WORKERPOOLSIZE", 10),
  JobQueue:           Getenvint64("JOBQUEUESIZE", 100),
  ProcessConcurrency: Getenvint64("PROCESSCONCURRENCY", 5),
  mysql 默认

