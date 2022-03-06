package main

import (
	"caneweb/cane"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := cane.New()
	server.GET("/hello", cane.HandleFunc(hello))
	server.POST("/post", cane.HandleFunc(createPost))
	server.GET("/hi/:name/say", cane.HandleFunc(sayHi))
	fmt.Println("Server start at localhost:5445...")
	log.Fatal(server.Run(":5445"))
}

func hello(c *cane.Ctx) {
	c.String(http.StatusOK, "hello %s", c.Query("name"))
}

func createPost(c *cane.Ctx) {
	title := c.FormValue("title")
	desc := c.FormValue("desc")
	c.JSON(http.StatusOK, cane.Map{
		"post_title":  title,
		"description": desc,
	})
}

func sayHi(c *cane.Ctx) {
	c.String(http.StatusOK, "Hi there, %s", c.Var("name"))
}
