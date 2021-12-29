package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gofiber/fiber/v2"
	tests "github.com/jeffotoni/gusermeli/apicore/pkg/tests"
	drepo "github.com/jeffotoni/gusermeli/userwrite/domain/repo"
)

//go test -v -run ^TestPut$
func TestPut(t *testing.T) {
	type args struct {
		method      string
		ctype       string
		header      map[string]string
		url         string
		urlid       string
		handlerfunc func(c *fiber.Ctx) error
		user        drepo.User
	}

	tt := []struct {
		name     string
		args     args
		want     int //status code
		bodyShow bool
	}{
		// {name: "test_user_put", args: args{
		// 	method:      "PUT",
		// 	ctype:       "application/json",
		// 	header:      map[string]string{"Authorization": fmts.ConcatStr("Bearer ", os.Getenv("token3"))},
		// 	url:         "/v1/user/29145037094",
		// 	handlerfunc: UserPut,
		// 	user: drepo.User{
		// 		FirstName: "",
		// 		LastName:  "",
		// 		Birthday:  "",
		// 	},
		// }, want: 400, bodyShow: true},

		// {name: "test_user_put", args: args{
		// 	method:      "PUT",
		// 	ctype:       "application/json",
		// 	header:      map[string]string{},
		// 	url:         "/v1/user/29145037000",
		// 	handlerfunc: UserPut,
		// 	user: drepo.User{
		// 		FirstName: "Paul",
		// 		LastName:  "Churchill",
		// 		Birthday:  "1920-08-20",
		// 	},
		// }, want: 400, bodyShow: true},

		{name: "test_user_put", args: args{
			method:      "PUT",
			ctype:       "application/json",
			header:      map[string]string{},
			url:         "/v1/user/:id",
			urlid:       "/v1/user/29145037094",
			handlerfunc: UserPut,
			user: drepo.User{
				FirstName: "Paul2",
				LastName:  "Churchill2",
				Birthday:  "1925-08-20",
			},
		}, want: 200, bodyShow: true},
	}

	for _, tt := range tt {
		requestBody, err := json.Marshal(&tt.args.user)
		if err != nil {
			t.Errorf("Error json.Marshal:%s", err.Error())
			return
		}
		tt1 := tt
		t.Run(tt1.name, func(t *testing.T) {
			//t.Parallel()
			if tt1.bodyShow {
				fmt.Println("Json: ", string(requestBody))
			}

			tests.TestNewRequest(t, tt1.args.url, tt1.args.urlid, tt1.args.method,
				tt1.args.handlerfunc,
				bytes.NewBuffer(requestBody),
				tt1.args.ctype, tt1.args.header, tt1.want, tt1.bodyShow)
		})
	}
}
