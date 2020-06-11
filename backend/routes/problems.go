package routes

import (
	"backend/views"
	"github.com/kataras/iris/v12"
)

func RegisterProblemsRouter (app *iris.Application) {
	problems := app.Party("/problems")
	problems.Get("/list", views.IndexView)		// for test
}