package rdb

import (
	"os"

	"github.com/go-redis/redis"
)

type DataBase struct {
	Connection *redis.Client
}

type DAO interface {
	Connect() DataBase
}

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RDB_ADDR"),
		Password: os.Getenv("RDB_PSWD"),
		DB:       0,
	})

	return client
}
