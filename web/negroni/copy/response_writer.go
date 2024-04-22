package copy

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	http.Flusher

	Status() int

	Written() bool

	Size() int

	Before(func (ResponseWriter))
}

type beforeFunc func(writer ResponseWriter)

func NewResponseWriter(rw http.ResponseWriter) ResponseWriter {
	nrw := &responseWriter{
		ResponseWriter: rw,
	}

	if _, ok := rw.(http.CloseNotifier); ok {
		return &responseWriterCloseNotifer{nrw}
	}

	return nrw
}

type responseWriter struct {
	http.ResponseWriter
	status     int
	size       int
	beforeFunc []beforeFunc
}

func (rw *responseWriter) WriteHeader(s int) {
	rw.status = s
	rw.callBefore()
	rw.ResponseWriter.WriteHeader(s)
}



func (rw *responseWriter) callBefore() {

}

type responseWriterCloseNotifer struct {
	*responseWriter
}

func (r responseWriterCloseNotifer) CloseNotify() <-chan bool {
	return r.ResponseWriter.(http.CloseNotifier).CloseNotify()
}
