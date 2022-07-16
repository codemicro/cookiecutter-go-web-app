package endpoints

import (
	"github.com/codemicro/go-fiber-sql/application/db"
	"github.com/codemicro/go-fiber-sql/application/urls"
	"github.com/codemicro/go-fiber-sql/application/util"
	"github.com/gofiber/fiber/v2"
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
		ErrorHandler: util.JSONErrorHandler,
	})

	app.Get(urls.Index, e.Index)

	return app
}
