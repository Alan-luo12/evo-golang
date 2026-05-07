PS C:\Users\罗宇轩\Desktop\go> hey -n 10000 -c 100 http://localhost:8080/HealthHandler

Summary:
  Total:        0.4035 secs
  Slowest:      0.0252 secs
  Fastest:      0.0001 secs
  Average:      0.0040 secs
  Requests/sec: 24782.8588
  
  Total data:   710000 bytes
  Size/request: 71 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [483]   |■■
  0.005 [9174]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [169]   |■
  0.010 [73]    |
  0.013 [0]     |
  0.015 [0]     |
  0.018 [7]     |
  0.020 [26]    |
  0.023 [43]    |
  0.025 [24]    |


Latency distribution:
  10%% in 0.0035 secs
  25%% in 0.0037 secs
  50%% in 0.0038 secs
  75%% in 0.0040 secs
  90%% in 0.0044 secs
  95%% in 0.0047 secs
  99%% in 0.0170 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0187 secs
  DNS-lookup:   0.0001 secs, 0.0000 secs, 0.0147 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0053 secs
  resp wait:    0.0037 secs, 0.0001 secs, 0.0091 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0035 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 -m POST -H "Content-Type: application/json" -d '{\"message\":\"test\",\"panic\":false}' http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.2138 secs
  Slowest:      0.0243 secs
  Fastest:      0.0001 secs
  Average:      0.0021 secs
  Requests/sec: 23383.5534
  
  Total data:   205000 bytes
  Size/request: 41 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4729]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [215]   |■■
  0.007 [5]     |
  0.010 [0]     |
  0.012 [0]     |
  0.015 [0]     |
  0.017 [1]     |
  0.019 [15]    |
  0.022 [21]    |
  0.024 [13]    |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0018 secs
  50%% in 0.0019 secs
  75%% in 0.0021 secs
  90%% in 0.0023 secs
  95%% in 0.0026 secs
  99%% in 0.0169 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0190 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0181 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0036 secs
  resp wait:    0.0018 secs, 0.0001 secs, 0.0047 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0038 secs

Status code distribution:
  [200] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> 
PS C:\Users\罗宇轩\Desktop\go> hey -n 10000 -c 100 "http://localhost:8080/SlowHandler?ms=10"

Summary:
  Total:        1.0751 secs
  Slowest:      0.0335 secs
  Fastest:      0.0097 secs
  Average:      0.0107 secs
  Requests/sec: 9301.0459
  
  Total data:   520000 bytes
  Size/request: 52 bytes

Response time histogram:
  0.010 [1]     |
  0.012 [9896]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.014 [3]     |
  0.017 [0]     |
  0.019 [0]     |
  0.022 [0]     |
  0.024 [0]     |
  0.026 [0]     |
  0.029 [17]    |
  0.031 [32]    |
  0.033 [51]    |


Latency distribution:
  10%% in 0.0101 secs
  25%% in 0.0103 secs
  50%% in 0.0104 secs
  75%% in 0.0106 secs
  90%% in 0.0108 secs
  95%% in 0.0110 secs
  99%% in 0.0264 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0206 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0183 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0026 secs
  resp wait:    0.0104 secs, 0.0096 secs, 0.0135 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> hey -n 1000 -c 50 "http://localhost:8080/SlowHandler?ms=200"

Summary:
  Total:        4.0301 secs
  Slowest:      0.2182 secs
  Fastest:      0.1996 secs
  Average:      0.2014 secs
  Requests/sec: 248.1346
  
  Total data:   53000 bytes
  Size/request: 53 bytes

Response time histogram:
  0.200 [1]     |
  0.202 [900]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.203 [49]    |■■
  0.205 [0]     |
  0.207 [0]     |
  0.209 [0]     |
  0.211 [0]     |
  0.213 [0]     |
  0.215 [1]     |
  0.216 [17]    |■
  0.218 [32]    |■


Latency distribution:
  10%% in 0.2001 secs
  25%% in 0.2003 secs
  50%% in 0.2005 secs
  75%% in 0.2010 secs
  90%% in 0.2015 secs
  95%% in 0.2145 secs
  99%% in 0.2177 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0008 secs, 0.0000 secs, 0.0172 secs
  DNS-lookup:   0.0007 secs, 0.0000 secs, 0.0158 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:    0.2005 secs, 0.1996 secs, 0.2026 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200] 1000 responses



PS C:\Users\罗宇轩\Desktop\go> hey -n 2000 -c 50 -m POST -H "Content-Type: application/json" -d '{\"name\":\"bench_task\",\"delay_time\":1000}' http://localhost:8080/Submit

Summary:
  Total:        0.1373 secs
  Slowest:      0.0227 secs
  Fastest:      0.0003 secs
  Average:      0.0034 secs
  Requests/sec: 14566.9713
  
  Total data:   172000 bytes
  Size/request: 86 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [515]   |■■■■■■■■■■■■■■■
  0.005 [1358]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.007 [74]    |■■
  0.009 [2]     |
  0.012 [0]     |
  0.014 [0]     |
  0.016 [5]     |
  0.018 [16]    |
  0.020 [26]    |■
  0.023 [3]     |


