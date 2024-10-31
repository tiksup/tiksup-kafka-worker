package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tiksup/tiksup-kafka-worker/internal/config"
	"github.com/tiksup/tiksup-kafka-worker/internal/database"
	"github.com/tiksup/tiksup-kafka-worker/pkg/eventstream"
)

func KafkaWorker(configMap *kafka.ConfigMap, mC database.MongoConnection) error {
	var kafkaData eventstream.KafkaData
	kafaDB := &eventstream.KafkaRepository{Collection: mC.Collection, CTX: mC.CTX}

	consumer, err := config.KafKaConsumer(configMap)
	if err != nil {
		log.Fatalf("Kafka worker error: %v", err)
	}
	defer consumer.Close()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error getting Kafka information: %v\n", err)
		}

		fmt.Printf("%s\n", msg.Value)

		if err := json.Unmarshal(msg.Value, &kafkaData); err != nil {
			log.Fatalf("Error to Unmarshall message: %v\n", err)
		}

		/* if err := kafaDB.UpdateUserInfo(kafkaData); err != nil {
			log.Printf("Error to insert kafka information on database: %v\n", err)
		} */

		if err := kafaDB.UpdateUserInfo(kafkaData); err != nil {
			fmt.Println("Ha ocurrido un error", err)
		}
		fmt.Println("User info insert to database")
		if kafkaData.Next {
			// go MovieWorker(client, db, kafkaData, mC)
			fmt.Println("Requesting for more data")
		}
	}
}
