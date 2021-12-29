package headers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RemoteAddr(r *http.Request) string {
	if len(r.Header.Get("X-Real-IP")) > 0 {
		return r.Header.Get("X-Real-IP")
	}
	if len(r.Header.Get("X-Forwarded-For")) > 0 {
		return r.Header.Get("X-Forwarded-For")
	}
	return "127.0.0.1"
}

func IP(c *fiber.Ctx) string {
	ipReal := string(c.Request().Header.Peek("X-Real-IP"))
	if len(ipReal) <= 0 {
		return "127.0.0.1"
	}
	return ipReal
}

func UserAgent(c *fiber.Ctx) string {
	agent := string(c.Request().Header.Peek("User-Agent"))
	if len(agent) <= 0 {
		return "no agent"
	}
	return agent
}

func MsgID(c *fiber.Ctx) string {
	msgid := string(c.Request().Header.Peek("Msgid"))
	return msgid
}

func Host(c *fiber.Ctx) string {
	host := string(c.Request().Header.Peek("Host"))
	return host
}

func ContentLength(c *fiber.Ctx) int {
	contentlength := string(c.Request().Header.Peek("Content-Length"))
	cl, _ := strconv.Atoi(contentlength)
	return cl
}

func ContentType(c *fiber.Ctx) string {
	contentype := string(c.Request().Header.Peek("Content-Type"))
	return contentype
}
