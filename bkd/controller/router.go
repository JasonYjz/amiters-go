package controller

import (
	"amiters-go/controller/api"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"os"
)

func Router() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/post").Handle(new(api.PostController))
	})

	app.Configure(iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler: false,
		DisablePathCorrection:   false,
		EnablePathEscape:        false,
		TimeFormat:              "2006-01-02 15:04:05",
		Charset:                 "utf-8",
	}))
	err := app.Run(iris.Addr("localhost:8080"))
	if err != nil {
		os.Exit(-1)
	}
}
