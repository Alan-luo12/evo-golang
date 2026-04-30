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
```

4. go语言天然并发的优势？
天然支持并发，自动创建协程理想QPS极高
