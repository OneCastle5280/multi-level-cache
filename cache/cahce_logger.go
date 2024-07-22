package cache

import "fmt"

type (

	// Logger
	// @Description: 日志打印组件接口
	Logger interface {
		Debug(format string, v ...any)
		Info(format string, v ...any)
		Warn(format string, v ...any)
		Error(format string, v ...any)
	}

	// MlcLogger
	// @Description: Logger 实现类
	MlcLogger struct {
		logger Logger
	}

	// DefaultLogger
	// @Description: 默认日志打印组件
	DefaultLogger struct {
	}

	// LogLevel 日志等级
	LogLevel int
)

func (d DefaultLogger) Debug(format string, v ...any) {
	fmt.Printf(format, v...)
}

func (d DefaultLogger) Info(format string, v ...any) {
	fmt.Printf(format, v...)
}

func (d DefaultLogger) Warn(format string, v ...any) {
	fmt.Printf(format, v...)
}

func (d DefaultLogger) Error(format string, v ...any) {
	fmt.Printf(format, v...)
}

var (
	// 日志打印等级
	logLevel LogLevel
	// mlc logger
	mlcLogger = MlcLogger{
		logger: &DefaultLogger{},
	}
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// SetLoggerLevel
//
//	@Description: 设置日志打印等级
//	@param level
func SetLoggerLevel(level LogLevel) {
	if level < DEBUG || level > ERROR {
		panic("log level illegal")
	}
	logLevel = level
}

// SetLogger
//
//	@Description: 设置自定义 Logger
//	@param logger
func SetLogger(logger Logger) {
	if logger == nil {
		panic("logger is nil")
	}
	mlcLogger.logger = logger
}

// Debug
//
//	@Description: debug
//	@param format
//	@param v
func Debug(format string, v ...any) {
	if logLevel > DEBUG {
		return
	}
	mlcLogger.logger.Debug(format, v...)
}

// Info
//
//	@Description: info
//	@param format
//	@param v
func Info(format string, v ...any) {
	if logLevel > INFO {
		return
	}
	mlcLogger.logger.Info(format, v...)
}

// Warn
//
//	@Description: warn
//	@param format
//	@param v
func Warn(format string, v ...any) {
	if logLevel > WARN {
		return
	}
	mlcLogger.logger.Info(format, v...)
}

// Error
//
//	@Description: error
//	@param format
//	@param v
func Error(format string, v ...any) {
	if logLevel > ERROR {
		return
	}
	mlcLogger.logger.Info(format, v...)
}
