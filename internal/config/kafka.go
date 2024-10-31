/*
* This file creates the connection to some kafka server
* and initializes a consumer.
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

package config

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaConfig() kafka.ConfigMap {
	configMap := kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KAFKA_SERVER"),
		"group.id":           "user-info-consumer",
		"auto.offset.reset":  "latest",
		"session.timeout.ms": 300000,
	}
	return configMap
}

func KafKaConsumer(configMap *kafka.ConfigMap) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, err
	}

	topic := os.Getenv("KAFKA_TOPIC")
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}
	log.Println("\033[32mSUBSCRIBE TO KAFKA TOPIC\033[0m")

	return consumer, nil
}
