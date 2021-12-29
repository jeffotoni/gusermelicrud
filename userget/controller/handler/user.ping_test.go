package handler

import (
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	mgoConn "github.com/jeffotoni/gusermeli/apicore/pkg/mongo"
	tests "github.com/jeffotoni/gusermeli/apicore/pkg/tests"
)

//go test -v -run ^TestPing$
func TestPing(t *testing.T) {
	ctx, err := mgoConn.Connect()
	if err != nil {
		log.Println("error connect MongoDb")
		return
	}
	defer mgoConn.Disconnect(ctx)

	tests.TestNewRequest(t, "/v1/user/ping", "", "GET", Ping, nil, "application/json", map[string]string{}, 200, true)
}

//go test -v -run ^TestPing2$
func TestPing2(t *testing.T) {
	ctx, err := mgoConn.Connect()
	if err != nil {
		log.Println("error connect MongoDb")
		return
	}
	defer mgoConn.Disconnect(ctx)

	bodyResponse := `{"pong":"🏓"}`
	app := fiber.New()
	app.Get("/v1/user/ping", Ping)

	req := httptest.NewRequest("GET", "/v1/user/ping", nil)
	//req.Header.Add("Authorization", "Bearer $token")
	resp, err := app.Test(req)

	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")

	body, err := io.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, bodyResponse, string(body))
}
