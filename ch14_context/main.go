package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	var result int
	var total int
	ctx := r.Context()
	for {
		select {
		case <-ctx.Done():
			Log(ctx, Debug, "timeout")
			fmt.Fprintf(w, "Sum: %d, iterations: %d, %v\n", result, total, ctx.Err())
			return
		default:
			num := rand.Intn(100_000)
			if num == 1234 {
				Log(ctx, Info, "number found")
				fmt.Fprintf(w, "Sum: %d, iterations: %d, found 1,234\n", result, total)
				return
			}
			result += num
			total++
		}
	}
}

func generateTimeOutMiddleware(duration time.Duration) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancelFunc := context.WithTimeout(r.Context(), duration)
			defer cancelFunc()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func getMux() http.Handler {
	mux := http.NewServeMux()

	timeOut := generateTimeOutMiddleware(time.Second)
	mux.Handle("/", timeOut(http.HandlerFunc(home)))

	return loggigMiddleware(mux)
}

func main() {
	http.ListenAndServe(":4000", getMux())
}
