package drepo

import (
	"context"
	"errors"
	"fmt"

	crypt "github.com/jeffotoni/gusermeli/apicore/pkg/crypt"
	"github.com/jeffotoni/gusermeli/apicore/pkg/fmts"
	mgoConn "github.com/jeffotoni/gusermeli/apicore/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID              string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName       string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Birthday        string `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Cpf             string `json:"cpf,omitempty" bson:"cpf,omitempty"`
	Email           string `json:"email,omitempty" bson:"email,omitempty"`
	Password        string `json:"password,omitempty" bson:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty" bson:"confirm_password,omitempty"`

	CreatedAt string `json:"create,omitempty" bson:"create,omitempty"`
	UpdatedAt string `json:"update,omitempty" bson:"update,omitempty"`
	IP        string `json:"ip,omitempty" bson:"ip,omitempty"`
	Agent     string `json:"agent,omitempty" bson:"agent,omitempty"`
}

func (u *User) Create(ctx context.Context) (err error) {
	collection := mgoConn.Coll()
	if collection == nil {
		return errors.New("Mongo DB not Connected")
	}
	if len(u.ID) <= 0 {
		return errors.New("ID is required")
	}

	var found1 User
	filter := bson.M{"_id": bson.M{"$eq": u.ID}}
	err = collection.FindOne(ctx, filter, options.FindOne()).Decode(&found1)
	if err != mgoConn.ErrNoDocuments && err != nil {
		return errors.New(fmts.ConcatStr("Error findOne ID:", err.Error()))
	}
	if len(found1.ID) > 0 {
		return errors.New("ID already exists")
	}

	filter = bson.M{"cpf": bson.M{"$eq": u.Cpf}}
	err = collection.FindOne(ctx, filter, options.FindOne()).Decode(&found1)
	if err != mgoConn.ErrNoDocuments && err != nil {
		fmt.Println("err:", err)
		return errors.New("Error findOne Cpf")
	}
	if len(found1.Cpf) > 0 {
		return errors.New("Cpf already exists")
	}

	filter = bson.M{"email": bson.M{"$eq": u.Email}}
	err = collection.FindOne(ctx, filter, options.FindOne()).Decode(&found1)
	if err != mgoConn.ErrNoDocuments && err != nil {
		fmt.Println("err:", err)
		return errors.New("Error findOne Email")
	}
	if len(found1.Email) > 0 {
		return errors.New("Email already exists")
	}

	password, err := crypt.Blowfish(u.Password)
	if err != nil {
		return
	}
	u.Password = password
	u.ConfirmPassword = ""
	if _, err = collection.InsertOne(ctx, &u, options.InsertOne()); err != nil {
		return
	}

	return
}
