package middleware

import "net/http"

func MiddlewareCorsHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Content-type, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if (*r).Method == "OPTIONS" {
			return
		}

		handler.ServeHTTP(w, r)
	})
}
