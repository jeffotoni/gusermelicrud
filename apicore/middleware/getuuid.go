package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/utils"
)

const MsgID = "msgID"

func MsgUUID(app *fiber.App) {
	app.Use(requestid.New(requestid.Config{
		Header: MsgID,
	}))
	return
}

func GetUUID(c *fiber.Ctx) string {
	uuid := string(c.Request().Header.Peek(MsgID))
	if len(uuid) == 0 {
		uuid = utils.UUID()
		c.Request().Header.Set(MsgID, uuid)
	}
	return uuid
}
