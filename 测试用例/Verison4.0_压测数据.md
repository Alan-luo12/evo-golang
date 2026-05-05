PS C:\Users\罗宇轩\Desktop\go> hey -c 100 -n 10000 http://localhost:8080/HealthHandler

Summary:
  Total:        0.4311 secs
  Slowest:      0.0296 secs
  Fastest:      0.0001 secs
  Average:      0.0042 secs
  Requests/sec: 23198.8796
  
  Total data:   710000 bytes
  Size/request: 71 bytes

Response time histogram:
  0.000 [1]     |
  0.003 [377]   |■■
  0.006 [9363]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.009 [132]   |■
  0.012 [27]    |
  0.015 [0]     |
  0.018 [2]     |
  0.021 [22]    |
  0.024 [33]    |
  0.027 [31]    |
  0.030 [12]    |


Latency distribution:
  10%% in 0.0037 secs
  25%% in 0.0038 secs
  50%% in 0.0040 secs
  75%% in 0.0043 secs
  90%% in 0.0047 secs
  95%% in 0.0050 secs
  99%% in 0.0175 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0207 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0203 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0058 secs
  resp wait:    0.0040 secs, 0.0001 secs, 0.0094 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0048 secs

Status code distribution:
  [200] 10000 responses




PS C:\Users\罗宇轩\Desktop\go> hey -c 50 -n 1000 "http://localhost:8080/SlowHandler?ms=200"

Summary:
  Total:        4.0284 secs
  Slowest:      0.2184 secs
  Fastest:      0.1997 secs
  Average:      0.2013 secs
  Requests/sec: 248.2381
  
  Total data:   53000 bytes
  Size/request: 53 bytes

Response time histogram:
  0.200 [1]     |
  0.202 [943]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.203 [6]     |
  0.205 [0]     |
  0.207 [0]     |
  0.209 [0]     |
  0.211 [0]     |
  0.213 [0]     |
  0.215 [0]     |
  0.217 [14]    |■
  0.218 [36]    |■■


Latency distribution:
  10%% in 0.2001 secs
  25%% in 0.2002 secs
  50%% in 0.2005 secs
  75%% in 0.2007 secs
  90%% in 0.2010 secs
  95%% in 0.2148 secs
  99%% in 0.2180 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0008 secs, 0.0000 secs, 0.0171 secs
  DNS-lookup:   0.0007 secs, 0.0000 secs, 0.0150 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0018 secs
  resp wait:    0.2005 secs, 0.1996 secs, 0.2020 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0004 secs

Status code distribution:
  [200] 1000 responses




PS C:\Users\罗宇轩\Desktop\go> hey -c 50 -n 5000 "http://localhost:8080/Getstatus?id=1"

Summary:
  Total:        1.2330 secs
  Slowest:      0.0501 secs
  Fastest:      0.0003 secs
  Average:      0.0117 secs
  Requests/sec: 4054.9997
  
  Total data:   320000 bytes
  Size/request: 64 bytes

Response time histogram:
  0.000 [1]     |
  0.005 [290]   |■■■
  0.010 [3750]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.015 [41]    |
  0.020 [19]    |
  0.025 [46]    |
  0.030 [338]   |■■■■
  0.035 [441]   |■■■■■
  0.040 [33]    |
  0.045 [15]    |
  0.050 [26]    |


Latency distribution:
  10%% in 0.0065 secs
  25%% in 0.0073 secs
  50%% in 0.0078 secs
  75%% in 0.0086 secs
  90%% in 0.0303 secs
  95%% in 0.0319 secs
  99%% in 0.0379 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0182 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0164 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0010 secs
  resp wait:    0.0115 secs, 0.0003 secs, 0.0395 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0010 secs

Status code distribution:
  [200] 5000 responses


PS C:\Users\罗宇轩\Desktop\go> hey -c 50 -n 5000 -m POST -T "application/json" -d '{\"name\":\"stress_task\",\"delay_time\":10}' http://localhost:8080/Submit

Summary:
  Total:        1.5404 secs
  Slowest:      0.0548 secs
  Fastest:      0.0006 secs
  Average:      0.0148 secs
  Requests/sec: 3245.8246
  
  Total data:   280014 bytes
  Size/request: 56 bytes

Response time histogram:
  0.001 [1]     |
  0.006 [1025]  |■■■■■■■■■■■■■■■■■■■■■■■■
  0.011 [1678]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.017 [268]   |■■■■■■
  0.022 [320]   |■■■■■■■■
  0.028 [1086]  |■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.033 [526]   |■■■■■■■■■■■■■
  0.039 [42]    |■
  0.044 [2]     |
  0.049 [22]    |■
  0.055 [30]    |■


Latency distribution:
  10%% in 0.0027 secs
  25%% in 0.0068 secs
  50%% in 0.0096 secs
  75%% in 0.0253 secs
  90%% in 0.0283 secs
  95%% in 0.0300 secs
  99%% in 0.0456 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0210 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0189 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0030 secs
  resp wait:    0.0146 secs, 0.0005 secs, 0.0533 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0011 secs

