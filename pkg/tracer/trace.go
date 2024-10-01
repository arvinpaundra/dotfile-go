package tracer

import (
	config "github.com/arvinpaundra/dotfile-go/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
	Err error
)

func InitLogger() {
	isDevelopment := config.C.GinMode != gin.ReleaseMode

	cfg := zap.Config{
		Development:      isDevelopment,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevel(),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			CallerKey:      "caller",
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	Log, Err = cfg.Build()
	if Err != nil {
		panic(Err)
	}
	defer Log.Sync()
}
