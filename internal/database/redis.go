package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func GetRedisConnection(ctx context.Context) (*redis.Client, error) {
	REDIS_PORT := os.Getenv("REDIS_PORT")
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	uri := fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT)

	client := redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: REDIS_PASSWORD,
		DB:       0,
		Protocol: 2,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	log.Println("\033[32mCONNECTED TO REDIS DATABASE\033[0m")

	return client, nil
}
