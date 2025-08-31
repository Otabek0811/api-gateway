package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zap.SugaredLogger
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string) *Logger {
	var lvl zapcore.Level
	switch strings.ToLower(level) {
	case "error":
		lvl = zapcore.ErrorLevel
	case "warn":
		lvl = zapcore.WarnLevel
	case "info":
		lvl = zapcore.InfoLevel
	case "debug":
		lvl = zapcore.DebugLevel
	default:
		lvl = zapcore.InfoLevel
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		lvl,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

	return &Logger{logger: logger}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.log("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info(message)
	} else {
		l.logger.Infof(message, args...)
	}
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Warn(message)
	} else {
		l.logger.Warnf(message, args...)
	}
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.log("error", message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.log("fatal", message, args...)
	os.Exit(1)
}

func (l *Logger) log(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.logFormatted(level, msg.Error(), args...)
	case string:
		l.logFormatted(level, msg, args...)
	default:
		l.logFormatted(level, fmt.Sprintf("%s message: %+v", level, message), args...)
	}
}

func (l *Logger) logFormatted(level string, message string, args ...interface{}) {
	switch level {
	case "debug":
		if len(args) == 0 {
			l.logger.Debug(message)
		} else {
			l.logger.Debugf(message, args...)
		}
	case "info":
		if len(args) == 0 {
			l.logger.Info(message)
		} else {
			l.logger.Infof(message, args...)
		}
	case "warn":
		if len(args) == 0 {
			l.logger.Warn(message)
		} else {
			l.logger.Warnf(message, args...)
		}
	case "error":
		if len(args) == 0 {
			l.logger.Error(message)
		} else {
			l.logger.Errorf(message, args...)
		}
	case "fatal":
		if len(args) == 0 {
			l.logger.Fatal(message)
		} else {
			l.logger.Fatalf(message, args...)
		}
	}
}
