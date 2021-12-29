package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gusermeli/apicore/pkg/fmts"
	tests "github.com/jeffotoni/gusermeli/apicore/pkg/tests"
	drepo "github.com/jeffotoni/gusermeli/userwrite/domain/repo"
)

//go test -v -run ^TestPost$
func TestPost(t *testing.T) {
	type args struct {
		method      string
		ctype       string
		header      map[string]string
		url         string
		handlerfunc func(c *fiber.Ctx) error
		user        drepo.User
	}

	tt := []struct {
		name     string
		args     args
		want     int //status code
		bodyShow bool
	}{
		{name: "test_user_post", args: args{
			method:      "POST",
			ctype:       "application/json",
			header:      map[string]string{"Authorization": fmts.ConcatStr("Bearer ", os.Getenv("token3"))},
			url:         "/v1/user",
			handlerfunc: UserPost,
			user: drepo.User{
				FirstName:       "",
				LastName:        "",
				Birthday:        "",
				Cpf:             "",
				Email:           "",
				Password:        "",
				ConfirmPassword: "",
				CreatedAt:       "",
				UpdatedAt:       "",
				IP:              "",
				Agent:           "",
			},
		}, want: 400, bodyShow: true},

		{name: "test_user_post", args: args{
			method:      "POST",
			ctype:       "application/json",
			header:      map[string]string{},
			url:         "/v1/user",
			handlerfunc: UserPost,
			user: drepo.User{
				FirstName:       "Paul",
				LastName:        "Churchill",
				Birthday:        "1920-08-20",
				Cpf:             "291.450.370-00",
				Email:           "paul@gmail.com",
				Password:        "123456",
				ConfirmPassword: "123456",
			},
		}, want: 400, bodyShow: true},

		{name: "test_user_post", args: args{
			method:      "POST",
			ctype:       "application/json",
			header:      map[string]string{},
			url:         "/v1/user",
			handlerfunc: UserPost,
			user: drepo.User{
				FirstName:       "Paul",
				LastName:        "Churchill",
				Birthday:        "1920-08-20",
				Cpf:             "291.450.370-94",
				Email:           "paul@gmail.com",
				Password:        "123456",
				ConfirmPassword: "123456",
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

			tests.TestNewRequest(t, tt1.args.url, "", tt1.args.method,
				tt1.args.handlerfunc,
				bytes.NewBuffer(requestBody),
				tt1.args.ctype, tt1.args.header, tt1.want, tt1.bodyShow)
		})
	}
}
