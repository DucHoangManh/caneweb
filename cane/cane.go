package cane

import "net/http"

type Handler interface {
	Serve(c *Ctx)
}

type HandleFunc func(c *Ctx)

func (f HandleFunc) Serve(c *Ctx) {
	f(c)
}

type Engine struct {
	router *router
}

// constructor of cane web framework
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (e *Engine) addRoute(method, pattern string, handler Handler) {
	e.router.addRoute(method, pattern, handler)
}

// define some simple operations
func (e *Engine) GET(pattern string, handler Handler) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler Handler) {
	e.addRoute(http.MethodPost, pattern, handler)
}

// implement the standard package Handler interface and transform incoming request to our handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newCtx(w, r)
	e.router.handle(ctx)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
