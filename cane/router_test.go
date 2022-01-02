package cane

import (
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("POST", "/test", HandleFunc(sampleHandler))
	r.addRoute("GET", "/post", HandleFunc(sampleHandler))
	return r
}

func TestRouter(t *testing.T) {
	router := newTestRouter()
	if _, ok := router.handlers["POST-/test"]; !ok {
		t.Error("err route configuration")
	}
	if _, ok := router.handlers["GET-/post"]; !ok {
		t.Error("err route configuration")
	}
	if _, ok := router.handlers["GET-/test"]; ok {
		t.Error("err route configuration")
	}

}

func sampleHandler(c *Ctx) {
	c.String(200, "%s", "test")
}
