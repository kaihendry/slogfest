package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	http.HandleFunc("/foo", sleep())
	http.HandleFunc("/bar", basic())

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func sleep() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sleepQuery := r.URL.Query().Get("sleep")
		sleep, _ := strconv.Atoi(sleepQuery)

		slog.Info("sleeping", "sleep", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)
		fmt.Fprintf(w, "Slept for %d seconds", sleep)

		slog.Info("slept", "sleep", sleep, "duration", time.Since(start))

		slog.Info("finished",
			slog.Group("req",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String())),
			slog.Int("status", http.StatusOK),
			slog.Duration("duration", time.Since(start)))
	}
}

func basic() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		slog.Info("finished",
			slog.Group("req",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String())),
			slog.Int("status", http.StatusOK),
			slog.Duration("duration", time.Since(start)))
	}
}
