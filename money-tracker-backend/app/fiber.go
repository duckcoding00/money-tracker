package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Application struct {
	config Config
}

type Config struct {
	portAddress string
	dbConfig    DBConfig
}

type DBConfig struct {
	dbAddr      string
	maxOpenCons int
	maxIdleCons int
	maxIdleTime string
}

func (app *Application) Config() *fiber.App {
	r := fiber.New()

	// setup cors

	v1 := r.Group("/v1")
	api := v1.Group("/api")

	api.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "test oke",
		})
	})

	// setup user routes
	app.userRoutes(api)

	return r
}

// users API
func (app *Application) userRoutes(v1 fiber.Router) {
	users := v1.Group("/users")

	users.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("oke")
	})
}

func (app *Application) Run(r *fiber.App) error {
	log.Infof("server running at port: %v", app.config.portAddress)
	return r.Listen(fmt.Sprintf(":%s", app.config.portAddress))
}
