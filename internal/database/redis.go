/*
* This file initializes a new connection to the redis database,
* returns a client to work with the queries.
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
