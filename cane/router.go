package cane

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]Handler
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]Handler),
	}
}

// create route from method and path
func getRoute(method, path string) (route string) {
	var builder strings.Builder
	fmt.Fprintf(&builder, "%s%s%s", method, "-", path)
	route = builder.String()
	return
}

// add a new route to the router
func (r *router) addRoute(method, path string, handler Handler) {
	log.Printf("add route %s %s", method, path)
	route := getRoute(method, path)
	r.handlers[route] = handler
}

//route request to appropriate handler, error code if no handler found
func (r *router) handle(c *Ctx) {
	route := getRoute(c.Method, c.Path)
	if handler, ok := r.handlers[route]; ok {
		handler.Serve(c)
	} else {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.String(http.StatusNotFound, "No route found")
	}
}
