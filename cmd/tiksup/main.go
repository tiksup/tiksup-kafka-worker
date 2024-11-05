/*
* This file is the entry point of the project,
* main.go runs and initializes the entire project.
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

/*
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⡀⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿
⠀⠀⢘⢸⠀⡇⠁⠰⠀⡀⠆⢀⢰⠁⡇⡆⢸⢀⣁⣗⣺⢘⡀⡇⣱⠸⠀⠀⠆⠸⠐⠀⡇⣎⢸⠐⡃⣗⢸⣈⡂⡇⢰⢸⠀⡆⠂⢀⠰⠀⡄⠘⣿⣿⣿⣿⣿⣿
⠀⠀⠜⢸⠀⡇⠃⢠⠀⡄⡄⢸⢠⠀⡇⡄⢸⠤⠇⡧⢼⠸⠄⡇⢣⢰⠀⡀⡄⢸⢠⠀⡇⡼⢸⠤⠇⡧⢼⠸⠤⡇⢢⢰⠀⡅⡄⠈⢠⠀⡃⠸⣿⣿⣿⣿⣿⣿
⠀⠀⢱⢸⠀⡇⠀⢸⠀⣩⣇⢸⢨⡀⡏⡁⢸⠉⡆⣏⣹⢸⠁⡇⣎⢸⠀⠁⡁⠘⢈⠀⠃⣱⢸⣉⣇⣏⣽⣰⠁⡇⢈⢸⠀⠁⠀⠈⠈⠀⠃⢰⣿⣿⣿⣿⣿⣿
⠀⠀⢠⢰⠀⡁⡄⢈⠀⢻⣿⣷⣬⣤⠇⠣⢸⠐⡤⡇⢸⢰⠀⡇⡜⢈⠀⠀⠁⢠⠈⠀⡇⢣⢸⢐⣾⣿⣿⣦⠀⡇⠘⢘⠀⠅⠁⢠⠈⠀⡄⢠⣿⣿⣿⣿⣿⣿
⠀⠀⢈⢸⠀⡇⠁⠘⠀⡇⣿⡏⠙⢿⣷⣭⣹⢬⡉⡇⣼⣈⠄⡇⡱⢸⠀⡄⡇⢸⢸⡁⡇⣮⣽⣿⡿⠃⣿⣯⠄⡇⢹⢸⠀⠇⠂⠀⠰⠀⡀⢈⣿⣿⣿⣿⣿⣿
⠀⠀⠘⠘⠀⡇⠀⠸⠀⡄⣿⡇⠀⠀⠉⠿⣿⣷⣅⡇⢾⢘⣖⣇⣣⣸⣠⣄⣇⣸⣤⣤⣿⣿⠏⠉⢀⠀⢹⣿⠀⡇⢠⢸⠀⠀⠄⠀⢠⠀⠀⠘⣿⣿⣿⣿⣿⣿
⠀⠀⠸⢸⠀⡇⠇⢰⠀⠁⣻⣿⠀⠀⠀⠀⠈⠙⠻⣿⣿⠿⠛⠿⠛⠛⠛⠛⠛⠿⠿⠿⣿⡄⠀⠀⠀⠀⠘⣿⣇⡇⢍⢸⠠⡃⡤⢸⢀⠀⡃⠸⣿⣿⣿⣿⣿⣿
⠀⠀⢲⢸⠒⡇⠆⢨⠀⠃⢸⣿⠀⠀⠀⠀⠀⠀⠀⠀⠙⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣷⡚⢺⠐⡇⠅⠸⠘⠀⠇⢲⣿⣿⣿⣿⣿⣿
⠀⠀⢥⢸⠤⡇⡄⢸⢠⠆⢼⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣼⡷⠾⢿⣬⡻⣿⣯⡤⡧⡇⢰⠼⠀⡦⢤⣿⣿⣿⣿⣿⣿
⠀⠀⣚⢸⡂⡇⡃⢸⠘⣆⡾⣿⡆⠀⠀⠀⠀⢀⣤⣾⠿⠿⠶⣶⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⠁⠀⠀⠀⠈⠻⣿⣿⣿⣖⠏⢸⢰⠃⡇⢚⣿⣿⣿⣿⣿⣿
⠀⠀⠼⢸⠠⡇⠁⢸⠀⣤⣼⡿⡃⠀⠀⠀⠠⣾⠏⠀⠀⠀⠀⠀⠙⢿⣷⣄⠀⠀⠀⠀⢐⠀⢿⣇⠀⠀⠀⠀⠀⠀⠈⢿⣞⣿⣧⢘⢠⠀⠁⠸⣿⣿⣿⣿⣿⣿
⠀⠀⣱⢸⠁⡇⣆⢸⣲⣿⡟⠀⠀⠀⠀⠀⢰⣿⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⡀⠀⠀⠀⠀⠀⠈⣿⡀⠀⠀⠀⠀⠀⠀⠀⢿⣾⣿⣏⢈⠀⡇⢰⣿⣿⣿⣿⣿⣿
⠀⠀⢠⢸⠀⡇⠄⣸⣿⡟⠀⠀⠀⠀⠀⠁⠈⣿⠀⠀⠀⠀⠀⠀⡀⠀⠀⢹⡧⠀⠀⠀⠀⠀⠀⠺⣷⡄⠀⠀⢰⣾⣶⠀⢸⣿⢻⣿⡜⠀⡇⢠⣿⣿⣿⣿⣿⣿
⠰⢶⣭⣸⢈⡇⡈⣿⡿⠀⠀⠀⠀⠀⠀⠀⠀⠻⣧⡀⠀⠀⠀⣿⣿⠀⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⠉⢿⣦⡀⠀⠙⠁⢀⣸⡿⣆⣿⣿⣀⠆⣈⣿⣿⣿⣿⣿⣿
⠀⠀⠘⢹⠻⠶⣵⣿⠃⠀⠀⠀⠀⠀⠀⠀⢀⠈⠹⣷⣄⠀⠀⠈⠁⠀⣀⣼⠿⣂⡀⢠⣦⣦⣶⣶⣶⣶⡟⠛⠷⠦⢤⡾⠋⠀⠈⠙⠛⠿⣷⣾⣿⣿⣿⣿⣿⣿
⠀⠀⡹⢼⣀⡇⢸⣿⠷⠆⠀⠀⠀⠀⠀⠀⠀⠀⢀⣬⣽⠿⢶⣤⣤⣴⠿⠋⠀⠀⠉⠉⠻⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣨⣿⣿⣿⣿⣿⣿⣿
⢀⣀⣰⣸⣀⣇⣸⣿⣄⡀⠀⠀⠀⠀⠀⠀⢰⠞⠉⠀⠀⠀⠀⠀⠀⢀⡀⠀⠀⠀⠀⠀⠀⠀⠉⠛⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠁⣿⣿⣿⣿⣿⣿
⠈⠉⢉⢹⠛⡏⡽⣿⣿⠀⠀⠀⠀⠀⠄⠆⡇⠀⠀⠀⠀⠀⠀⠀⠀⣿⠛⠛⠶⣦⣤⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⣿⣿⣿⣿⣿⣿
⠀⠀⡚⢺⣒⡗⢓⠻⣿⣷⡀⠀⠀⠀⠂⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⡟⠀⠀⠀⠀⠀⠉⠙⠛⠶⠶⣤⣤⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⣿⣿⣿⣿⣿⣿
⠀⠀⠼⢸⠀⡇⠧⢸⠙⣿⣿⣄⠀⠀⢀⠀⠸⡀⠀⠀⠀⠀⠀⠀⠀⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠙⣛⣿⠟⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⣿
⠀⠀⣰⢸⠀⡇⡆⢸⡀⡉⠻⣿⣵⣄⡀⠀⠀⠙⢄⠀⠀⠀⠀⠀⠀⠘⢷⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⡾⠛⠁⠀⠀⠀⠀⠀⠀⣠⣿⡿⣿⣿⣿⣿⣿⣿
⠀⠀⢢⢸⠀⡇⠄⢈⠀⠇⠧⢨⠻⣿⣿⣢⢀⠀⠈⢑⢦⡀⠀⠀⠀⠀⠀⠉⠛⠶⣦⣤⣀⣀⣀⣤⣶⠾⠛⠉⠀⠀⠀⠀⠀⢀⣠⣶⣿⡿⡏⢦⣿⣿⣿⣿⣿⣿
⠀⠀⢈⢈⠀⠇⡀⠐⠀⡀⠇⢸⢰⣀⡟⠻⣿⣿⣶⣠⣄⢚⣷⠦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣤⣶⣿⣿⢛⢹⡈⡇⠈⣿⣿⣿⣿⣿⣿
⠀⠀⠸⢸⠀⡇⠃⠠⠀⠀⡄⠈⢰⠂⡇⡔⢸⠉⠛⡿⢿⢿⣿⣿⣿⣭⣷⣶⣤⣄⣀⣀⣤⣤⣤⣤⣤⣴⣶⣾⡿⡿⢿⢻⠙⡆⡗⢺⢠⠀⠁⠘⣿⣿⣿⣿⣿⣿
⠀⠀⠹⢸⠀⡇⠁⢀⠀⠁⡁⠈⢈⠀⡏⡉⢸⠉⠇⣏⣹⢸⡀⡏⢏⠉⠉⠙⡛⢛⢛⠛⡉⡻⢹⢉⡍⣏⣸⠹⣉⡇⢈⢹⠈⠃⠁⠈⢈⠀⠃⠸⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
*/
package main

