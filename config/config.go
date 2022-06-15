package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Kafka struct {
	Host            string
	TopicTesting    string
	ConsumerGroupId string
}

func init() {
	// Membuat variable environment dari file .env
	if err := godotenv.Load("./config/.env"); err != nil {
		panic("load file env: " + err.Error())
	}

	Kafka.Host = os.Getenv("KAFKA_HOST")
	Kafka.TopicTesting = os.Getenv("KAFKA_TEST_TOPIC")
	Kafka.ConsumerGroupId = os.Getenv("KAFKA_GROUP_ID")
}
