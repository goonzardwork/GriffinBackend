package rdb

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

type DataBase struct {
	Context    context.Context
	Connection *redis.Client
}

func (db DataBase) Connect() DataBase {
	ctx := context.Background()
	cli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("RDB_ADDR"),
		Password: os.Getenv("RDB_PSWD"),
		DB:       0,
	})
	db.Connection, db.Context = cli, ctx
	return db
}
