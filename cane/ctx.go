package cane

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Map map[string]interface{}

type Ctx struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request data
	Path   string
	Method string
	// response data
	StatusCode int
}

// newCtx create new Ctx with original data
func newCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// responseWriter
func (c *Ctx) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Ctx) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Ctx) String(code int, formatString string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, err := fmt.Fprintf(c.Writer, formatString, values...)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Ctx) JSON(code int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Ctx) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

// request
func (c *Ctx) FormValue(key string) string {
	return c.Req.FormValue(key)
}

func (c *Ctx) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
