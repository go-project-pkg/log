# go-project-pkg/log

Wrap [zap](https://github.com/uber-go/zap) for easy using.

## Installation

```sh
$ go get -u github.com/go-project-pkg/log
```

## Usage

Use default logger:

```go
import "github.com/go-project-pkg/log"

func main() {
    defer log.Sync()

    log.Info("Hello world!")
    log.Info("Hello ", log.String("string_key", "value"), log.Int("int_key", 666))
    log.Infof("Hello %s!", "world")
    log.Infow("Hello ", "string_key", "value", "int_key", 666)

    log.WithName("logger1").Warn("I am logger1")
    log.WithName("logger2").Warn("I am logger2")

    log.WithFields(log.String("f1", "value"), log.Int("f2", 888)).Error("Hello world!")
    log.WithName("logger3").WithFields(log.String("f1", "value"), log.Int("f2", 888)).Error("Hello world!")
}
```

Custom your own logger:

```go
import "github.com/go-project-pkg/log"

func main() {
    defer log.Sync()

    opts := &log.Options{
        RotateOptions: &log.RotateOptions{ // take effect when EnableRotate is true
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
        EnableRotate: true, // log files rotate or not
    }

    log.Init(opts)

    log.Info("Hello world!")
    log.Info("Hello ", log.String("string_key", "value"), log.Int("int_key", 666))
    log.Infof("Hello %s!", "world")
    log.Infow("Hello ", "string_key", "value", "int_key", 666)

    log.WithName("logger1").Warn("I am logger1")
    log.WithName("logger2").Warn("I am logger2")

    log.WithFields(log.String("f1", "value"), log.Int("f2", 888)).Error("Hello world!")
    log.WithName("logger3").WithFields(log.String("f1", "value"), log.Int("f2", 888)).Error("Hello world!")

    // log files rotation test
    for i := 0; i <= 20000; i++ {
      	log.Infof("hello world: %d", i)
    }
}
```

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
