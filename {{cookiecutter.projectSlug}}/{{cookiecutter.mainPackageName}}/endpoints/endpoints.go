package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"{{cookiecutter.__mainPackagePath}}/config"
	"{{cookiecutter.__mainPackagePath}}/db"
	"{{cookiecutter.__mainPackagePath}}/paths"
	"{{cookiecutter.__mainPackagePath}}/util"
)

type Endpoints struct {
	db *db.DB
}

func New(dbi *db.DB) *Endpoints {
	return &Endpoints{
		db: dbi,
	}
}

func (e *Endpoints) SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler:          util.JSONErrorHandler,
		DisableStartupMessage: !config.Debug.Enabled,

		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Get(paths.Index, e.Index)

	return app
}
