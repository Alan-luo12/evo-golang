整体架构
text
┌─────────────────────────────────────────────────────────────┐
│                         Client                               │
│  1. 生成 timestamp + nonce                                   │
│  2. 计算 signature = HMAC-SHA256(method+path+ts+nonce+body)  │
│  3. 发送请求（携带 X-Sign, X-TimeStamp, X-Nonce）            │
└─────────────────────────┬───────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                      API Gateway / Server                    │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              Signature Middleware                     │    │
│  │  1. 校验必填头 (X-Sign, X-TimeStamp, X-Nonce)        │    │
│  │  2. 校验 Nonce 长度 (8-128)                          │    │
│  │  3. 校验时间戳窗口 (默认 60s)                        │    │
│  │  4. Redis SETNX 防重放                               │    │
│  │  5. 重新计算签名并比对                               │    │
│  └─────────────────────────────────────────────────────┘    │
│                          │                                   │
│                          ▼                                   │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐                  │
│  │  Limiter │→│   Auth   │→│  Handler │                  │
│  └──────────┘  └──────────┘  └──────────┘                  │
└─────────────────────────────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                         Redis                                │
│  Key: security:nounce:{nonce}                               │
│  Value: "1"                                                  │
│  TTL: SignWindow (秒)                                       │
└─────────────────────────────────────────────────────────────┘
1.2 核心代码结构
go
// 签名生成
func SignHMACSHA256(payload string, secret string) string {
    mac := hmac.New(sha256.New, []byte(secret))
    mac.Write([]byte(payload))
    return hex.EncodeToString(mac.Sum(nil))
}

// 防重放（原子操作）
func (r *RedisRepo) UseNonceOnce(ctx context.Context, nonce string, ttl time.Duration) (bool, error) {
    key := fmt.Sprintf("security:nounce:%s", nonce)
    return r.redisClient.SetNX(ctx, key, "1", ttl).Result()
}
1.3 中间件顺序
text
Request → Recover → Trace → Log → Limiter → Sign → Auth → Handler → Response
优先级：Panic 捕获优先 → 链路追踪 → 日志 → 限流 → 签名 → 认证