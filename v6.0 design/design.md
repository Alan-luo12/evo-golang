## V6 设计决策

1.设计了 workers 到 dispatchpool 到 jobqueuechannel 和 processcurrency 的一套限制措施  实现了  对redis的dequeue的保护，以及对sqlite的保护，当currency设置为1的时候可以发现很大减少了sqlite被锁的情况，但是一旦增加为2，就会在并发场景下出现大量lock


2设计了限流桶，不使用传统channel令牌桶，因为精度查资源消耗搞，采用纯计算令牌桶