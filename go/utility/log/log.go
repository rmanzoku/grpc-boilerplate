package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Logger struct {
	*zap.Logger
}

var logger *Logger

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	l, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	logger = &Logger{l}
}
func GetLogger() *Logger {
	return logger
}

func Print(args ...interface{}) {
	logger.Sugar().Info(args)
}

func isIgnoreResponseLoggingMethod(s string) bool {
	list := []string{
		"/auth.AuthService/GetNow",
	}
	for _, svc := range list {
		if s == svc {
			return true
		}
	}
	return false
}

func AccessLog(ctx context.Context, req interface{}, res interface{}, incomingErr error, info *grpc.UnaryServerInfo, nano int64) {

	// uid, err := ctxutil.ExtractUid(ctx)
	// if err != nil {
	// 	uid = 0
	// }

	fields := []zap.Field{
		zap.String("type", "access"),
		zap.Any("req", req),
		zap.Error(incomingErr),
		zap.Float32("time_ms", float32(nano/1000)/1000),
		//		zap.Uint32("uid", uid),
		zap.String("service", info.FullMethod),
	}

	if !isIgnoreResponseLoggingMethod(info.FullMethod) {
		fields = append(fields, zap.Any("res", res))
	} else {
		fields = append(fields, zap.String("res", "ignored"))
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		f := []zap.Field{
			zap.Strings("user-agent", md.Get("user-agent")),
			zap.Strings("ref", md.Get("referer")),
			zap.Strings("lang", md.Get("accept-language")),
			zap.Strings("addr", md.Get("x-forwarded-for")),
		}
		fields = append(fields, f...)
	}

	if incomingErr != nil {
		logger.Error("access", fields...)
	} else {
		logger.Info("access", fields...)
	}
}
