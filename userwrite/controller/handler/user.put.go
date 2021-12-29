package handler

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	mw "github.com/jeffotoni/gusermeli/apicore/middleware"
	"github.com/jeffotoni/gusermeli/apicore/pkg/env"
	"github.com/jeffotoni/gusermeli/apicore/pkg/fmts"
	hd "github.com/jeffotoni/gusermeli/apicore/pkg/headers"
	valid "github.com/jeffotoni/gusermeli/apicore/pkg/validador"
	zlog "github.com/jeffotoni/gusermeli/apicore/pkg/zerolog"
	drepo "github.com/jeffotoni/gusermeli/userwrite/domain/repo"
	"github.com/rs/zerolog/log"
)

var (
	TIMEOUT_PUT = env.GetDuration("TIMEOUT_PUT", 60*time.Second)
)

//UserPut
func UserPut(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	code := 400
	msgID := mw.GetUUID(c)

	ctx := context.WithValue(context.Background(), "msgID", msgID)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_PUT)
	defer cancel()

	if len(string(c.Body())) <= 0 {
		log.Error().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "userwrite").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "(s StructConnect) UserPut(c)").
			Msg("Body is required")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"Body is required"}`))
	}

	var user drepo.UserUp
	err := c.BodyParser(&user)
	if err != nil {
		log.Error().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "userwrite").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "BodyParser(&user)").
			Msg(err.Error())
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"`, err.Error(), `"}`))
	}

	user.UpdatedAt = time.Now().Format("2006-01-02T15:04:05.000Z")
	user.IP = hd.IP(c)
	user.Agent = hd.UserAgent(c)
	user.Cpf = c.Params("id")
	user.Cpf = strings.Replace(user.Cpf, "-", "", -1)
	user.Cpf = strings.Replace(user.Cpf, ".", "", -1)
	if !valid.IsCPF(user.Cpf) {
		log.Error().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "userwrite").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "user.IsCPF").
			Msg("invalid CPF")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"invalid CPF"}`))
	}

	if !valid.Birthday(user.Birthday) { // birdthday < 18 => error
		log.Error().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "userwrite").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "user.Birthday").
			Msg("invalid Birthday")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"invalid Birthday"}`))
	}

	err = user.Update(ctx)
	if err != nil {
		log.Error().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "userwrite").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "user.Update").
			Msg(err.Error())
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"`, err.Error(), `"}`))
	}

	return c.Status(200).SendString("")
}
