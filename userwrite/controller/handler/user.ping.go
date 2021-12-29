package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	mw "github.com/jeffotoni/gusermeli/apicore/middleware"
	hd "github.com/jeffotoni/gusermeli/apicore/pkg/headers"
	mgoConn "github.com/jeffotoni/gusermeli/apicore/pkg/mongo"
	zlog "github.com/jeffotoni/gusermeli/apicore/pkg/zerolog"
	"github.com/rs/zerolog/log"
)

//Ping pong
func Ping(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	code := 400
	msgID := mw.GetUUID(c)
	err := mgoConn.Ping()
	if err != nil {
		log.Info().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "api.user").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "(s Conn) Ping(c)").
			Str("func", "Ping(c)").Msg("")
		return c.Status(code).SendString("")
	}
	code = 200
	return c.Status(code).SendString(`{"pong":"🏓"}`)
}
