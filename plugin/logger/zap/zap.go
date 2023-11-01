package zap

import (
	"backstage/abstract/logger"
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"strings"
	"time"
)

const (
	Name                = "Zap Logger"
	DefaultFileName     = "default"
	DefaultCallerSkip   = 1
	DefaultMaxAge       = time.Hour * 24 * 30
	DefaultRotationTime = time.Hour * 24
	DefaultLevel        = zapcore.DebugLevel
)

type _Logger struct {
	opts       *Options
	logger     *zap.SugaredLogger
	ioWriter   *rotatelogs.RotateLogs
	fileName   string
	callerSkip int
}

func NewLogger(opts ...Option) (logger.Logger, error) {

	clock := rotatelogs.Clock(rotatelogs.Local)
	level := DefaultLevel // -1, Debug; 0, Info; 1, Warn; 2, Error
	maxAge := DefaultMaxAge
	callerSkip := DefaultCallerSkip
	fileName := DefaultFileName
	rotationTime := DefaultRotationTime

	options := Options{level: DefaultLevel}

	for _, o := range opts {
		o(&options)
	}

	if options.fileName != "" {
		fileName = filepath.Join(options.filePath, options.fileName)
	}

	if options.callerSkip > 0 {
		callerSkip = options.callerSkip
	}

	if options.level > zapcore.DebugLevel && options.level <= zapcore.ErrorLevel {
		level = options.level
	}

	if options.rotationTime.Milliseconds() > 0 {
		rotationTime = options.rotationTime
	}

	if options.maxAge.Milliseconds() > 0 {
		maxAge = options.maxAge
	}

	if options.clock != nil {
		fmt.Println("options.clock != nil")
		clock = options.clock
	}

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "call",
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	zLvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	ioWriter, err := rotatelogs.New(
		strings.Replace(fileName, ".log", "", -1)+"-%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithClock(clock),
	)
	if err != nil {
		return nil, err
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(ioWriter), zLvl),
	)

	return &_Logger{
		opts:     &options,
		ioWriter: ioWriter,
		logger:   zap.New(core, zap.AddCaller(), zap.AddCallerSkip(callerSkip)).Sugar(),
	}, nil
}

func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(filepath.Base(caller.FullPath()))
}

func (p *_Logger) Name() string {
	return Name
}

func (p *_Logger) Log(level logger.Level, v ...interface{}) {
	switch level {
	case logger.Info:
		p.logger.Info(v)
	case logger.Debug:
		p.logger.Debug(v)
	case logger.Warn:
		p.logger.Warn(v)
	case logger.Error:
		p.logger.Error(v)
	}
}

func (p *_Logger) Logf(level logger.Level, format string, v ...interface{}) {
	switch level {
	case logger.Debug:
		p.logger.Debugf(format, v...)
	case logger.Info:
		p.logger.Infof(format, v...)
	case logger.Warn:
		p.logger.Warnf(format, v...)
	case logger.Error:
		p.logger.Errorf(format, v...)
	}
}

func (p *_Logger) Destroy() {
	p.logger.Sync()
	p.ioWriter.Close()
}
