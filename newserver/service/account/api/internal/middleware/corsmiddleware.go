package middleware

import "net/http"

type CorsmiddlewareMiddleware struct {
}

func NewCorsmiddlewareMiddleware() *CorsmiddlewareMiddleware {
	return &CorsmiddlewareMiddleware{}
}

func (m *CorsmiddlewareMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
