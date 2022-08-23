package admin

import (
	"fmt"
	"hr/api/admin/controller"
	"hr/config"

	uc "github.com/hechen0210/utils/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/recover"
)

type App struct {
	config *uc.ConfigData
	app    *iris.Application
}

func NewApp() (*App, error) {
	err := config.InitStorage()
	if err != nil {
		return nil, err
	}
	return &App{
		config: config.GetConfig(),
		app:    iris.New(),
	}, nil
}

func (a *App) Run() {
	a.app.Logger().SetLevel("debug")
	a.app.UseRouter(recover.New())
	a.loadRoute()
	addr := fmt.Sprintf("%s:%s", a.config.Get("addr.host").ToString(), a.config.Get("addr.port").ToString())
	a.app.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func (a *App) loadRoute() {
	login := controller.NewLoginController()
	a.app.Post("/login", login.Login)
	a.app.Get("/logout", login.Logout)
	demand := controller.NewDemandController()
	a.app.PartyFunc("/demand", func(party router.Party) {
		party.Get("/list", demand.GetList)
		party.Get("/info/:id", demand.GetInfo)
	})
	company := controller.NewCompanyController()
	a.app.PartyFunc("/company", func(party router.Party) {
		party.Get("/list", company.GetList)
		party.Get("/info/:id", company.GetInfo)
		party.Post("/update", company.Update)
		party.Post("/delete/:id", company.Delete)
	})
}
