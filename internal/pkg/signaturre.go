package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
	"time"
)

// 这里使用strings.Builder去实现对payload的拼装把所有重要的信息都拼装在一起
func BuildSignPayload(method string, path string, timestamp int64, nonce string, body []byte) string {
	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString(path)
	sb.WriteString(strconv.FormatInt(timestamp, 10))
	sb.WriteString(nonce)
	sb.Write(body)
	return sb.String()
}

// 返回一个本地根据传入的payload和secret生成的转换位stirng类型的哈希值
func SignHMACSHA256(payload string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	// 转换为十六进制字符串
	return hex.EncodeToString(mac.Sum(nil))
}

// 通过hmac提供的固定时间的安全比对方法，对传入的sign还有本地生成的sign进行比对，返回布尔值
func VertifySign(payload string, secret string, sign string) bool {
	//1111111
	log.Printf("[DEBUG] secret=%q", secret)
	log.Printf("[DEBUG] payload=%q", payload)
	expected := SignHMACSHA256(payload, secret)
	log.Printf("expected sign %s, sign %s", expected, sign)
	//11111111
	return hmac.Equal([]byte(expected), []byte(strings.TrimSpace(sign)))
}

// 通过Unix方法把整数timestamp转湖岸为time.TIme进行计算diff，比对是否在时间窗口之内，作用是防止请求端拿着过期的timestamp进行攻击
func TimeStampInWindow(now time.Time, ts int64, window time.Duration) bool {
	timestamp := time.Unix(ts, 0)
	diff := now.Sub(timestamp)
	if diff <= 0 {
		log.Printf("timestamp %d is in window %s", ts, window)
		return false
	}

	return diff <= window
}

// 计算ttl
func TimeStampTTL(now time.Time, ts int64, window time.Duration) time.Duration {
	timestamp := time.Unix(ts, 0)
	ttl := window - now.Sub(timestamp)
	if ttl <= 0 {
		ttl = 0
	}

	return ttl
}

// 判断nonce是否合法
func ValidNonce(nonce string) bool {
	l := len(strings.TrimSpace(nonce))
	return l >= 8 && l <= 128
}
