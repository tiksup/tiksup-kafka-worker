/*
* This file contains a repository for inserting data into a redis
* message queue for batch processing.
* Copyright (C) 2024-2025 jsusmachaca
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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
