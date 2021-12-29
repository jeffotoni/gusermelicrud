package handler

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	mw "github.com/jeffotoni/gusermeli/apicore/middleware"
	"github.com/jeffotoni/gusermeli/apicore/pkg/crypt"
	gemail "github.com/jeffotoni/gusermeli/apicore/pkg/email/gmail"
	"github.com/jeffotoni/gusermeli/apicore/pkg/env"
	"github.com/jeffotoni/gusermeli/apicore/pkg/fmts"
	hd "github.com/jeffotoni/gusermeli/apicore/pkg/headers"
	valid "github.com/jeffotoni/gusermeli/apicore/pkg/validador"
	zlog "github.com/jeffotoni/gusermeli/apicore/pkg/zerolog"
	drepo "github.com/jeffotoni/gusermeli/userwrite/domain/repo"
	"github.com/rs/zerolog/log"
)

var (
	TIMEOUT_POST = env.GetDuration("TIMEOUT_POST", 60*time.Second)
)

//UserPost
func UserPost(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	code := 400
	msgID := mw.GetUUID(c)

	ctx := context.WithValue(context.Background(), "msgID", msgID)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_POST)
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
			Str("handler", "(s StructConnect) UserPost(c)").
			Msg("Body is required")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"Body is required"}`))
	}

	var user drepo.User
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

	if len(user.Password) <= 5 {
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
			Str("handler", "user.Password").
			Msg("invalid Password")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"invalid Password"}`))
	}

	if user.Password != user.ConfirmPassword {
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
			Str("handler", "user.Password != user.ConfirmPassword").
			Msg("invalid ConfirmPassword")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"invalid Confirm Password"}`))
	}

	user.ID = crypt.GSha1(utils.UUID())
	user.CreatedAt = time.Now().Format("2006-01-02T15:04:05.000Z")
	user.UpdatedAt = user.CreatedAt
	user.IP = hd.IP(c)
	user.Agent = hd.UserAgent(c)
	user.Cpf = strings.Replace(user.Cpf, ".", "", -1)
	user.Cpf = strings.Replace(user.Cpf, "-", "", -1)
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

	if !valid.RegexEmail(user.Email) {
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
			Str("handler", "user.RegexEmail").
			Msg("invalid RegexEmail")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"invalid Email"}`))
	}

	if len(user.FirstName) == 0 {
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
			Str("handler", "user.FirstName").
			Msg("invalid FirstName")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"invalid FirstName"}`))
	}

	if len(user.LastName) == 0 {
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
			Str("handler", "user.LastName").
			Msg("invalid LastName")
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"invalid LastName"}`))
	}

	err = user.Create(ctx)
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
			Str("handler", "user.Create").
			Msg(err.Error())
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"`, err.Error(), `"}`))
	}

	tokenx := crypt.GSha1(utils.UUID())
	go func(tokenx, email string) {
		link_pass := fmts.ConcatStr("https://domain.com", "/setemail/index.html?t=", tokenx)
		gemail.SendUser([]string{email}, "Successfully created", email, link_pass, "https://domain.com")
	}(tokenx, user.Email)

	return c.Status(200).SendString(fmts.ConcatStr(`{"id":"`, user.ID, `","email":"`, user.Email, `","cpf":"`, user.Cpf, `"}`))
}
