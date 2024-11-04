package eventstream

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Database *redis.Client
	CTX      context.Context
}

func (rClient *RedisRepository) MessageQueue(useID string, doc KafkaData) error {
	rdb := rClient.Database
	ctx := rClient.CTX

	jsonData, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("user:%s:documents", useID)
	err = rdb.RPush(ctx, key, jsonData).Err()
	if err != nil {
		return err
	}

	return nil
}
