package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjackV2 "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
	"xl-common/pkg/conf"
)

const (
	Console = "console"
	File    = "file"
)

var (
	Leavel = zap.DebugLevel
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func Init(c conf.LoggerConf) {
	Leavel = getLoggerLevel(c.Level)
	if len(c.Model) == 0 {
		c.Model = Console
	}
	var writeSyncer zapcore.WriteSyncer
	if c.Model == Console {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	} else if c.Model == File {
		if len(c.FileName) == 0 {
			c.FileName = "./logs/im.log"
		}
		if c.MaxSize == 0 {
			c.MaxSize = 1024 * 2
		}
		if c.MaxBackups == 0 {
			c.MaxBackups = 10
		}
		if c.MaxAge == 0 {
			c.MaxAge = 7
		}
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&lumberjackV2.Logger{
			Filename:   c.FileName,
			MaxSize:    c.MaxSize, // megabytes
			MaxBackups: c.MaxBackups,
			MaxAge:     c.MaxAge, // days
		}))
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		writeSyncer,
		Leavel,
	)
	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}

func getLoggerLevel(level string) zapcore.Level {
	var loggerLevel zapcore.Level
	switch level {
	case "error":
		loggerLevel = zapcore.ErrorLevel
	case "debug":
		loggerLevel = zapcore.DebugLevel
	case "info":
		loggerLevel = zapcore.PanicLevel
	case "warn":
		loggerLevel = zapcore.WarnLevel
	case "panic":
		loggerLevel = zapcore.PanicLevel
	case "fatal":
		loggerLevel = zapcore.FatalLevel
	case "dPanic":
		loggerLevel = zapcore.DPanicLevel
	default:
		loggerLevel = zapcore.InfoLevel
	}
	return loggerLevel
}