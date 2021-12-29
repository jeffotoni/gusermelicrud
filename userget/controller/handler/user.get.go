package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	mw "github.com/jeffotoni/gusermeli/apicore/middleware"
	"github.com/jeffotoni/gusermeli/apicore/pkg/env"
	"github.com/jeffotoni/gusermeli/apicore/pkg/fmts"
	hd "github.com/jeffotoni/gusermeli/apicore/pkg/headers"
	redisu "github.com/jeffotoni/gusermeli/apicore/pkg/redis/user"
	zlog "github.com/jeffotoni/gusermeli/apicore/pkg/zerolog"
	drepo "github.com/jeffotoni/gusermeli/userget/domain/repo"
	"github.com/rs/zerolog/log"
)

var (
	TIMEOUT_GET = env.GetDuration("TIMEOUT_GET", 60*time.Second)
)

//UserGet
func UserGet(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	code := 400
	msgID := mw.GetUUID(c)

	ctx := context.WithValue(context.Background(), "msgID", msgID)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_GET)
	defer cancel()

	var generic []string

	var exist bool
	m := strings.Replace(c.Query("firstname"), "[", "", -1)
	m = strings.Replace(m, "]", "", -1)
	if strings.Index(m, ",") > 0 {
		generic = strings.Split(m, ",")
		exist = true
	} else if len(m) > 0 {
		generic = append(generic, m)
		exist = true
	}

	if !exist {
		m := strings.Replace(c.Query("cpf"), "[", "", -1)
		m = strings.Replace(m, "]", "", -1)
		m = strings.Replace(m, ".", "", -1)
		m = strings.Replace(m, "-", "", -1)
		//var cpfs []string
		if strings.Index(m, ",") > 0 {
			generic = strings.Split(m, ",")
			exist = true
		} else if len(m) > 0 {
			generic = append(generic, m)
			exist = true
		}
	}

	if !exist {
		m := strings.Replace(c.Query("lastname"), "[", "", -1)
		m = strings.Replace(m, "]", "", -1)
		if strings.Index(m, ",") > 0 {
			generic = strings.Split(m, ",")
			exist = true
		} else if len(m) > 0 {
			generic = append(generic, m)
			exist = true
		}
	}

	if !exist {
		m := strings.Replace(c.Query("email"), "[", "", -1)
		m = strings.Replace(m, "]", "", -1)
		if strings.Index(m, ",") > 0 {
			generic = strings.Split(m, ",")
			exist = true
		} else if len(m) > 0 {
			generic = append(generic, m)
			exist = true
		}
	}

	if !exist {
		log.Error().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "userget").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "user.UpGet").
			Msg("Parameter invalid")
		return c.Status(code).SendString(`{"msg":""Parameter invalid"}`)
	}

	fmt.Println("generic:::", generic)

	//cache redis
	KEY := strings.Join(generic, " ")
	ruJson, err := redisu.Get(KEY)
	if len(ruJson) > 0 && len(KEY) > 0 {
		fmt.Println("get the cache")
		return c.Status(200).SendString(ruJson)
	}

	userJson, err := drepo.Get(ctx, generic)
	if err != nil {
		log.Error().
			Str("@timestamp", time.Now().Format("2006-01-02T15:04:05.000Z")).
			Str("data", time.Now().Format("2006-01-02 15:04:05")).
			Str("msgid", msgID).
			Str("version", zlog.LOG_VERSION).
			Str("service", "userget").
			Str("method", c.Method()).
			Str("url", c.OriginalURL()).
			Int("content_length", hd.ContentLength(c)).
			Int("status", code).
			Str("host", hd.Host(c)).
			Str("remote_ip", hd.IP(c)).
			Str("agent", hd.UserAgent(c)).
			Str("region", zlog.LOG_REGION).
			Str("handler", "user.UserGet").
			Str("func", "drepo.Get").
			Msg(err.Error())
		return c.Status(code).SendString(fmts.ConcatStr(`{"msg":"`, err.Error(), `"}`))
	}

	if len(ruJson) == 0 && len(userJson) > 0 && len(KEY) > 0 {
		redisu.PutTTL(KEY, userJson, time.Duration(CACHE_REDIS)*time.Second)
	}

	code = 200
	if len(userJson) == 0 {
		code = 204
	}
	return c.Status(code).SendString(userJson)
}
