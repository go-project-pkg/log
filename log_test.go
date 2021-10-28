package log

import "testing"

func Test_Init(t *testing.T) {
	defer std.Sync()

	opts := &Options{
		RotateOptions: &RotateOptions{ // useful if EnableRotate is true
			MaxSize:    1,     // MB
			MaxAge:     0,     // Day
			MaxBackups: 2,     // saved files count
			LocalTime:  true,  // use local time in log file name
			Compress:   false, // gzip
		},
		Name:              "",        // logger name
		Level:             "debug",   // debug, info, warn, error, panic, dpanic, fatal
		Format:            "console", // json, console/text
		DisableColor:      false,
		DisableCaller:     false,
		DisableStacktrace: false,
		OutputPaths: []string{ // aplication's all levels logs
			"stdout", // os.Stdout
			"./app.log",
		},
		ErrorOutputPaths: []string{ // zap internal errors, not include application's any level logs
			"stderr", // os.Stderr
			"./error.log",
		},
		EnableRotate: true, // if rotate log files
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

	// log file rotate test
	for i := 0; i <= 20000; i++ {
		std.Infof("hello world: %d", i)
	}
}
