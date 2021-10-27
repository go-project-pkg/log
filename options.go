package log

import (
	"go.uber.org/zap/zapcore"
)

const consoleFormat = "console"

var defaultOpts = newOptions()

// Options for logger.
type Options struct {
	// logger name, default ""
	Name              string
	Level             string
	Format            string
	DisableColor      bool
	DisableCaller     bool
	DisableStacktrace bool
	Development       bool
	OutputPaths       []string
	ErrorOutputPaths  []string
}

// Return default log options instance.
func newOptions() *Options {
	return &Options{
		Name:              "",
		Level:             zapcore.InfoLevel.String(),
		Format:            consoleFormat,
		DisableColor:      false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Development:       false,
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
}
