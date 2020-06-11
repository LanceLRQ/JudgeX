package core

import (
	"backend/routes"
	"backend/views"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

var Server JudgeXServer

type JudgeXServer struct {
	App *iris.Application
}

func InitJudgeXServer () error {
	Server := &JudgeXServer{}
	Server.App = iris.New()
	Server.App.Use(recover.New())
	Server.App.Use(logger.New())

	if Settings.DebugMode {
		Server.App.Logger().SetLevel("debug")
	} else {
		Server.App.Logger().SetLevel("info")
	}

	Server.InitRoutes()

	err := Server.App.Run(iris.Addr(fmt.Sprintf("%s:%d", Settings.Server.Listen, Settings.Server.Port)))
	return err
}


func (server *JudgeXServer) InitRoutes() {
	server.App.Get("/", views.IndexView)
	routes.RegisterProblemsRouter(server.App)
}