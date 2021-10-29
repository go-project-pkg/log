package log

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func Test_Init(t *testing.T) {
	defer Sync()

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

	Debug("Hello world!")
	Debug("Hello ", String("string_key", "value"), Int("int_key", 666))
	Debugf("Hello %s!", "world")
	Debugw("Hello ", "string_key", "value", "int_key", 666)

	println()

	Info("Hello world!")
	Info("Hello ", String("string_key", "value"), Int("int_key", 666))
	Infof("Hello %s!", "world")
	Infow("Hello ", "string_key", "value", "int_key", 666)

	println()

	WithName("logger1").Warn("I am logger1")
	WithName("logger2").Warn("I am logger2")

	println()

	WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")
	WithName("logger3").WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")

	println()

	//DPanic("dpanic message")
	//Panic("panic message")
	//Fatal("fatal message")

	// log files rotation test
	for i := 0; i <= 20000; i++ {
		Infof("hello world: %d", i)
	}
}

func Test_ToContext_FromContext(t *testing.T) {
	defer Sync()

	ctx := WithFields(String("f1", "value"), Int("f2", 888)).ToContext(context.Background())
	FromContext(ctx).Info("hello world!")
}

func Test_SetHooks(t *testing.T) {
	defer Sync()

	monitorHook1 := func(entry Entry) error {
		if entry.Level >= ErrorLevel {
			fmt.Println("hook1 alert!")
		}
		return nil
	}

	monitorHook2 := func(entry Entry) error {
		if entry.Level >= ErrorLevel {
			fmt.Println("hook2 alert!")
		}

		// This error log is zap internal error, and will write to 'ErrorOutputPaths'.
		return errors.New("alert hook failed")
	}

	SetHooks(monitorHook1, monitorHook2)

	Error("set hooks: server failed")
}

func Test_WithHooks(t *testing.T) {
	monitorHook3 := func(entry Entry) error {
		if entry.Level >= ErrorLevel {
			fmt.Println("hook3 alert!")
		}
		return nil
	}

	logger := New(nil)
	defer logger.Sync()

	logger.WithHooks(monitorHook3).Error("with hooks: server failed")

	logger.Error("no hooks: server failed")
}
