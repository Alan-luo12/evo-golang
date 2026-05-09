PS C:\Users\罗宇轩\Desktop\go> hey --% -n 10000 -c 100 http://localhost:8080/HealthHandler

Summary:
  Total:        0.4419 secs
  Slowest:      0.0401 secs
  Fastest:      0.0001 secs
  Average:      0.0043 secs
  Requests/sec: 22628.1716
  
  Total data:   710000 bytes
  Size/request: 71 bytes

Response time histogram:
  0.000 [1]     |
  0.004 [5513]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [4304]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.012 [63]    |
  0.016 [19]    |
  0.020 [0]     |
  0.024 [12]    |
  0.028 [34]    |
  0.032 [28]    |
  0.036 [24]    |
  0.040 [2]     |


Latency distribution:
  10%% in 0.0037 secs
  25%% in 0.0038 secs
  50%% in 0.0041 secs
  75%% in 0.0044 secs
  90%% in 0.0050 secs
  95%% in 0.0053 secs
  99%% in 0.0215 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0238 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0229 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0025 secs
  resp wait:    0.0041 secs, 0.0001 secs, 0.0162 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0008 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 2. Echo JSON
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 5000 -c 50 -m POST -H "Content-Type: application/json" -d "{\"message\":\"test\",\"panic\":false}" http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.2334 secs
  Slowest:      0.0286 secs
  Fastest:      0.0001 secs
  Average:      0.0023 secs
  Requests/sec: 21422.8362
  
  Total data:   205000 bytes
  Size/request: 41 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4702]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [246]   |■■
  0.009 [1]     |
  0.012 [0]     |
  0.014 [0]     |
  0.017 [0]     |
  0.020 [0]     |
  0.023 [10]    |
  0.026 [24]    |
  0.029 [16]    |


Latency distribution:
  10%% in 0.0018 secs
  25%% in 0.0019 secs
  50%% in 0.0020 secs
  75%% in 0.0022 secs
  90%% in 0.0027 secs
  95%% in 0.0031 secs
  99%% in 0.0214 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0233 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0239 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0035 secs
  resp wait:    0.0020 secs, 0.0001 secs, 0.0058 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0039 secs

Status code distribution:
  [200] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 3. SlowHandler 10ms
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 10000 -c 100 "http://localhost:8080/SlowHandler?ms=10"

Summary:
  Total:        1.0807 secs
  Slowest:      0.0387 secs
  Fastest:      0.0096 secs
  Average:      0.0107 secs
  Requests/sec: 9253.1667
  
  Total data:   520000 bytes
  Size/request: 52 bytes

Response time histogram:
  0.010 [1]     |
  0.013 [9899]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.015 [0]     |
  0.018 [0]     |
  0.021 [0]     |
  0.024 [0]     |
  0.027 [0]     |
  0.030 [0]     |
  0.033 [20]    |
  0.036 [45]    |
  0.039 [35]    |


Latency distribution:
  10%% in 0.0101 secs
  25%% in 0.0103 secs
  50%% in 0.0104 secs
  75%% in 0.0106 secs
  90%% in 0.0108 secs
  95%% in 0.0110 secs
  99%% in 0.0303 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0259 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0234 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0030 secs
  resp wait:    0.0104 secs, 0.0096 secs, 0.0130 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 4. SlowHandler 200ms
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 1000 -c 50 "http://localhost:8080/SlowHandler?ms=200"

Summary:
  Total:        4.0330 secs
  Slowest:      0.2233 secs
  Fastest:      0.1998 secs
  Average:      0.2016 secs
  Requests/sec: 247.9558
  
  Total data:   53000 bytes
  Size/request: 53 bytes

Response time histogram:
  0.200 [1]     |
  0.202 [947]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.204 [2]     |
  0.207 [0]     |
  0.209 [0]     |
  0.212 [0]     |
  0.214 [0]     |
  0.216 [0]     |
  0.219 [0]     |
  0.221 [19]    |■
  0.223 [31]    |■


