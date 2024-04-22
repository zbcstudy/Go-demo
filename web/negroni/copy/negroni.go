package copy

import "net/http"

const (
	defaultAddress = ":8080"
)

type Handler interface {
	ServerHttp(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc) ServerHttp(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(rw, r, next)
}

type middleware struct {
	handler Handler
	next    *middleware
}

func (m middleware) serverHttp(rw http.ResponseWriter, r *http.Request) {
	m.handler.ServerHttp(rw, r, m.next.serverHttp)
}

func Wrap(handler http.Handler) Handler {
	return HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(rw, r)
		next(rw, r)
	})
}

func WrapFunc(handlerFunc http.HandlerFunc) Handler {
	return HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		handlerFunc(rw, r)
		next(rw, r)
	})
}

type Negroni struct {
	middleware middleware
	handlers   []Handler
}

func New(handlers ...Handler) *Negroni {
	return &Negroni{
		handlers:   handlers,
		middleware: build(handlers),
	}
}

func (n *Negroni) With(handlers ...Handler) *Negroni {
	return New(
		append(n.handlers, handlers...)...,
	)
}

func Classic() *Negroni {
	return New()
}

func build(handlers []Handler) middleware {
	var next middleware

	if len(handlers) == 0 {
		return voidMiddleware()
	} else if len(handlers) > 1 {
		next = build(handlers[1:])
	} else {
		next = voidMiddleware()
	}

	return middleware{
		handler: handlers[0],
		next:    &next,
	}
}

func voidMiddleware() middleware {
	return middleware{
		HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {}),
		&middleware{},
	}
}
