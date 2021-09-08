package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	LoggerConfig zap.Config
}

// NewOptions 自定义Options
func NewOptions() Options {
	options := Options{}
	options.withLoggerOptions()
	options.withEncoderOptions()
	return options
}

//withLoggerOptions 设置Logger日志配置项
func (o *Options) withLoggerOptions() {

	o.LoggerConfig = zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         consoleFormat,
		OutputPaths:      []string{"stdout", "example.log"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

//withLoggerOptions 设置Encoder编码器配置项
func (o *Options) withEncoderOptions() {

	o.LoggerConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// NewDevelopmentConfig 源于zap.NewDevelopmentConfig()
func NewDevelopmentConfig() Options {
	options := Options{}
	options.LoggerConfig = zap.NewDevelopmentConfig()
	options.LoggerConfig.EncoderConfig.EncodeTime = timeEncoder
	options.LoggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return options
}

// NewProductionConfig 源于zap.NewProductionConfig()
func NewProductionConfig() Options {
	options := Options{}
	options.LoggerConfig = zap.NewProductionConfig()
	options.LoggerConfig.EncoderConfig.EncodeTime = timeEncoder
	options.LoggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return options
}