Latency distribution:
  10%% in 0.2001 secs
  25%% in 0.2003 secs
  50%% in 0.2005 secs
  75%% in 0.2008 secs
  90%% in 0.2011 secs
  95%% in 0.2189 secs
  99%% in 0.2228 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0010 secs, 0.0000 secs, 0.0218 secs
  DNS-lookup:   0.0009 secs, 0.0000 secs, 0.0208 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:    0.2005 secs, 0.1998 secs, 0.2024 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200] 1000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 5. Submit（限流内）
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 2000 -c 50 -m POST -H "Content-Type: application/json" -d "{\"name\":\"bench\",\"delay_time\":1000}" http://localhost:8080/Submit

Summary:
  Total:        0.2271 secs
  Slowest:      0.0681 secs
  Fastest:      0.0002 secs
  Average:      0.0052 secs
  Requests/sec: 8806.5496
  
  Total data:   124480 bytes
  Size/request: 62 bytes

Response time histogram:
  0.000 [1]     |
  0.007 [1936]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.014 [2]     |
  0.021 [1]     |
  0.027 [20]    |
  0.034 [2]     |
  0.041 [0]     |
  0.048 [10]    |
  0.055 [16]    |
  0.061 [5]     |
  0.068 [7]     |


Latency distribution:
  10%% in 0.0018 secs
  25%% in 0.0032 secs
  50%% in 0.0047 secs
  75%% in 0.0051 secs
  90%% in 0.0056 secs
  95%% in 0.0061 secs
  99%% in 0.0509 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0003 secs, 0.0000 secs, 0.0238 secs
  DNS-lookup:   0.0005 secs, 0.0000 secs, 0.0219 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0038 secs
  resp wait:    0.0046 secs, 0.0001 secs, 0.0461 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200] 515 responses
  [429] 1485 responses




PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 7. Getstatus 缓存击穿（不存在的 ID）
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 5000 -c 50 "http://localhost:8080/Getstatus?id=999999999"

Summary:
  Total:        1.4607 secs
  Slowest:      0.0506 secs
  Fastest:      0.0006 secs
  Average:      0.0138 secs
  Requests/sec: 3422.9261
  
  Total data:   258400 bytes
  Size/request: 51 bytes

Response time histogram:
  0.001 [1]     |
  0.006 [480]   |■■■■■■■
  0.011 [2944]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.016 [124]   |■■
  0.021 [120]   |■■
  0.026 [164]   |■■
  0.031 [658]   |■■■■■■■■■
  0.036 [449]   |■■■■■■
  0.041 [19]    |
  0.046 [2]     |
  0.051 [39]    |■


Latency distribution:
  10%% in 0.0059 secs
  25%% in 0.0076 secs
  50%% in 0.0083 secs
  75%% in 0.0239 secs
  90%% in 0.0306 secs
  95%% in 0.0324 secs
  99%% in 0.0365 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0218 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0204 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:    0.0135 secs, 0.0006 secs, 0.0385 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [400] 4464 responses
  [500] 536 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 8. 404 路由
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 5000 -c 50 http://localhost:8080/xxxapi

Summary:
  Total:        0.2284 secs
  Slowest:      0.0269 secs
  Fastest:      0.0001 secs
  Average:      0.0022 secs
  Requests/sec: 21888.8406
  
  Total data:   95000 bytes
  Size/request: 19 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4843]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [105]   |■
  0.008 [1]     |
  0.011 [0]     |
  0.014 [0]     |
  0.016 [0]     |
  0.019 [0]     |
  0.022 [3]     |
  0.024 [35]    |
  0.027 [12]    |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0019 secs
  50%% in 0.0020 secs
  75%% in 0.0022 secs
  90%% in 0.0025 secs
  95%% in 0.0026 secs
  99%% in 0.0211 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0215 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0214 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0030 secs
  resp wait:    0.0019 secs, 0.0001 secs, 0.0061 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0028 secs

