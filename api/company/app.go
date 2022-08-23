package company

import (
	"fmt"
	"hr/api/company/controller"
	"hr/config"
	"hr/util"

	// "hr/util"

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
	a.app.Use(util.NewJwt(util.Jwt{
		Exception: []string{"/login"},
		Secret:    a.config.Get("jwt.secret").ToString(),
	}).Serve)
	a.loadRoute()
	addr := fmt.Sprintf("%s:%s", a.config.Get("addr.host").ToString(), a.config.Get("addr.port").ToString())
	a.app.Run(iris.Addr(addr), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func (a *App) loadRoute() {
	loginController := controller.NewLoginController()
	a.app.Post("/login", loginController.Login)
	a.app.Get("/logout", loginController.Logout)
	demandController := controller.NewDemandController()
	a.app.PartyFunc("/demand", func(p router.Party) {
		p.Get("/", demandController.GetList)
		p.Post("/edit", demandController.Edit)
	})
	staffController := controller.NewStaffController()
	a.app.PartyFunc("/staff", func(p router.Party) {
		p.Get("/", staffController.GetList)
		p.Get("/{id:int}", staffController.GetInfo)
	})
	salaryController := controller.NewSalaryController()
	a.app.PartyFunc("/salary", func(p router.Party) {
		p.Get("/", salaryController.GetList)
		p.Get("/{id:int}", salaryController.GetInfo)
	})
}
