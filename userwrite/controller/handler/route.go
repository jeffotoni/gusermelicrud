// Go Api server
// @jeffotoni
package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	mw "github.com/jeffotoni/gusermeli/apicore/middleware"
	hd "github.com/jeffotoni/gusermeli/apicore/pkg/headers"
)

func AllRoutes(app *fiber.App) {
	app.Use(mw.MaxBody(BODY_LIMIT)) //maximo para requests normais
	mw.Cors(app)
	mw.Logger(app)
	mw.Compress(app)
	mw.MsgUUID(app)

	app.Post("/v1/user/ping", limiter.New(limiter.Config{
		Next:       nil,
		Max:        LIMIT_RATE_PING,
		Expiration: 5 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), Ping)

	app.Post("/v1/user", limiter.New(limiter.Config{
		Next:       nil,
		Max:        LIMIT_RATE_POST,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), UserPost)

	app.Put("/v1/user/:id", limiter.New(limiter.Config{
		Next:       nil,
		Max:        LIMIT_RATE_PUT,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		}}), UserPut)

}
