package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field
type Level = zapcore.Level

const (
	consoleFormat = "console"
	jsonFormat    = "json"
)

//日志等级
const (
	InfoLevel   = zap.InfoLevel   // 0, default level
	WarnLevel   = zap.WarnLevel   // 1
	ErrorLevel  = zap.ErrorLevel  // 2
	DPanicLevel = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zap.FatalLevel // 5
	DebugLevel = zap.DebugLevel // -1
)

var (
	Info   = std.Info
	Warn   = std.Warn
	Error  = std.Error
	DPanic = std.DPanic
	Panic  = std.Panic
	Fatal  = std.Fatal
	Debug  = std.Debug

	Infow   = std.Infow
	Warnw   = std.Warnw
	Errorw  = std.Errorw
	DPanicw = std.DPanicw
	Panicw  = std.Panicw
	Fatalw  = std.Fatalw
	Debugw  = std.Debugw

	Infof   = std.Infof
	Warnf   = std.Warnf
	Errorf  = std.Errorf
	DPanicf = std.DPanicf
	Panicf  = std.Panicf
	Fatalf  = std.Fatalf
	Debugf  = std.Debugf

	WithValues = std.WithValues
)

// function variables for all field types
// in github.com/uber-go/zap/field.go
var (
	Any         = zap.Any
	Array       = zap.Array
	Object      = zap.Object
	Binary      = zap.Binary
	Bool        = zap.Bool
	Bools       = zap.Bools
	ByteString  = zap.ByteString
	ByteStrings = zap.ByteStrings
	Complex64   = zap.Complex64
	Complex64s  = zap.Complex64s
	Complex128  = zap.Complex128
	Complex128s = zap.Complex128s
	Duration    = zap.Duration
	Durations   = zap.Durations
	Err         = zap.Error
	Errors      = zap.Errors
	Float32     = zap.Float32
	Float32s    = zap.Float32s
	Float64     = zap.Float64
	Float64s    = zap.Float64s
	Int         = zap.Int
	Ints        = zap.Ints
	Int8        = zap.Int8
	Int8s       = zap.Int8s
	Int16       = zap.Int16
	Int16s      = zap.Int16s
	Int32       = zap.Int32
	Int32s      = zap.Int32s
	Int64       = zap.Int64
	Int64s      = zap.Int64s
	Namespace   = zap.Namespace
	Reflect     = zap.Reflect
	Stack       = zap.Stack
	String      = zap.String
	Stringer    = zap.Stringer
	Strings     = zap.Strings
	Time        = zap.Time
	Times       = zap.Times
	Uint        = zap.Uint
	Uints       = zap.Uints
	Uint8       = zap.Uint8
	Uint8s      = zap.Uint8s
	Uint16      = zap.Uint16
	Uint16s     = zap.Uint16s
	Uint32      = zap.Uint32
	Uint32s     = zap.Uint32s
	Uint64      = zap.Uint64
	Uint64s     = zap.Uint64s
	Uintptr     = zap.Uintptr
	Uintptrs    = zap.Uintptrs
)
