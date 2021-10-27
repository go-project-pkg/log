package log

import (
	"testing"
)

func Test_Init(t *testing.T) {
	defer std.Sync()

	opts := &Options{
		Format: "console",
	}

	Init(opts)

	std.Info("Hello world!\n")
	std.Info("Hello world!", String("key", "value"), Int("key", 2))

	std.Infof("Hello world!")
	std.Infof("Hello %s!\n", "levin")

	std.Infow("Hello world!")
	std.Infow("Hello ", String("key", "value"))
	std.Infow("Hello ", "key", "value")
	std.Infow("Hello world!", "foo", "bar")

	std.WithName("l1").Infow("nihao", "key", "value")
	std.WithName("l2").Infow("nihao", "key", "value")
	std.WithName("l3").Infow("nihao", "key", "value")

	std.WithFields(String("nihao", "xxx")).Infow("Hello world!", "bar", "foo")
	std.WithFields(String("jin", "ying"), Int("wang", 6)).Infow("Hello world!", "bar", "foo")

	logger1 := std.WithName("levin").WithFields(String("key1", "value1"))
	logger1.Info("info message")
	logger1.Warn("warning message")
}
