package handlers

import (
	"encoding/json"
	iris "github.com/kataras/iris"
	"wizeline.github.com/hivanreyes/proxy-app/api/middleware"
)

// HandlerRedirection should redirect traffic
func HandlerRedirection(app *iris.Application){
	app.Get("/ping", middleware.ProxyMiddleware, proxyHandler)
}

func proxyHandler(c iris.Context){
	response, err := json.Marshal(middleware.Que)
	if err != nil {
		c.JSON(iris.Map{"status": 400, "result": "parse error"})
		return
	}

	c.JSON(iris.Map{"result": string(response) })
}