package ctxutil

import (
	"context"
	"time"
)

type key int

const (
	timeKey key = iota
	datastoreKey
)

func Setup(ctx context.Context) context.Context {
	ctx = SetTime(ctx, time.Now().Unix())
	ctx = SetDataStore(ctx)
	return ctx
}

func SetTime(ctx context.Context, time int64) context.Context {
	return context.WithValue(ctx, timeKey, time)
}

func ExtractTime(ctx context.Context) int64 {
	return ctx.Value(timeKey).(int64)
}
