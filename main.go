package main

import (
	"net/http"
	"os"

	"github.com/kaihendry/slogd"
	"golang.org/x/exp/slog"
)

func main() {
	opts := slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Value.Kind() == slog.KindDuration {
				a.Value = slog.Int64Value(a.Value.Duration().Milliseconds())
			}
			return a
		},
	}

	slog.SetDefault(slog.New(opts.NewJSONHandler(os.Stdout)))

	good()
	bad()
}

func good() {
	url := "https://httpbin.org/delay/2"
	var err error
	defer slogd.New("fetching", "url", url).Stop(&err)
	_, err = http.Get(url)

}

func bad() {
	url := "http://1.2.3.4"
	var err error
	defer slogd.New("fetching", "url", url).Stop(&err)
	_, err = http.Get(url)

}