Status code distribution:
  [404] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 9. Panic 恢复
PS C:\Users\罗宇轩\Desktop\go> hey --% -n 1000 -c 20 -m POST -H "Content-Type: application/json" -d "{\"message\":\"crash\",\"panic\":true}" http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.1437 secs
  Slowest:      0.0272 secs
  Fastest:      0.0002 secs
  Average:      0.0028 secs
  Requests/sec: 6957.9932
  
  Total data:   57000 bytes
  Size/request: 57 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [534]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [445]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [0]     |
  0.011 [0]     |
  0.014 [0]     |
  0.016 [0]     |
  0.019 [0]     |
  0.022 [0]     |
  0.025 [8]     |■
  0.027 [12]    |■


Latency distribution:
  10%% in 0.0009 secs
  25%% in 0.0012 secs
  50%% in 0.0025 secs
  75%% in 0.0033 secs
  90%% in 0.0036 secs
  95%% in 0.0041 secs
  99%% in 0.0253 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0004 secs, 0.0000 secs, 0.0230 secs
  DNS-lookup:   0.0004 secs, 0.0000 secs, 0.0217 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:    0.0023 secs, 0.0001 secs, 0.0050 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [500] 1000 responses



PS C:\Users\罗宇轩\Desktop\go> # 并发 100
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 100 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad100\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0062 secs
  Slowest:      0.0613 secs
  Fastest:      0.0002 secs
  Average:      0.0104 secs
  Requests/sec: 9648.1765
  
  Total data:   15742230 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.006 [6459]  |■
  0.012 [271320]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.018 [11432] |■■
  0.025 [188]   |
  0.031 [64]    |
  0.037 [21]    |
  0.043 [2]     |
  0.049 [0]     |
  0.055 [5]     |
  0.061 [13]    |


Latency distribution:
  10%% in 0.0095 secs
  25%% in 0.0099 secs
  50%% in 0.0103 secs
  75%% in 0.0109 secs
  90%% in 0.0116 secs
  95%% in 0.0122 secs
  99%% in 0.0143 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0220 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0212 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0008 secs
  resp wait:    0.0103 secs, 0.0001 secs, 0.0407 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200] 3405 responses
  [429] 286100 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 200
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 200 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad200\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0132 secs
  Slowest:      0.0605 secs
  Fastest:      0.0002 secs
  Average:      0.0200 secs
  Requests/sec: 9988.2825
  
  Total data:   16282008 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.006 [548]   |
  0.012 [5573]  |■
  0.018 [2100]  |
  0.024 [285132]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.030 [6094]  |■
  0.036 [185]   |
  0.042 [28]    |
  0.048 [34]    |
  0.054 [47]    |
  0.061 [38]    |


Latency distribution:
  10%% in 0.0190 secs
  25%% in 0.0194 secs
  50%% in 0.0199 secs
  75%% in 0.0208 secs
  90%% in 0.0220 secs
  95%% in 0.0228 secs
  99%% in 0.0261 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0249 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0240 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0007 secs
  resp wait:    0.0200 secs, 0.0001 secs, 0.0367 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200] 2934 responses
  [429] 296846 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 500
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 500 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad500\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0346 secs
  Slowest:      0.7451 secs
  Fastest:      0.0002 secs
  Average:      0.0507 secs
  Requests/sec: 9945.7506
  
  Total data:   15989686 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.075 [293812]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.149 [180]   |
  0.224 [76]    |
  0.298 [45]    |
  0.373 [55]    |
  0.447 [15]    |
  0.522 [14]    |
  0.596 [74]    |
  0.671 [52]    |
  0.745 [45]    |


