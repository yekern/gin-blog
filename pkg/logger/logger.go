package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog() {
	encodingConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "Logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stack",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encodingConfig)
	chunkLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("logs/%s", viper.GetString("log.filename")),
		MaxSize:    viper.GetInt("log.max_size"),
		MaxBackups: viper.GetInt("log.max_number"),
		MaxAge:     viper.GetInt("log.expired_date"),
		Compress:   viper.GetBool("log.compress"),
	}
	writeSync := zapcore.AddSync(chunkLogger)
	var l zapcore.Level
	err := l.UnmarshalText([]byte("debug"))
	if err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, writeSync, l)
	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
}

func Record() *zap.Logger {
	return zap.L()
}
