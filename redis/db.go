package storage

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

var con *redis.Client

// New :: Connect to redis
func New(url string, pass string, db int) error {
	con = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: pass,
		DB:       db,
	})
	_, err := con.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) (interface{}, error) {
	resp, err := con.Get(key).Result()
	if err != nil {
		return nil, err
	}
	var res interface{}
	json.Unmarshal([]byte(resp), &res)
	return res, nil
}

func Push(key string, val interface{}) error {
	str, err := json.Marshal(val)
	if err != nil {
		return err
	}
	err = con.Set(key, str, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
