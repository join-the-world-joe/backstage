package log

import "backstage/abstract/logger"

var _logger logger.Logger
var _level = logger.Debug

func SetLogger(l logger.Logger) {
	_logger = l
}

func SetLevel(level string) {
	switch level {
	case logger.DebugLevel:
		_level = logger.Debug
	case logger.InfoLevel:
		_level = logger.Info
	case logger.WarnLevel:
		_level = logger.Warn
	case logger.ErrorLevel:
		_level = logger.Error
	default:
		_level = logger.Debug
	}
}

func Debug(v ...interface{}) {
	if _level <= logger.Debug {
		_logger.Log(logger.Debug, v...)
	}
}

func DebugF(format string, v ...interface{}) {
	if _level <= logger.Debug {
		_logger.Logf(logger.Debug, format, v...)
	}
}

func Info(v ...interface{}) {
	if _level <= logger.Info {
		_logger.Log(logger.Info, v...)
	}
}

func InfoF(format string, v ...interface{}) {
	if _level <= logger.Info {
		_logger.Logf(logger.Info, format, v...)
	}
}

func Warn(v ...interface{}) {
	if _level <= logger.Warn {
		_logger.Log(logger.Warn, v...)
	}
}

func WarnF(format string, v ...interface{}) {
	if _level <= logger.Warn {
		_logger.Logf(logger.Warn, format, v...)
	}
}

func Error(v ...interface{}) {
	if _level <= logger.Error {
		_logger.Log(logger.Error, v...)
	}
}

func ErrorF(format string, v ...interface{}) {
	if _level <= logger.Error {
		_logger.Logf(logger.Error, format, v...)
	}
}
