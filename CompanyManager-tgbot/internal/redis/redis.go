package redis

import (
	"encoding/json"
	"fmt"
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
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

func (r *redisClient) Set(msg int) {
	slice, b := r.needChange(msg)
	if b{
		slice = append(slice, msg)
		cacheEntry, err := json.Marshal(slice)
		if err != nil {
			logger.Log.Errorf("Can't marshal slice to send to cash: %v", err)
		}
		err = r.client.Set("UsersIDs", cacheEntry, 0).Err()
		if err != nil {
			logger.Log.Errorf("Can't set slice to cash: %v", err)
		}
	}
}

func (r *redisClient) Get() ([]int, error) {
	val, err := r.client.Get("UsersIDs").Result()


	var slice []int
	err = json.Unmarshal([]byte(val), &slice)

	fmt.Println(slice)
	if err != nil {
		return nil, err
	}

	return slice, nil
}

func (r *redisClient) needChange(msg int)  ([]int, bool) {
	slice, _ := r.Get()

	for _, item := range slice {
		if msg == item {
			return nil, false
		}
	}
	return slice, true
}