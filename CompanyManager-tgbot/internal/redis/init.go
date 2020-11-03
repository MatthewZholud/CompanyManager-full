package redis

import (
	"github.com/MatthewZholud/CompanyManager-full/CompanyManager-tgbot/internal/logger"
	"github.com/go-redis/redis"
	"os"
)

type redisClient struct {
	client *redis.Client
}

func Initialize(redisUrl string) *redisClient {
	conn := redis.NewClient(&redis.Options{
		Addr: redisUrl,
	})
	if err := conn.Ping().Err(); err != nil {
		logger.Log.Fatalf("Unable to connect to redis: %v", err)
		}
	return &redisClient{
		client: conn,
	}
}

func StartRedis() *redisClient {
	redis := Initialize(os.Getenv("REDIS_URL"))
	return redis
}
