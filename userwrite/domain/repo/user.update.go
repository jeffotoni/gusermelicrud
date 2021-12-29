package drepo

import (
	"context"
	"encoding/json"
	"errors"

	mgoConn "github.com/jeffotoni/gusermeli/apicore/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type UserUp struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Birthday  string `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Cpf       string `json:"cpf,omitempty" bson:"cpf,omitempty"`

	UpdatedAt string `json:"update,omitempty" bson:"update,omitempty"`
	IP        string `json:"ip,omitempty" bson:"ip,omitempty"`
	Agent     string `json:"agent,omitempty" bson:"agent,omitempty"`
}

func (u *UserUp) Update(ctx context.Context) (err error) {
	collection := mgoConn.Coll()
	if collection == nil {
		return errors.New("Mongo DB not Connected")
	}

	var c UserUp
	filter := bson.M{"cpf": bson.M{"$eq": u.Cpf}}
	err = collection.FindOne(
		ctx,
		filter,
	).Decode(&c)

	if err != nil {
		return
	}

	var newCc UserUp
	cc0, _ := json.Marshal(&c)
	json.Unmarshal(cc0, &newCc)

	cc1, _ := json.Marshal(&u)
	json.Unmarshal(cc1, &newCc)

	update := bson.M{"$set": newCc}
	if _, err = collection.UpdateOne(ctx, filter, update); err != nil {
		return
	}

	return
}
