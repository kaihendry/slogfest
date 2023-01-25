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
	// log with compact timestamps
	logger := slog.New(slog.NewTextHandler(os.Stdout))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sleepQuery := r.URL.Query().Get("sleep")
		sleep, _ := strconv.Atoi(sleepQuery)

		logger.Info("sleeping", "sleep", sleep)
		time.Sleep(time.Duration(sleep) * time.Second)
		fmt.Fprintf(w, "Slept for %d seconds", sleep)

		logger.Info("slept", "sleep", sleep, "duration", time.Since(start))

		slog.Info("finished",
			slog.Group("req",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String())),
			slog.Int("status", http.StatusOK),
			slog.Duration("duration", time.Since(start)))

	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