import (
	"context"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/tiksup/tiksup-kafka-worker/internal/config"
	"github.com/tiksup/tiksup-kafka-worker/internal/database"
	"github.com/tiksup/tiksup-kafka-worker/internal/service"
	"github.com/tiksup/tiksup-kafka-worker/pkg/trigger"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db        *mongo.Database
	configMap kafka.ConfigMap
	ctx       = context.TODO()
	rctx      = context.Background()
	mongoConn database.MongoConnection
	redisConn database.RedisConnection
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("\033[31mNot .env file found. Using system variables\033[0m")
	}

	configMap = config.KafkaConfig()

	var err error
	db, err = database.GetMongoConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to mongo: %v", err)
	}

	rdb, err := database.GetRedisConnection(rctx)
	if err != nil {
		log.Fatal("Error trying connect to redis")
	}

	mongoConn = database.MongoConnection{Database: db, CTX: ctx}
	redisConn = database.RedisConnection{Database: rdb, CTX: rctx}
}

func main() {
	GRPC_SERVER := os.Getenv("GRPC_SERVER")

	client, err := config.CreateEventClient(GRPC_SERVER)
	if err != nil {
		log.Fatalf("Error trying connect to grpc server: %v", err)
	}
	defer client.Close()

	gRPC := trigger.GRPCRepository{Client: client.Client, CTX: context.Background()}
	service.KafkaWorker(&configMap, mongoConn, redisConn, gRPC)
}
