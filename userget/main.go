// Go Api server
// @jeffotoni
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	mgoConn "github.com/jeffotoni/gusermeli/apicore/pkg/mongo"
	"github.com/jeffotoni/gusermeli/userget/config"
	route "github.com/jeffotoni/gusermeli/userget/controller/handler"
)

func main() {
	ctx, err := mgoConn.Connect()
	if err != nil {
		log.Println("error connect MongoDb")
		return
	}
	defer mgoConn.Disconnect(ctx)

	app := fiber.New(fiber.Config{
		Concurrency: config.FIBER_CONCURRENCY,
		BodyLimit:   config.FIBER_BODY_LIMIT,
	})
	route.AllRoutes(app)
	app.Listen(config.HTTPPORT)
}
