package main
import (
	handlers "wizeline.github.com/hivanreyes/proxy-app/api/handlers"
	utils "wizeline.github.com/hivanreyes/proxy-app/api/utils"
	server "wizeline.github.com/hivanreyes/proxy-app/api/server"
)


/*
	Router iris
	Env Vars
*/

func main(){
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}