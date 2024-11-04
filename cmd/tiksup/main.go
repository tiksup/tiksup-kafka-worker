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

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/tiksup/tiksup-kafka-worker/internal/config"
	"github.com/tiksup/tiksup-kafka-worker/internal/database"
	"github.com/tiksup/tiksup-kafka-worker/internal/service"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db        *mongo.Database
	configMap kafka.ConfigMap
	ctx       = context.TODO()
	mongoConn database.MongoConnection
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
	mongoConn = database.MongoConnection{Database: db, CTX: ctx}
}

func main() {
	service.KafkaWorker(&configMap, mongoConn)
}