Status code distribution:
  [200] 1 responses
  [500] 4999 responses





PS C:\Users\罗宇轩\Desktop\go> hey -c 50 -n 5000 http://localhost:8080/api_not_exist

Summary:
  Total:        0.2149 secs
  Slowest:      0.0228 secs
  Fastest:      0.0001 secs
  Average:      0.0021 secs
  Requests/sec: 23266.2675
  
  Total data:   95000 bytes
  Size/request: 19 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [4521]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [420]   |■■■■
  0.007 [8]     |
  0.009 [0]     |
  0.011 [0]     |
  0.014 [0]     |
  0.016 [4]     |
  0.018 [19]    |
  0.021 [20]    |
  0.023 [7]     |


Latency distribution:
  10%% in 0.0017 secs
  25%% in 0.0018 secs
  50%% in 0.0020 secs
  75%% in 0.0021 secs
  90%% in 0.0024 secs
  95%% in 0.0026 secs
  99%% in 0.0156 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0000 secs, 0.0168 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0167 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0037 secs
  resp wait:    0.0018 secs, 0.0001 secs, 0.0050 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0045 secs

Status code distribution:
  [404] 5000 responses



PS C:\Users\罗宇轩\Desktop\go> hey -c 100 -n 10000 "http://localhost:8080/SlowHandler?ms=10"

Summary:
  Total:        1.0739 secs
  Slowest:      0.0326 secs
  Fastest:      0.0097 secs
  Average:      0.0107 secs
  Requests/sec: 9311.7794
  
  Total data:   520000 bytes
  Size/request: 52 bytes

Response time histogram:
  0.010 [1]     |
  0.012 [9896]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.014 [3]     |
  0.017 [0]     |
  0.019 [0]     |
  0.021 [0]     |
  0.023 [0]     |
  0.026 [2]     |
  0.028 [18]    |
  0.030 [44]    |
  0.033 [36]    |


Latency distribution:
  10%% in 0.0101 secs
  25%% in 0.0103 secs
  50%% in 0.0105 secs
  75%% in 0.0106 secs
  90%% in 0.0109 secs
  95%% in 0.0111 secs
  99%% in 0.0255 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0002 secs, 0.0000 secs, 0.0194 secs
  DNS-lookup:   0.0002 secs, 0.0000 secs, 0.0185 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0020 secs
  resp wait:    0.0104 secs, 0.0096 secs, 0.0142 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [200] 10000 responses



PS C:\Users\罗宇轩\Desktop\go> hey -c 50 -n 1000 "http://localhost:8080/SlowHandler?ms=500"

Summary:
  Total:        10.0305 secs
  Slowest:      0.5184 secs
  Fastest:      0.4998 secs
  Average:      0.5014 secs
  Requests/sec: 99.6959
  
  Total data:   53000 bytes
  Size/request: 53 bytes

Response time histogram:
  0.500 [1]     |
  0.502 [931]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.503 [18]    |■
  0.505 [0]     |
  0.507 [0]     |
  0.509 [0]     |
  0.511 [0]     |
  0.513 [0]     |
  0.515 [0]     |
  0.517 [14]    |■
  0.518 [36]    |■■


Latency distribution:
  10%% in 0.5001 secs
  25%% in 0.5003 secs
  50%% in 0.5006 secs
  75%% in 0.5009 secs
  90%% in 0.5013 secs
  95%% in 0.5147 secs
  99%% in 0.5177 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0008 secs, 0.0000 secs, 0.0175 secs
  DNS-lookup:   0.0007 secs, 0.0000 secs, 0.0156 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0009 secs
  resp wait:    0.5006 secs, 0.4997 secs, 0.5029 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0005 secs

Status code distribution:
  [200] 1000 responses



Summary:
  Total:        0.0923 secs
  Slowest:      0.0188 secs
  Fastest:      0.0001 secs
  Average:      0.0018 secs
  Requests/sec: 10836.8662
  
  Total data:   57000 bytes
  Size/request: 57 bytes

Response time histogram:
  0.000 [1]     |
  0.002 [675]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.004 [303]   |■■■■■■■■■■■■■■■■■■
  0.006 [1]     |
  0.008 [0]     |
  0.009 [0]     |
  0.011 [0]     |
  0.013 [0]     |
  0.015 [0]     |
  0.017 [8]     |
  0.019 [12]    |■


Latency distribution:
  10%% in 0.0003 secs
  25%% in 0.0010 secs
  50%% in 0.0017 secs
  75%% in 0.0021 secs
  90%% in 0.0023 secs
  95%% in 0.0028 secs
  99%% in 0.0171 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0003 secs, 0.0000 secs, 0.0158 secs
  DNS-lookup:   0.0003 secs, 0.0000 secs, 0.0146 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0011 secs
  resp wait:    0.0014 secs, 0.0001 secs, 0.0038 secs
  resp read:    0.0000 secs, 0.0000 secs, 0.0006 secs

Status code distribution:
  [500] 1000 responses

