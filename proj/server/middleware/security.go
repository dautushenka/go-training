package middleware

import (
	"context"
	"go-training/proj/core/repository"
	"go-training/proj/core/security"
	"go-training/proj/server"
	"net/http"
	"strings"
)

type Security struct {
	UserRepository repository.UserRepository
}

func (s *Security) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authentication")
		if bearerToken != "" {
			tokenSlice := strings.Split(bearerToken, "Bearer ")
			if len(tokenSlice) > 1 {
				token := tokenSlice[1]
				userId, err := security.ValidateToken(token)
				if err == nil {
					user, err := s.UserRepository.GetUserById(userId)
					if err == nil {
						ctx := context.WithValue(r.Context(), "User", user)
						r = r.WithContext(ctx)
					}
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Security) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := r.Context().Value("User"); user == nil {
			server.WriteErrorResponse(w, "Access Denied", 403)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
