package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/jeffotoni/gusermeli/apicore/pkg/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ErrNoDocuments = mongo.ErrNoDocuments
	client         *mongo.Client
	collectionName = env.GetString("MONGODB_COLLECTION", "users")
	mgoSchema      = env.GetString("MONGODB_SCHEMA", "mongodb")
	mgoURI         = env.GetString("MONGODB_URI", "localhost:27017")
	mgoUser        = env.GetString("MONGODB_USERNAME", "root")
	mgoPass        = env.GetString("MONGODB_PASSWORD", "123456")
	mgoDb          = env.GetString("MONGODB_DATABASE", "meliusers")
	mgoOpts        = env.GetString("MONGODB_OPTIONS", "authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false")
	mgoPingTimeout = env.GetDuration("MONGODB_PING_TIMEOUT", 5*time.Second)
	mgoConnTimeout = env.GetDuration("MONGODB_CONNECTION_TIMEOUT", 5*time.Second)
	mgoMinPoolSize = env.GetInt("MONGODB_MIN_POOL", 0)
	mgoMaxPoolSize = env.GetInt("MONGODB_MAX_POOL", 100)
)

func stringURI() (connectionString string) {
	if len(mgoUser) > 0 && len(mgoPass) > 0 {
		connectionString = mgoSchema + "://" + mgoUser + ":" + mgoPass + "@" + mgoURI + "/" + mgoDb + "?" + mgoOpts
	} else {
		connectionString = mgoSchema + "://" + mgoURI + "/" + mgoDb + "?" + mgoOpts
	}
	return
}

func Connect() (ctx context.Context, err error) {
	opts := options.Client()
	fmt.Println(stringURI())
	opts.ApplyURI(stringURI())
	min := uint64(mgoMinPoolSize)
	opts.MinPoolSize = &min
	max := uint64(mgoMaxPoolSize)
	opts.MaxPoolSize = &max
	client, err = mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}
	ctx, _ = context.WithTimeout(context.Background(), mgoConnTimeout)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func Disconnect(ctx context.Context) {
	client.Disconnect(ctx)
}

func Coll() *mongo.Collection {
	if client == nil {
		//log.Println("Mongo DB not Connected")
		return nil
	}
	return client.Database(mgoDb).Collection(collectionName)
}

func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), mgoPingTimeout)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}
