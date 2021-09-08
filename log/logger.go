package log

type Logger interface {
	DPanic(msg string, fields ...Field)
	DPanicf(template string, args ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})

	Info(msg string, fields ...Field)
	Infof(format string, v ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Debug(msg string, fields ...Field)
	Debugf(format string, v ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Warn(msg string, fields ...Field)
	Warnf(format string, v ...interface{})
	Warnw(msg string, keysAndValues ...interface{})

	Error(msg string, fields ...Field)
	Errorf(format string, v ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Panic(msg string, fields ...Field)
	Panicf(format string, v ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	Fatal(msg string, fields ...Field)
	Fatalf(format string, v ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Write(p []byte) (n int, err error)
	WithValues(keysAndValues ...interface{}) Logger

	//WithName(name string) Logger

	//WithContext(ctx context.Context) context.Context
}

//
//func Errorw(msg string, keysAndValues ...interface{}) {
//	std.base.Sugar().Errorw(msg, keysAndValues...)
//}
