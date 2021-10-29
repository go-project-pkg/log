package log

import (
	"context"
	"testing"
)

func Test_Init(t *testing.T) {
	defer std.Sync()

	opts := &Options{
		// Take effect when EnableRotate is true.
		RotateOptions: &RotateOptions{
			// Maximum size in megabytes of the log file before it gets rotated.
			// Default: 100, if the value is 0, the log files will not be rotated.
			MaxSize: 1,
			// Saved days, default 0, means no limit.
			MaxAge: 30,
			// Saved count, default 0, means no limit.
			MaxBackups: 2,
			// Use local time in log file name, default false.
			LocalTime: true,
			// Gzip log files, default false.
			Compress: false,
		},
		Name:              "",        // logger name
		Level:             "debug",   // debug, info, warn, error, dpanic, panic, fatal
		Format:            "console", // json, console/text
		DisableColor:      false,
		DisableCaller:     false,
		DisableStacktrace: false,
		// Aplication's all levels logs.
		OutputPaths: []string{
			"stdout", // os.Stdout
			"./app.log",
		},
		// Only include zap internal errors, not include application's any level logs.
		ErrorOutputPaths: []string{
			"stderr", // os.Stderr
			"./error.log",
		},
		// Enable log files rotation feature or not.
		EnableRotate: true,
	}

	Init(opts)

	std.Debug("Hello world!")
	std.Debug("Hello ", String("string_key", "value"), Int("int_key", 666))
	std.Debugf("Hello %s!", "world")
	std.Debugw("Hello ", "string_key", "value", "int_key", 666)

	println()

	std.Info("Hello world!")
	std.Info("Hello ", String("string_key", "value"), Int("int_key", 666))
	std.Infof("Hello %s!", "world")
	std.Infow("Hello ", "string_key", "value", "int_key", 666)

	println()

	std.WithName("logger1").Warn("I am logger1")
	std.WithName("logger2").Warn("I am logger2")

	println()

	std.WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")
	std.WithName("logger3").WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")

	println()

	//std.DPanic("dpanic message")
	//std.Panic("panic message")
	//std.Fatal("fatal message")

	// log files rotation test
	for i := 0; i <= 20000; i++ {
		std.Infof("hello world: %d", i)
	}
}

func Test_ToContext_FromContext(t *testing.T) {
	defer std.Sync()

	ctx := std.WithFields(String("f1", "value"), Int("f2", 888)).ToContext(context.Background())
	std.FromContext(ctx).Info("hello world!")
}
