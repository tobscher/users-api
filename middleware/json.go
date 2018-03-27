package middleware

import "net/http"

// ContentType is a middleware to set the content type.
func ContentType(h http.Handler, contentType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		h.ServeHTTP(w, r)
	})
}
