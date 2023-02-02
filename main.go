package main

import (
	"os"
	"time"

	"github.com/kaihendry/gotrace"
	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	// https://github.com/apex/log/blob/master/_examples/trace/trace.go

	ctx := slog.Group("fields",
		slog.String("app", "myapp"),
		slog.String("env", "prod"))

	for range time.Tick(time.Second) {
		_ = work(ctx)
	}

}

func work(ctx slog.Attr) (err error) {
	path := "README.md"
	defer gotrace.Trace("opening").Stop(&err)
	_, err = os.Open(path)
	return
}