Latency distribution:
  10%% in 0.0023 secs
  25%% in 0.0026 secs
  50%% in 0.0029 secs
  75%% in 0.0033 secs
  90%% in 0.0036 secs
  95%% in 0.0052 secs
  99%% in 0.0191 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0164 secs
  DNS-lookup:   0.0003 secs, 0.0000 secs, 0.0151 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0042 secs
  resp wait:    0.0029 secs, 0.0003 secs, 0.0065 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0026 secs

Status code distribution:
  [200] 2000 responses





PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 http://localhost:8080/xxxapi

Summary:
  Total:        0.2112 secs
  Slowest:      0.0223 secs
  Fastest:      0.0001 secs
  Average:      0.0021 secs
  Requests/sec: 23674.7020
  
  Total data:   95000 bytes
  Size/request: 19 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [4504]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [433]   |■■■■
  0.007 [12]    |
  0.009 [0]     |
  0.011 [0]     |
  0.013 [0]     |
  0.016 [0]     |
  0.018 [19]    |
  0.020 [13]    |
  0.022 [18]    |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0018 secs
  50%% in 0.0019 secs
  75%% in 0.0021 secs
  90%% in 0.0023 secs
  95%% in 0.0026 secs
  99%% in 0.0163 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0176 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0178 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0033 secs
  resp wait:    0.0018 secs, 0.0001 secs, 0.0060 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0037 secs

Status code distribution:
  [404] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> hey -n 1000 -c 20 -m POST -H "Content-Type: application/json" -d '{\"message\":\"crash\",\"panic\":true}' http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.1256 secs
  Slowest:      0.0196 secs
  Fastest:      0.0002 secs
  Average:      0.0024 secs
  Requests/sec: 7964.7702
  
  Total data:   57000 bytes
  Size/request: 57 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [507]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [459]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [13]    |■
  0.008 [0]     |
  0.010 [0]     |
  0.012 [0]     |
  0.014 [0]     |
  0.016 [0]     |
  0.018 [8]     |■
  0.020 [12]    |■


Latency distribution:
  10%% in 0.0004 secs
  25%% in 0.0011 secs
  50%% in 0.0021 secs
  75%% in 0.0031 secs
  90%% in 0.0034 secs
  95%% in 0.0037 secs
  99%% in 0.0178 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0003 secs, 0.0000 secs, 0.0158 secs
  DNS-lookup:   0.0003 secs, 0.0000 secs, 0.0149 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:    0.0020 secs, 0.0001 secs, 0.0049 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0007 secs

Status code distribution:
  [500] 1000 responses



PS C:\Users\罗宇轩\Desktop\go> 



>>注意这一条是在缓存时间之内的查询表示一切正常
PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=336034243949887489"

Summary:
  Total:        0.2459 secs
  Slowest:      0.0238 secs
  Fastest:      0.0002 secs
  Average:      0.0024 secs
  Requests/sec: 20336.3717
  
  Total data:   420000 bytes
  Size/request: 84 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [4385]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [560]   |■■■■■
  0.007 [4]     |
  0.010 [0]     |
  0.012 [0]     |
  0.014 [0]     |
  0.017 [0]     |
  0.019 [5]     |
  0.021 [26]    |
  0.024 [19]    |


Latency distribution:
  10%% in 0.0020 secs
  25%% in 0.0021 secs
  50%% in 0.0022 secs
  75%% in 0.0024 secs
  90%% in 0.0027 secs
  95%% in 0.0031 secs
  99%% in 0.0177 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0195 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0180 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:    0.0021 secs, 0.0002 secs, 0.0055 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0033 secs

Status code distribution:
  [200] 5000 responses


>>注意这是缓存击穿的数据
 PS C:\Users\罗宇轩\Desktop\go> hey -n 5000 -c 50 "http://localhost:8080/Getstatus?id=335957101248380900"

Summary:
  Total:        1.4461 secs
  Slowest:      0.0524 secs
  Fastest:      0.0004 secs
  Average:      0.0138 secs
  Requests/sec: 3457.4644
  
  Total data:   246800 bytes
  Size/request: 49 bytes

Response time histogram:
  0.000 [1]     |
  0.006 [332]   |■■■■
  0.011 [3141]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.016 [40]    |■
  0.021 [78]    |■
  0.026 [125]   |■■
  0.032 [1084]  |■■■■■■■■■■■■■■
  0.037 [154]   |■■
  0.042 [3]     |
  0.047 [13]    |
  0.052 [29]    |


Latency distribution:
  10%% in 0.0067 secs
  25%% in 0.0073 secs
  50%% in 0.0079 secs
  75%% in 0.0270 secs
  90%% in 0.0296 secs
  95%% in 0.0309 secs
  99%% in 0.0361 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0193 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0171 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:    0.0136 secs, 0.0004 secs, 0.0365 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [400] 4928 responses
  [500] 72 responses
