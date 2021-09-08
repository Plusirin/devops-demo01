package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 日志创建方式包括两类： NewLog() 和 NewRotateLog()，区别在于 NewRotateLog() 支持轮转；
//var std = NewLog()
//var std = NewRotateLog()
var std = NewRotateLog()

// 日志配置项：包括两类主环境， NewProductionConfig() 和 NewDevelopmentConfig()，以及自定义配置 NewOptions()
// LogOptions Config: Include NewProductionConfig() and NewDevelopmentConfig(),or NewOptions()
//var options = NewOptions()
//var options = NewDevelopmentConfig()
//var options = NewProductionConfig()
var options = NewDevelopmentConfig()

//NewLog New create a new Log (not support log rotating).
func NewLog() *Log {
	l, err := options.LoggerConfig.Build(
		zap.AddStacktrace(zapcore.PanicLevel),
		zap.AddCallerSkip(2),
	)
	if err != nil {
		panic(err)
	}
	log := &Log{
		base:  l,
		level: zapcore.DebugLevel,
	}

	return log
}

// NewRotateLog New create a new Log (support log rotating).
func NewRotateLog() *Log {
	return NewTeeWithRotate(options)
}

type Log struct {
	base  *zap.Logger
	level Level
}

func (l *Log) Debug(msg string, fields ...Field) {
	l.base.Debug(msg, fields...)
}

func (l *Log) Info(msg string, fields ...Field) {
	l.base.Info(msg, fields...)
}

func (l *Log) Warn(msg string, fields ...Field) {
	l.base.Warn(msg, fields...)
}

func (l *Log) Error(msg string, fields ...Field) {
	l.base.Error(msg, fields...)
}

func (l *Log) DPanic(msg string, fields ...Field) {
	l.base.DPanic(msg, fields...)
}

func (l *Log) Panic(msg string, fields ...Field) {
	l.base.Panic(msg, fields...)
}

func (l *Log) Fatal(msg string, fields ...Field) {
	l.base.Fatal(msg, fields...)
}

func (l *Log) Debugf(template string, args ...interface{}) {
	l.base.Sugar().Debugf(template, args)
}

func (l *Log) Infof(template string, args ...interface{}) {
	l.base.Sugar().Infof(template, args)
}

func (l *Log) Warnf(template string, args ...interface{}) {
	l.base.Sugar().Warnf(template, args)
}

func (l *Log) Errorf(template string, args ...interface{}) {
	l.base.Sugar().Errorf(template, args)
}

func (l *Log) DPanicf(template string, args ...interface{}) {
	l.base.Sugar().DPanicf(template, args)
}

func (l *Log) Panicf(template string, args ...interface{}) {
	l.base.Sugar().Panicf(template, args)
}

func (l *Log) Fatalf(template string, args ...interface{}) {
	l.base.Sugar().Fatalf(template, args)
}

func (l *Log) Debugw(msg string, keysAndValues ...interface{}) {
	l.base.Sugar().Debugw(msg, keysAndValues...)
}

func (l *Log) Infow(msg string, keysAndValues ...interface{}) {
	l.base.Sugar().Infow(msg, keysAndValues...)
}

func (l *Log) Warnw(msg string, keysAndValues ...interface{}) {
	l.base.Sugar().Warnw(msg, keysAndValues...)
}

func (l *Log) Errorw(msg string, keysAndValues ...interface{}) {
	l.base.Sugar().Errorw(msg, keysAndValues...)
}

func (l *Log) DPanicw(msg string, keysAndValues ...interface{}) {
	l.base.Sugar().DPanicw(msg, keysAndValues...)
}

func (l *Log) Panicw(msg string, keysAndValues ...interface{}) {
	l.base.Sugar().Panicw(msg, keysAndValues...)
}

func (l *Log) Fatalw(msg string, keysAndValues ...interface{}) {
	l.base.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *Log) WithValues(keysAndValues ...interface{}) Log {
	newLogger := l.base.With(handleFields(l.base, keysAndValues)...)

	return NewLogger(newLogger)
}

func (l *Log) Write(p []byte) (n int, err error) {
	l.base.Info(string(p))

	return len(p), nil
}

// NewLogger creates a new logr.Logger using the given Zap Logger to log.
func NewLogger(l *zap.Logger) Log {
	return Log{
		base:  l,
		level: zap.InfoLevel,
	}
}

func handleFields(l *zap.Logger, args []interface{}, additional ...zap.Field) []zap.Field {
	// a slightly modified version of zap.SugaredLogger.sweetenFields
	if len(args) == 0 {
		// fast-return if we have no suggared fields.
		return additional
	}

	// unlike Zap, we can be pretty sure users aren't passing structured
	// fields (since logr has no concept of that), so guess that we need a
	// little less space.
	fields := make([]zap.Field, 0, len(args)/2+len(additional))
	for i := 0; i < len(args); {
		// check just in case for strongly-typed Zap fields, which is illegal (since
		// it breaks implementation agnosticism), so we can give a better error message.
		if _, ok := args[i].(zap.Field); ok {
			l.DPanic("strongly-typed Zap Field passed to logr", zap.Any("zap field", args[i]))

			break
		}

		// make sure this isn't a mismatched key
		if i == len(args)-1 {
			l.DPanic("odd number of arguments passed as key-value pairs for logging", zap.Any("ignored key", args[i]))

			break
		}

		// process a key-value pair,
		// ensuring that the key is a string
		key, val := args[i], args[i+1]
		keyStr, isString := key.(string)
		if !isString {
			// if the key isn't a string, DPanic and stop logging
			l.DPanic(
				"non-string key argument passed to logging, ignoring all later arguments",
				zap.Any("invalid key", key),
			)

			break
		}

		fields = append(fields, zap.Any(keyStr, val))
		i += 2
	}

	return append(fields, additional...)
}

func (l *Log) Sync() error {
	return l.base.Sync()
}

func Sync() error {
	if std != nil {
		return std.Sync()
	}
	return nil
}

func Default() *Log {
	return std
}
