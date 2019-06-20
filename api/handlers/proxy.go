package handlers

import (
	iris "github.com/kataras/iris"
	"wizeline.github.com/hivanreyes/proxy-app/api/middleware"
)

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application){
	app.Get("/ping", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context){
	c.JSON(iris.Map{"result": "ok"})
}