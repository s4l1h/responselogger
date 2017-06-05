package responselogger

import (
	"net/http"
)

// func exampleCallBackFn(b []byte){
// 	log.Print(string(b))
// }

// CallBackFN responselogger callback function
type CallBackFN func([]byte) []byte

// New create and return new Response Logger
func New(w http.ResponseWriter, fn CallBackFN) *ResponseLogger {
	return &ResponseLogger{w: w, fn: fn}
}

// ResponseLogger object
type ResponseLogger struct {
	w  http.ResponseWriter
	fn CallBackFN
}

// Header http.ResponseWriter implement
func (r *ResponseLogger) Header() http.Header {
	return r.w.Header()
}

// WriteHeader http.ResponseWriter implement
func (r *ResponseLogger) WriteHeader(code int) {
	r.w.WriteHeader(code)
}

// Write http.ResponseWriter implement
func (r *ResponseLogger) Write(b []byte) (int, error) {
	return r.w.Write(r.fn(b)) // call fn and write
}
