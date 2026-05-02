### Version1.0
暂无数据，仅为初始化尝试，熟悉语法

### Verison 1.5

简要说明：本次v1.5是对v1的补充主要增加了简单中间件和压测，没有任何复杂业务和数据库之类，是一个 最小 HTTP 系统

结果
///基线压测
  Total:        30.0050 secs
  Slowest:      0.0417 secs
  Fastest:      0.0001 secs
  Average:      0.0044 secs
  Requests/sec: 22745.9757
  
///延迟接口压测
Error distribution:
  [213859]      Get "http://localhost:8080/SlowHandler?ms=2000": dial tcp [::1]:8080: connectex: No connection could be made because the target machine actively refused it.

    Total:        31.9984 secs
  Slowest:      5.6993 secs
  Fastest:      1.9997 secs
  Average:      2.1526 secs
  Requests/sec: 7354.7973
  
  Total data:   708939 bytes
  Size/request: 33 bytes
///问题记录
1. panic 是否导致服务退出？（是 / 否）
没有

2. 高并发时延迟分布是否恶化？
严重恶化

3. 错误 JSON 是否导致 panic？
没有paic 是 INvalid json

4. go语言天然并发的优势？
天然支持并发，自动创建协程理想QPS极高

### Version 2.0

---hey -n 10000 -c 100 http://localhost:8080/HealthHandler

Summary:
  Total:        0.6865 secs
  Slowest:      0.0376 secs
  Fastest:      0.0001 secs
  Average:      0.0068 secs
  Requests/sec: 14567.6397

  Total data:   710000 bytes
  Size/request: 71 bytes

Latency distribution:
  10%% in 0.0059 secs
  25%% in 0.0062 secs
  50%% in 0.0065 secs
  75%% in 0.0069 secs
  90%% in 0.0076 secs
  95%% in 0.0082 secs
  99%% in 0.0195 secs


Status code distribution:
  [200] 10000 responses

---hey -n 1000 -c 50 http://localhost:8080/SlowHandler?ms=200

Summary:
  Total:        4.0344 secs
  Slowest:      0.2216 secs
  Fastest:      0.1997 secs
  Average:      0.2016 secs
  Requests/sec: 247.8679

  Total data:   53000 bytes
  Size/request: 53 bytes

Latency distribution:
  10%% in 0.2002 secs
  25%% in 0.2004 secs
  50%% in 0.2006 secs
  75%% in 0.2010 secs
  90%% in 0.2013 secs
  95%% in 0.2167 secs
  99%% in 0.2210 secs

Status code distribution:
  [200] 1000 responses


---hey -m POST -H "Content-Type: application/json" -d "{\"msg\":\"hello\"}" -n 5000 -c 50 http://localhost:8080/EchoRequestHandler

Summary:
  Total:        0.3579 secs
  Slowest:      0.0261 secs
  Fastest:      0.0001 secs
  Average:      0.0035 secs
  Requests/sec: 13969.0479

  Total data:   250000 bytes
  Size/request: 50 bytes


Latency distribution:
  10%% in 0.0030 secs
  25%% in 0.0031 secs
  50%% in 0.0033 secs
  75%% in 0.0036 secs
  90%% in 0.0040 secs
  95%% in 0.0042 secs
  99%% in 0.0168 secs

Status code distribution:
  [200] 5000 responses


---总结
接口	                            并发数	 请求数	  平均延迟	QPS	  成功率	 结论
HealthHandler	                    100	    10000	   6.8ms	  14567	100%	  极高性能
SlowHandler(ms=200)	              50	    1000	   201.6ms	247	  100%	  符合预期（受限于人工延迟）
EchoRequestHandler(POST JSON)	    50	    5000	   3.5ms	  13969	100%	  极高性能


问题 
1.手动嵌套洋葱在中间件多了之后就不好写了而且不能可插拔，所以直接采用 倒序循环包装
