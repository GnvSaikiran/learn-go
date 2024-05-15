package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

func logIPAdress(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		options := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler := slog.NewJSONHandler(os.Stderr, options)
		mySlog := slog.New(handler)

		ip, _, _ := strings.Cut(r.RemoteAddr, ":")
		mySlog.Debug("request info", "IP", ip)

		h.ServeHTTP(w, r)
	})
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	acceptHeader := r.Header.Get("Accept")
	if acceptHeader == "application/json" {
		tStruct := struct {
			DayOfWeek  string `json:"day_of_week"`
			DayOfMonth int    `json:"day_of_month"`
			Month      string `json:"month"`
			Year       int    `json:"year"`
			Hour       int    `json:"hour"`
			Minute     int    `json:"minute"`
			Second     int    `json:"second"`
		}{
			DayOfWeek:  t.Weekday().String(),
			DayOfMonth: t.Day(),
			Month:      t.Month().String(),
			Year:       t.Year(),
			Hour:       t.Hour(),
			Minute:     t.Minute(),
			Second:     t.Second(),
		}

		jsonTime, _ := json.Marshal(tStruct)
		w.Write(jsonTime)
		return
	}

	w.Write([]byte(t.Format(time.RFC3339)))
}

func getServerMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /time", handleTime)

	return logIPAdress(mux)
}

func startServer() {
	http.ListenAndServe(":4000", getServerMux())
}
