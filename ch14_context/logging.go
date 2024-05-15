package main

import (
	"context"
	"fmt"
	"net/http"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

type userKey struct{}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level

	inLevel, ok := LogLevelFromContext(ctx)
	if !ok {
		fmt.Println("error reading log level from context")
		return
	}

	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Info || inLevel == Debug) {
		fmt.Println(message)
	}
}

func ContextWithLogLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, userKey{}, level)
}

func LogLevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(userKey{}).(Level)
	return level, ok
}

func loggigMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := r.URL.Query().Get("log_level")
		ctx := ContextWithLogLevel(r.Context(), Level(level))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
