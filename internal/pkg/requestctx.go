package pkg

import (
	"context"
)

type CtxKey string

const (
	TraceIDKey     CtxKey = "traceid"
	AuthSubjectKey CtxKey = "authsubject"
)

// 设置traceid
func WithTraceID(ctx context.Context, traceid string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceid)
}

// 获取traceid
func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if !ok {
		return ""
	}
	return traceID
}

// 设置authsubject
func WithAuthSubject(ctx context.Context, authsubject string) context.Context {
	return context.WithValue(ctx, AuthSubjectKey, authsubject)
}

// 获取authsubject
func GetAuthSubject(ctx context.Context) string {
	authsubject, ok := ctx.Value(AuthSubjectKey).(string)
	if !ok {
		return ""
	}
	return authsubject
}
