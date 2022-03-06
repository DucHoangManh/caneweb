package cane

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node // each root have request method info (GET, POST...)
	handlers map[string]Handler
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]Handler),
	}
}

// create route from method and path
func getRoute(method, path string) (route string) {
	var builder strings.Builder
	_, _ = fmt.Fprintf(&builder, "%s%s%s", method, "-", path)
	route = builder.String()
	return
}

func parsePattern(pattern string) []string {
	split := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, s := range split {
		if s != "" {
			parts = append(parts, s)
			if s[0] == '*' {
				break
			}
		}
	}
	return parts
}

// add a new route to the router
func (r *router) addRoute(method, path string, handler Handler) {
	route := getRoute(method, path)
	parts := parsePattern(path)
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(path, parts, 0)
	r.handlers[route] = handler
}

func (r *router) getRoute(method, path string) (node *node, vars map[string]string) {
	partsToSearch := parsePattern(path)
	vars = make(map[string]string, 0)
	root, ok := r.roots[method]
	if !ok {
		return
	}

	node = root.search(partsToSearch, 0)
	if node != nil {
		// found a path, now fill the path variables
		parts := parsePattern(node.pattern)
		for i, v := range parts {
			if v[0] == ':' {
				vars[v[1:]] = partsToSearch[i]
			}
			if v[0] == '*' {
				vars[v[1:]] = strings.Join(partsToSearch[i:], "/")
			}
		}
		return
	}
	return nil, nil
}

//route request to appropriate handler, error code if no handler found
func (r *router) handle(c *Ctx) {
	node, vars := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Vars = vars
		key := getRoute(c.Method, node.pattern)
		r.handlers[key].Serve(c)
	} else {
		c.Writer.WriteHeader(http.StatusNotFound)
		c.String(http.StatusNotFound, "No route found")
	}
}
