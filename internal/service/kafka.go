package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tiksup/tiksup-kafka-worker/internal/config"
	"github.com/tiksup/tiksup-kafka-worker/pkg/eventstream"
	"github.com/tiksup/tiksup-kafka-worker/pkg/movie"
)

func KafkaWorker(client *http.Client, configMap *kafka.ConfigMap, db *sql.DB, mC movie.MongoConnection) error {
	var kafkaData eventstream.KafkaData
	kafaDB := &eventstream.KafkaRepository{DB: db}

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

		if err := json.Unmarshal(msg.Value, &kafkaData); err != nil {
			log.Fatalf("Error to Unmarshall message: %v\n", err)
		}

		if err := kafaDB.UpdateUserInfo(kafkaData); err != nil {
			log.Printf("Error to insert kafka information on database: %v\n", err)
		}
		if kafkaData.Next {
			go MovieWorker(client, db, kafkaData, mC)
		}
	}
}
