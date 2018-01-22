package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type MainController struct {
	// Optionally: context is auto-binded by Iris on each request,
	// remember that on each incoming request iris creates a new UserController each time,
	// so all fields are request-scoped by-default, only dependency injection is able to set
	// custom fields like the Service which is the same for all requests (static binding).
	Ctx iris.Context
}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *MainController) Get() mvc.Result {
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{},
	}
}
