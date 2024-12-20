package apiHandlers

import (
	"RestaurentMGT/api"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Router(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	group := app.Group("/RestaurentMGT/api")
	defaultGroup := app.Group("/")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Static("/", "./docs/rapiDoc/build")
	DefaultMappings(defaultGroup)
	RouteMappings(group)
}

func RouteMappings(cg fiber.Router) {
cg.Post("/CreateRestaurent", api.CreateRestaurentApi)
cg.Put("/UpdateRestaurent", api.UpdateRestaurentApi)
cg.Delete("/DeleteRestaurent", api.DeleteRestaurentApi)
cg.Get("/FindRestaurent", api.FindRestaurentApi)
cg.Get("/FindallRestaurent", api.FindallRestaurentApi)
cg.Post("/UploadRestaurent", api.UploadRestaurentApi)
cg.Get("/DownloadRestaurent", api.DownloadRestaurentApi)
cg.Post("/CreateCustomer", api.CreateCustomerApi)
cg.Put("/UpdateCustomer", api.UpdateCustomerApi)
cg.Delete("/DeleteCustomer", api.DeleteCustomerApi)
cg.Get("/FindCustomer", api.FindCustomerApi)
cg.Get("/FindallCustomer", api.FindallCustomerApi)
cg.Post("/UploadCustomer", api.UploadCustomerApi)
cg.Get("/DownloadCustomer", api.DownloadCustomerApi)
cg.Get("/FindallifCustomerByRestaurentId/RestaurentId", api.FindallifCustomerByRestaurentIdApi)

cg.Get("/swagger", api.SwaggerHandler)

}

func DefaultMappings(cg fiber.Router) {
	cg.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "RestaurentMGT-APP1706 service is up and running", "version": "1.0"})
	})
}