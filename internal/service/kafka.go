/*
* This file contains a listener function that is responsible for listening
* to the arrival of each message in Kafka to work with them.
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

package service

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tiksup/tiksup-kafka-worker/internal/config"
	"github.com/tiksup/tiksup-kafka-worker/internal/database"
	"github.com/tiksup/tiksup-kafka-worker/pkg/eventstream"
	"github.com/tiksup/tiksup-kafka-worker/pkg/movie"
	"github.com/tiksup/tiksup-kafka-worker/pkg/trigger"
)

func KafkaWorker(
	configMap *kafka.ConfigMap,
	mongoConn database.MongoConnection,
	rdbConn database.RedisConnection,
	gRPC trigger.GRPCRepository,
) {
	var kafkaData eventstream.UserInfo
	movieRepository := &movie.MovieRepository{Database: mongoConn.Database, CTX: mongoConn.CTX}
	rdbRepository := &eventstream.RedisRepository{Database: rdbConn.Database, CTX: rdbConn.CTX}

	consumer, err := config.KafKaConsumer(configMap)
	if err != nil {
		log.Fatalf("Kafka worker error: %v", err)
	}
	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Fatalf("Error getting Kafka information: %v\n", err)
		}

		if err := json.Unmarshal(msg.Value, &kafkaData); err != nil {
			log.Printf("Error to Unmarshall message: %v\n", err)
			continue
		}

		if err := rdbRepository.MessageQueue(kafkaData.UserID, kafkaData); err != nil {
			log.Printf("Error registering user info into queue")
			continue
		}

		if err := movieRepository.InsertHistory(kafkaData.UserID, kafkaData.MovieID); err != nil {
			log.Printf("Error registering history: %v\n", err)
			continue
		}

		if kafkaData.Next {
			if err := trigger.ThrowTrigger(gRPC.Client, gRPC.CTX, kafkaData.UserID); err != nil {
				log.Printf("Error triggering event: %v\n", err)
				continue
			}
		}
		log.Println("User info insert to updated")
	}
}
