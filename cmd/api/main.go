package main

import (
	"fmt"
	app2 "github.com/dElCIoGio/filestorage/internal/app"
	"github.com/dElCIoGio/filestorage/internal/platform/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {

	cfg := config.GetConfig()

	var app *app2.FileStorageApp

	switch cfg.Env {
	case config.PROD:
		app, _ = app2.NewProd(cfg)
		break
	case config.TEST:
		app, _ = app2.NewTest(cfg)
		break
	default:
	case config.DEV:
		app, _ = app2.NewDev(cfg)
		break
	}
	if app == nil {
		return
	}

	app.Logger.Info(fmt.Sprintf("Starting server on port %s", app.Config.Port))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	err := http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		return
	}
}
