package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"os"
)


type redisClient struct {
	client *redis.Client
}

func Initialize() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})
	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	return &redisClient{client: c}
}


func (r *redisClient) Set( msg int)  error {
	val, err := r.client.Get("UsersID").Result()

	var slice []int
	err = json.Unmarshal([]byte(val), &slice)
	for i := range slice{
		if msg == slice[i]{
			var ErrInvalidEntity = errors.New("Invalid presenter")
			return ErrInvalidEntity
		}
	}
	slice = append(slice, msg)
	cacheEntry, err := json.Marshal(slice)
	if err != nil {
		return err
	}

	err = r.client.Set("UsersID", cacheEntry, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisClient) Get() ([]int, error) {
	val, err := r.client.Get("UsersID").Result()
	if err == redis.Nil || err != nil {
		return nil, err
	}

	var slice []int
	err = json.Unmarshal([]byte(val), &slice)

	if err != nil {
		return nil, err
	}

	fmt.Println(slice)
	return slice, nil
}