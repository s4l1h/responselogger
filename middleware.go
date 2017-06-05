package responselogger

import "net/http"

// Middleware http.Handler middleware
func Middleware(h http.Handler, fn CallBackFN) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseLogger := New(w, fn)
		h.ServeHTTP(responseLogger, r)
	})
}
