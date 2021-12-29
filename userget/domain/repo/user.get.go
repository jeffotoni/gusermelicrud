package drepo

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	mgoConn "github.com/jeffotoni/gusermeli/apicore/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userGet struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Birthday  string `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Cpf       string `json:"cpf,omitempty" bson:"cpf,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`

	CreatedAt string `json:"create,omitempty" bson:"create,omitempty"`
	UpdatedAt string `json:"update,omitempty" bson:"update,omitempty"`
	IP        string `json:"ip,omitempty" bson:"ip,omitempty"`
	Agent     string `json:"agent,omitempty" bson:"agent,omitempty"`
}

// db.users.createIndex( { first_name: "text" } )
func Get(ctx context.Context, generic []string) (userJson string, err error) {
	collection := mgoConn.Coll()
	if collection == nil {
		return "", errors.New("Mongo DB not Connected")
	}

	findOptions := options.Find()
	findOptions.SetLimit(30)

	// $text: { $search: "Jefferson Paul" }
	filter := bson.M{"$text": bson.M{"$search": strings.Join(generic, " ")}}
	cur, err := collection.Find(
		ctx,
		filter,
		findOptions,
	)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var results []userGet
	for cur.Next(ctx) {
		var elem userGet
		err := cur.Decode(&elem)
		if err != nil {
			log.Println(err)
			continue
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
		return "", err
	}
	cur.Close(ctx)
	if len(results) > 0 {
		b, err := json.Marshal(&results)
		if err != nil {
			return "", err
		}
		userJson = string(b)
	}

	return
}
