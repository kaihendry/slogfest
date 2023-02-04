package main

import (
	"net/http"
	"os"

	"github.com/kaihendry/gotrace"
	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	url := "https://httpbin.org/delay/2"
	var err error
	defer gotrace.New("fetching", "url", url).Stop(err)
	_, err = http.Get(url)

}
