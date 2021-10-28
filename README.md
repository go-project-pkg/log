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
    log.Info("Hello ", String("string_key", "value"), Int("int_key", 666))
    log.Infof("Hello %s!", "world")
    log.Infow("Hello ", "string_key", "value", "int_key", 666)

    log.WithName("logger1").Warn("I am logger1")
    log.WithName("logger2").Warn("I am logger2")

    log.WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")
    log.WithName("logger3").WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")
}
```

Custom your own logger:

```go
import "github.com/go-project-pkg/log"

func main() {
    defer log.Sync()

    opts := &log.Options{
        RotateOptions: &log.RotateOptions{ // useful if EnableRotate is true
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

    log.Init(opts)

    log.Info("Hello world!")
    log.Info("Hello ", String("string_key", "value"), Int("int_key", 666))
    log.Infof("Hello %s!", "world")
    log.Infow("Hello ", "string_key", "value", "int_key", 666)

    log.WithName("logger1").Warn("I am logger1")
    log.WithName("logger2").Warn("I am logger2")

    log.WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")
    log.WithName("logger3").WithFields(String("f1", "value"), Int("f2", 888)).Error("Hello world!")

    // log file rotate test
    for i := 0; i <= 20000; i++ {
      	log.Infof("hello world: %d", i)
    }
}
```

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