Latency distribution:
  10%% in 0.0487 secs
  25%% in 0.0495 secs
  50%% in 0.0506 secs
  75%% in 0.0524 secs
  90%% in 0.0543 secs
  95%% in 0.0558 secs
  99%% in 0.0676 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0363 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0313 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0068 secs
  resp wait:    0.0503 secs, 0.0001 secs, 0.7217 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0072 secs

Status code distribution:
  [200] 2930 responses
  [429] 291439 responses

Error distribution:
  [4348]        Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> # 并发 1000
PS C:\Users\罗宇轩\Desktop\go> hey --% -z 30s -c 1000 -m POST -H "Content-Type: application/json" -d "{\"name\":\"grad1000\",\"delay_time\":100}" http://localhost:8080/Submit

Summary:
  Total:        30.0684 secs
  Slowest:      2.7808 secs
  Fastest:      0.0001 secs
  Average:      0.1011 secs
  Requests/sec: 10577.1206
  
  Total data:   15321558 bytes
  Size/request: 54 bytes

Response time histogram:
  0.000 [1]     |
  0.278 [280980]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.556 [59]    |
  0.834 [72]    |
  1.112 [62]    |
  1.390 [79]    |
  1.669 [98]    |
  1.947 [81]    |
  2.225 [101]   |
  2.503 [119]   |
  2.781 [77]    |


Latency distribution:
  10%% in 0.0723 secs
  25%% in 0.0988 secs
  50%% in 0.1012 secs
  75%% in 0.1047 secs
  90%% in 0.1077 secs
  95%% in 0.1096 secs
  99%% in 0.1364 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0592 secs
  DNS-lookup:   0.0035 secs, 0.0000 secs, 0.0670 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0205 secs
  resp wait:    0.0938 secs, 0.0001 secs, 2.7387 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0166 secs

Status code distribution:
  [200] 3381 responses
  [429] 278348 responses

Error distribution:
  [36308]       Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> hey --% -n 50000 -c 2000 -m POST -H "Content-Type: application/json" -d "{\"name\":\"brutal\",\"delay_time\":1000}" http://localhost:8080/Submit

Summary:
  Total:        4.9355 secs
  Slowest:      4.3714 secs
  Fastest:      0.0003 secs
  Average:      0.1802 secs
  Requests/sec: 10130.6462
  
  Total data:   1562288 bytes
  Size/request: 55 bytes

Response time histogram:
  0.000 [1]     |
  0.437 [28113] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.875 [9]     |
  1.312 [54]    |
  1.749 [9]     |
  2.186 [8]     |
  2.623 [8]     |
  3.060 [10]    |
  3.497 [39]    |
  3.934 [25]    |
  4.371 [100]   |


Latency distribution:
  10%% in 0.0315 secs
  25%% in 0.1345 secs
  50%% in 0.1632 secs
  75%% in 0.1905 secs
  90%% in 0.2274 secs
  95%% in 0.2473 secs
  99%% in 0.3054 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0008 secs, 0.0000 secs, 0.1009 secs
  DNS-lookup:   0.0538 secs, 0.0000 secs, 0.1679 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0338 secs
  resp wait:    0.0500 secs, 0.0002 secs, 4.2954 secs
  resp read:    0.0002 secs, 0.0000 secs, 0.0833 secs

Status code distribution:
  [200] 937 responses
  [429] 27439 responses

Error distribution:
  [21624]       Post "http://localhost:8080/Submit": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

PS C:\Users\罗宇轩\Desktop\go> hey --% -z 10m -c 200 -m POST -H "Content-Type: application/json" -d "{\"name\":\"longhaul\",\"delay_time\":500}" http://localhost:8080/Submit

Summary:
  Total:        388.9770 secs
  Slowest:      0.0725 secs
  Fastest:      0.0002 secs
  Average:      0.0778 secs
  Requests/sec: 7921.7981
  
  Total data:   167619790 bytes
  Size/request: 167 bytes

