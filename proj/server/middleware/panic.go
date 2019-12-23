package middleware

import (
	"fmt"
	"go-training/proj/server"
	"net/http"
)

func ServerErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// Add logs
				fmt.Println(r)
				server.WriteErrorResponse(w, "Something went wrong", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
