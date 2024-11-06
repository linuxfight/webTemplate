package setup

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"webTemplate/cmd/app"
	v1 "webTemplate/internal/adapters/controller/api/v1"
)

func Setup(app *app.App) {
	app.Fiber.Use(cors.New(cors.ConfigDefault))

	app.Fiber.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1",
		FilePath: "./docs/swagger.json",
		Path:     "./docs",
		Title:    "Swagger API Docs",
	}))

	if viper.GetBool("settings.debug") {
		app.Fiber.Use(logger.New(logger.Config{TimeZone: viper.GetString("settings.timezone")}))
	}

	// Setup api v1 routes
	apiV1 := app.Fiber.Group("/api/v1")

	//middlewareHandler := middlewares.NewMiddlewareHandler(app)
	//
	// Setup user routes
	userHandler := v1.NewUserHandler(app)
	userHandler.Setup(apiV1)

	//
	//// Setup question routes
	//questionHandler := v1.NewQuestionHandler(app)
	//questionHandler.Setup(apiV1, middlewareHandler.IsAuthenticated)
	//
	//// Setup conference routes
	//conferenceHandler := v1.NewConferenceHandler(app)
	//conferenceHandler.Setup(apiV1, middlewareHandler.IsAuthenticated)
	//
	//eventHandler := v1.NewEventHandler(app)
	//eventHandler.Setup(apiV1, middlewareHandler.IsAuthenticated)
}