package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//Logger log
func Logger(app *fiber.App) {
	//app.Use(mw.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n"))
	if os.Getenv("ENV_AMBI") != "PROD" {
		app.Use(logger.New(logger.Config{
			Format:     "${pid} ${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			Output:     os.Stdout,
		}))
	}
	return
}
