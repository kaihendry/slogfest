package main

import (
	"net/http"
	"os"

	"github.com/kaihendry/gotrace"
	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	var err error
	defer gotrace.Trace("fetching", "env", "dev").Stop(&err)
	_, err = http.Get("https://httpbin.org/delay/2")

}
