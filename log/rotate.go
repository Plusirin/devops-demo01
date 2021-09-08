package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// RotateOptions 定义了日志轮转 配置项，主要包括：
// MaxSize: 日志最大Size;
// MaxAge: 日志最大Age;
// MaxBackups: 日志最大副本;
// Compress: 日志是否压缩;
type RotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type LevelEnablerFunc func(lvl Level) bool

// TeeOption 定义了日志轮转 文件配置项，主要包括：
// Filename: 日志文件名称;
// Ropt: 为 RotateOptions 日志轮转配置项;
// Lef: 定义该文件允许写入的日志级别;
type TeeOption struct {
	Filename string
	Ropt     RotateOptions
	Lef      LevelEnablerFunc
}

// NewTeeWithRotate 将 []TeeOption 通过 NewCore 方法创建 core，
// 并添加到 cores 变量中，通过调用 zap.NewTee 方法同时创建多个输出；
func NewTeeWithRotate(options Options) *Log {
	var cores []zapcore.Core

	// 获取 []TeeOption
	tops := newRotateOptions()

	// 遍历 tops，通过newCore()创建对应core，添加到 cores 返回列表中
	// 添加日志文件输出 core
	for _, top := range tops {
		top := top

		core := newCore(&top, &options)
		// cores 中包含 tops 中的 TeeOption 对象，一般为 access.log 和 error.log
		cores = append(cores, core)
	}

	options.LoggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// 添加os.Stdout输出core
	core := zapcore.NewCore(
		// 编码器配置为 Json 格式
		zapcore.NewConsoleEncoder(options.LoggerConfig.EncoderConfig),
		// 输出同步到标准输出和文件
		zapcore.WriteSyncer(
			//zapcore.AddSync(os.Stdout),
			zapcore.AddSync(os.Stdout),
		),
		// 设置本 core 控制器的日志级别
		DebugLevel,
	)

	cores = append(cores, core)

	logger := &Log{
		// 使用 zap.New() 方法创建 Logger 对象，并赋值给 base;
		base: zap.New(
			zapcore.NewTee(cores...),
			// 添加堆栈调用级别(zap.AddStacktrace()) 和 跳过内部堆栈(zap.AddCallerSkip())
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.PanicLevel),
			zap.AddCallerSkip(1),
		),
		level: zapcore.DebugLevel,
	}
	return logger
}

// newCore 通过 zapcore.NewCore() 初始化 *TeeOption 对象，
func newCore(top *TeeOption, options *Options) zapcore.Core {
	hook := lumberjack.Logger{
		Filename:   top.Filename,
		MaxSize:    top.Ropt.MaxSize,
		MaxBackups: top.Ropt.MaxBackups,
		MaxAge:     top.Ropt.MaxAge,
		Compress:   top.Ropt.Compress,
	}

	// 创建日志 os.Stdout 标准输出 core，并添加到 cores 中
	core := zapcore.NewCore(
		// 编码器配置为 Json 格式
		zapcore.NewJSONEncoder(options.LoggerConfig.EncoderConfig),
		// 输出同步到标准输出和文件
		zapcore.NewMultiWriteSyncer(
			//zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
		),
		// 设置本 core 控制器的日志级别
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(Level(lvl))
		}),
	)

	return core
}

// newRotateOptions 根据结构体中的参数，获取 []TeeOption
func newRotateOptions() []TeeOption {
	rotateOptions := &RotateOptions{
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 3,
		Compress:   true,
	}

	accessOpt := &TeeOption{
		Filename: "logs/access.log",
		Ropt:     *rotateOptions,
		Lef: func(lvl Level) bool {
			return lvl <= InfoLevel
		},
	}

	errorOpt := &TeeOption{
		Filename: "logs/error.log",
		Ropt:     *rotateOptions,
		Lef: func(lvl Level) bool {
			return lvl > InfoLevel
		},
	}

	return []TeeOption{
		*accessOpt,
		*errorOpt,
	}
}
