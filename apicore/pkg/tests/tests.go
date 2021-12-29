package tests

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	utils "github.com/gofiber/utils"
	mgoConn "github.com/jeffotoni/gusermeli/apicore/pkg/mongo"
	"github.com/stretchr/testify/assert"
)

func TestNewRequest(t *testing.T, url string, urlid string,
	method string,
	handlerFunc func(c *fiber.Ctx) error,
	data *bytes.Buffer, contentType string, header map[string]string, want int, bodyShow bool) {

	ctx, err := mgoConn.Connect()
	if err != nil {
		log.Println("error connect MongoDb")
		return
	}
	defer mgoConn.Disconnect(ctx)

	app := fiber.New()
	if method == "POST" {
		app.Post(url, handlerFunc)
	} else if method == "PUT" {
		app.Put(url, handlerFunc)
	} else if method == "GET" {
		app.Get(url, handlerFunc)
	} else if method == "DELETE" {
		app.Delete(url, handlerFunc)
	} else if method == "PATCH" {
		app.Patch(url, handlerFunc)
	}

	if len(urlid) > 0 {
		url = urlid
	}
	var req *http.Request
	if data == nil {
		req = httptest.NewRequest(method, url, nil)
	} else {
		req = httptest.NewRequest(method, url, data)
	}

	if len(contentType) > 0 {
		req.Header.Set("Content-Type", contentType)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := app.Test(req)
	if err != nil {
		if want == 400 {
			t.Logf("app.Test:%s", err.Error())
			return
		}
		t.Errorf("Error app.Test:%s", err.Error())
		return
	}

	utils.AssertEqual(t, nil, err, "app.Test(req)")
	assert.Equal(t, want, resp.StatusCode)

	if bodyShow {
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error io.ReadAll:%s", err.Error())
			return
		}

		t.Log("status:", resp.StatusCode)
		t.Log("\nResp :\n", string(body))
	}
}
