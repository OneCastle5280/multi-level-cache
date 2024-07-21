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
		logger *Logger
	}

	// DefaultLogger
	// @Description: 默认日志打印组件
	DefaultLogger struct {
	}

	// LogLevel 日志等级
	LogLevel int
)

func (d DefaultLogger) Debug(format string, v ...any) {
	fmt.Printf(format, v)
}

func (d DefaultLogger) Info(format string, v ...any) {
	fmt.Printf(format, v)
}

func (d DefaultLogger) Warn(format string, v ...any) {
	fmt.Printf(format, v)
}

func (d DefaultLogger) Error(format string, v ...any) {
	fmt.Printf(format, v)
}

var (
	// 日志打印等级
	logLevel LogLevel
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

func Debug(format string, v ...any) {

}
func Info(format string, v ...any) {

}
func Warn(format string, v ...any) {

}
func Error(format string, v ...any) {

}
