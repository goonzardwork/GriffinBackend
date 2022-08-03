package rdb

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

func JsonGet[T any](db *redis.Client, key, path string, container *T) error {
	cmd := redis.NewStringCmd("JSON.GET", key, path)
	db.Process(cmd)
	v, err := cmd.Result()
	_ = json.Unmarshal([]byte(v), &container)
	return err
}

func JsonSet[T any](db *redis.Client, key, path string, container *T) error {
	content, _ := json.Marshal(container)
	cmd := redis.NewStringCmd("JSON.SET", key, path, string(content))
	db.Process(cmd)
	err := cmd.Err()
	return err
}

func JsonArrAppend[T any](db *redis.Client, key, path, container *T) error {
	content, _ := json.Marshal(container)
	cmd := redis.NewStringCmd("JSON.ARRAPPEND", key, path, string(content))
	db.Process(cmd)
	err := cmd.Err()
	return err
}
