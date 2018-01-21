package controllers

import (
	"github.com/kataras/iris/mvc"
)

type MainController struct{}

// // Get serves
// // Method:   GET
// // Resource: http://localhost:8080
func (c *MainController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Main Controller 2</h1>",
	}
}
