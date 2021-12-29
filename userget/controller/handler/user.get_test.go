package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	tests "github.com/jeffotoni/gusermeli/apicore/pkg/tests"
)

//go test -v -run ^TestGet$
func TestGet(t *testing.T) {
	type args struct {
		method      string
		ctype       string
		header      map[string]string
		url         string
		urlid       string
		handlerfunc func(c *fiber.Ctx) error
	}

	tt := []struct {
		name     string
		args     args
		want     int //status code
		bodyShow bool
	}{

		{name: "test_user_get", args: args{
			method:      "GET",
			ctype:       "application/json",
			header:      map[string]string{},
			url:         "/v1/user",
			urlid:       `/v1/user?names=\[Paul\]`,
			handlerfunc: UserGet,
		}, want: 200, bodyShow: true},

		{name: "test_user_get", args: args{
			method:      "GET",
			ctype:       "application/json",
			header:      map[string]string{},
			url:         "/v1/user",
			urlid:       `/v1/user?names=\[Paul2\]`,
			handlerfunc: UserGet,
		}, want: 204, bodyShow: true},
	}

	for _, tt := range tt {

		tt1 := tt
		t.Run(tt1.name, func(t *testing.T) {
			tests.TestNewRequest(t, tt1.args.url, tt1.args.urlid, tt1.args.method,
				tt1.args.handlerfunc,
				//bytes.NewBuffer(requestBody),
				nil,
				tt1.args.ctype, tt1.args.header, tt1.want, tt1.bodyShow)
		})
	}
}
