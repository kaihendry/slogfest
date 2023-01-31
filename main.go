package main

import (
	"os"
	"time"

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
	path := "/mnt/redsamba/freenas.mp4"
	defer Trace("opening").Stop(&err)
	_, err = os.Open(path)
	return
}

// Just want to entend slog.Record with a Stop function
type traceEntry struct {
	r     slog.Record
	start time.Time
}

func Trace(msg string, kvs ...any) (v traceEntry) {
	slog.Info(msg, kvs...)
	v.r.Message = msg
	v.start = time.Now()
	return v
}

func (v traceEntry) Stop(err *error) {
	if err == nil || *err == nil {
		slog.Info(v.r.Message, "duration", time.Since(v.start).Milliseconds())
	} else {
		slog.Error(v.r.Message, *err, "duration", time.Since(v.start).Milliseconds())
	}
}
