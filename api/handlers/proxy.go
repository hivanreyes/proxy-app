package handlers

import (
	iris "github.com/kataras/iris"
)

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application){
	app.Get("/ping", func(c iris.Context){
		c.JSON(iris.Map{"result": "ok", "result2": "ok2"})
	})
}