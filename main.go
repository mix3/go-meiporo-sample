package main

import (
	"github.com/lestrrat/go-xslate"
	"github.com/mix3/go-meiporo-sample/model"
	"github.com/mix3/meiporo"
)

func main() {
	m := meiporo.New()

	// middleware
	m.Middlewares = []meiporo.Handler{}
	m.Use(&meiporo.LoggerMiddleware{})
	m.Use(&meiporo.RecoverMiddleware{
		RecoverHandler: ErrorHandler,
	})
	m.Use(&meiporo.StaticFileMiddleware{PublicPath: "public"})

	// routing
	m.Router.Add("GET", "/", RootHandler)
	m.Router.Add("GET", "/list", ListHandler)
	m.Router.Add("POST", "/create", CreateHandler)
	m.Router.Add("POST", "/switch/:id", SwitchHandler)
	m.Router.Add("POST", "/delete/:id", DeleteHandler)
	m.Router.Add("POST", "/delete", DeleteAllHandler)

	// xslate
	meiporo.Xslate.Configure(xslate.Args{
		"Loader": xslate.Args{
			"LoadPaths": []string{"tmpl"},
		},
	})

	db := model.GetDB()
	defer db.Close()

	m.Run()
}
