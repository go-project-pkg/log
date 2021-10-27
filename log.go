package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// global logger instance.
	std = New(defaultOpts)
	mu  sync.Mutex
)

// Init logger with given options instance.
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = New(opts)
}

// New logger
func New(opts *Options) *Logger {
	if opts == nil {
		opts = newOptions()
	}

	if opts.Format == "" {
		opts.Format = defaultOpts.Format
	}

	if len(opts.OutputPaths) == 0 {
		opts.OutputPaths = defaultOpts.OutputPaths
	}

	if len(opts.ErrorOutputPaths) == 0 {
		opts.ErrorOutputPaths = defaultOpts.ErrorOutputPaths
	}

	var level Level
	if err := level.UnmarshalText([]byte(opts.Level)); err != nil {
		level = InfoLevel
	}

	encodeLevel := zapcore.CapitalLevelEncoder
	if opts.Format == consoleFormat && !opts.DisableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "timestamp",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       opts.Development,
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opts.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}

	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	zap.RedirectStdLog(l)

	logger := newLogger(l)

	return logger
}
