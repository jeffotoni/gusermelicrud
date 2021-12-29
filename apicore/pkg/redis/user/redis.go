package redisu

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	DB     = 0
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       DB,
	})
)

var ctx = context.Background()

func SetDB(db int) {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
	return
}

func Put(KEY, pJSON string) (err error) {
	if len(pJSON) == 0 {
		return errors.New("json is required")
	}
	err = client.Set(ctx, KEY, pJSON, 0).Err()
	return
}

func PutTTL(KEY, pJSON string, dt time.Duration) (err error) {
	if len(pJSON) == 0 {
		return errors.New("json is required")
	}
	err = client.Set(ctx, KEY, pJSON, dt).Err()
	return
}

func Get(KEY string) (valJson string, err error) {
	if len(KEY) == 0 {
		err = errors.New("key is required")
		return
	}
	valJson, err = client.Get(ctx, KEY).Result()
	return
}

func Del(KEY string) (err error) {
	if len(KEY) == 0 {
		err = errors.New("key is required")
		return
	}
	client.Del(ctx, KEY)
	return
}

func Expire(key string, td time.Duration) (err error) {
	_, err = client.Expire(ctx, key, td).Result()
	return
}
