/*
* This file initializes a new connection to the mongodb database,
* returns a collection to work with the queries.
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

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection(ctx context.Context) (*mongo.Database, error) {
	MONGO_HOST := os.Getenv("MONGO_HOST")
	MONGO_PORT := os.Getenv("MONGO_PORT")
	MONGO_USER := os.Getenv("MONGO_USER")
	MONGO_PASSWORD := os.Getenv("MONGO_PASSWORD")
	MONGO_DB := os.Getenv("MONGO_DATABASE")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", MONGO_USER, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT)

	clientOpions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	database := client.Database(MONGO_DB)

	log.Println("\033[32mCONNECTED TO MONGODB DATABASE\033[0m")
	return database, nil
}
