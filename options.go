package log

import (
	"go.uber.org/zap/zapcore"
)

const consoleFormat = "console"

var defaultOpts = newOptions()

// RotateOptions for log rotate feature.
type RotateOptions struct {
	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool
}

// Options for logger.
type Options struct {
	*RotateOptions

	// logger name, default ""
	Name              string
	Level             string
	Format            string
	DisableColor      bool
	DisableCaller     bool
	DisableStacktrace bool
	OutputPaths       []string
	ErrorOutputPaths  []string
	EnableRotate      bool
}

// Return default log options instance.
func newOptions() *Options {
	return &Options{
		RotateOptions: &RotateOptions{
			MaxSize:    100,
			MaxAge:     0,
			MaxBackups: 0,
			LocalTime:  true,
			Compress:   false,
		},
		Name:              "",
		Level:             zapcore.InfoLevel.String(),
		Format:            consoleFormat,
		DisableColor:      false,
		DisableCaller:     false,
		DisableStacktrace: false,
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		EnableRotate: false,
	}
}