Response time histogram:
  0.000 [1]     |
  0.007 [3300]  |
  0.015 [20466] |■
  0.022 [84132] |■■■■
  0.029 [869715]        |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.036 [19767] |■
  0.044 [2242]  |                                                                                                            
  0.051 [307]   |                                                                                                            
  0.058 [24]    |                                                                                                            
  0.065 [28]    |                                                                                                            
  0.072 [18]    |


Latency distribution:
  10%% in 0.0215 secs
  25%% in 0.0243 secs
  50%% in 0.0254 secs
  75%% in 0.0263 secs
  90%% in 0.0272 secs
  95%% in 0.0278 secs
  99%% in 0.0335 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0000 secs, 0.0000 secs, 0.0252 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0230 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0012 secs
  resp wait:    0.0776 secs, 0.0001 secs, 0.0590 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [200] 12550 responses
  [429] 987450 responses



PS C:\Users\罗宇轩\Desktop\go> iwr -Uri http://localhost:8080/Submit -Method Post -ContentType "application/json" -Body '{"name":"cache_test","delay_time":500}' -UseBasicParsing


StatusCode        : 200
StatusDescription : OK
Content           : {"code":0,"msg":"Success","data":{"task_id":336346476630245377,"status":"submitted"}}
                    
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 86
                    Content-Type: application/json
                    Date: Sat, 09 May 2026 08:22:02 GMT
                    
                    {"code":0,"msg":"Success","data":{"task_id":336346476630245377,"status":"submitted"}}
                    
Forms             : 
Headers           : {[Content-Length, 86], [Content-Type, application/json], [Date, Sat, 09 May 2026 08:22:02 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : 
RawContentLength  : 86



PS C:\Users\罗宇轩\Desktop\go> hey --% -n 5000 -c 50 "http://localhost:8080/Getstatus?id=336346476630245377"

Summary:
  Total:        0.2227 secs
  Slowest:      0.0269 secs
  Fastest:      0.0002 secs
  Average:      0.0022 secs
  Requests/sec: 22451.7086
  
  Total data:   415000 bytes
  Size/request: 83 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4822]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [126]   |■
  0.008 [1]     |
  0.011 [0]     |
  0.014 [0]     |
  0.016 [0]     |
  0.019 [0]     |
  0.022 [1]     |
  0.024 [24]    |
  0.027 [25]    |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0019 secs
  50%% in 0.0020 secs
  75%% in 0.0021 secs
  90%% in 0.0024 secs
  95%% in 0.0026 secs
  99%% in 0.0215 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0222 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0225 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:    0.0019 secs, 0.0002 secs, 0.0057 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0022 secs

Status code distribution:
  [200] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> hey --% -n 5000 -c 50 "http://localhost:8080/Getstatus?id=999999999"

Summary:
  Total:        1.4506 secs
  Slowest:      0.0537 secs
  Fastest:      0.0008 secs
  Average:      0.0136 secs
  Requests/sec: 3446.7495
  
  Total data:   259300 bytes
  Size/request: 51 bytes

Response time histogram:
  0.001 [1]     |
  0.006 [554]   |■■■■■■■■
  0.011 [2911]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.017 [133]   |■■
  0.022 [139]   |■■
  0.027 [207]   |■■■
  0.033 [837]   |■■■■■■■■■■■■
  0.038 [167]   |■■
  0.043 [12]    |
  0.048 [21]    |
  0.054 [18]    |


Latency distribution:
  10%% in 0.0052 secs
  25%% in 0.0075 secs
  50%% in 0.0083 secs
  75%% in 0.0229 secs
  90%% in 0.0304 secs
  95%% in 0.0321 secs
  99%% in 0.0379 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0222 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0204 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0019 secs
  resp wait:    0.0134 secs, 0.0008 secs, 0.0393 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0009 secs

Status code distribution:
  [400] 4428 responses
  [500] 572 responses



PS C:\Users\罗宇轩\Desktop\go> 














