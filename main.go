package main

import (
	"net/http"
	"os"

	"github.com/kaihendry/slogd"
	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout)))
	url := "https://httpbin.org/delay/2"
	var err error
	defer slogd.New("fetching", "url", url).Stop(err)
	_, err = http.Get(url)
}
