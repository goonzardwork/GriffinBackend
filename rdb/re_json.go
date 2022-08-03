package rdb

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

func JsonGet[T any](db *redis.Client, key, path string, container *T) error {
	cmd := redis.NewStringCmd("JSON.GET", key, path)
	db.Process(cmd)
	fmt.Println(cmd)
	v, err := cmd.Result()
	_ = json.Unmarshal([]byte(v), &container)
	return err
}

func JsonSet[T any](db *redis.Client, key, path string, container *T) error {
	content, _ := json.Marshal(container)
	cmd := redis.NewSliceCmd("JSON.SET", key, path, string(content))

	db.Process(cmd)
	err := cmd.Err()
	return err
}

func JsonArrAppend[T any](db *redis.Client, key, path string, container *T) error {
	content, _ := json.Marshal(container)
	cmd := redis.NewSliceCmd("JSON.ARRAPPEND", key, path, string(content))

	db.Process(cmd)
	err := cmd.Err()
	return err
}

func JsonArrLen(db *redis.Client, key, path string) (int, error) {
	cmd := redis.NewSliceCmd("JSON.ARRLEN", key, path)
	db.Process(cmd)
	l, err := cmd.Result()
	if err != nil {
		return -1, err
	}
	intL, err := strconv.Atoi(fmt.Sprintf("%v", l[0]))
	return intL, err
}
