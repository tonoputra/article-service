package helper

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type dbRedis struct {
	Conn *redis.Client
	Err  error
}

func (_d *dbRedis) Init() {
	_d.Conn, _d.Err = _d.Connect()
}

func (_d *dbRedis) Connect() (*redis.Client, error) {
	var (
		REDIS_HOST     = os.Getenv("REDIS_HOST")
		REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
		REDIS_DB, _    = strconv.Atoi(os.Getenv("REDIS_DB"))
	)

	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST,
		Password: REDIS_PASSWORD,
		DB:       REDIS_DB,
	})

	log.Println("Redis Connected")
	return client, nil
}

func (_d *dbRedis) Get(ctx context.Context, key string) (result string, err error) {
	result, err = _d.Conn.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		log.Println(REDIS_NIL, err.Error())
	case err != nil:
		log.Println(ERR_REDIS_GET_DATA, err.Error())
	case result == "":
		log.Println(REDIS_VAL_EMPTY, err.Error())
	}

	return
}

func (_d *dbRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	if err = _d.Conn.SetEX(ctx, key, value, expiration).Err(); err != nil {
		log.Println(ERR_REDIS_SET_DATA, err.Error())
	}

	return
}

var Redis = &dbRedis{}
