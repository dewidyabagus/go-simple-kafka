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

var App struct {
	Env        string
	Host       string
	ListenPort string
}

func init() {
	// Membuat variable environment dari file .env
	if err := godotenv.Load("./config/.env"); err != nil {
		panic("load file env: " + err.Error())
	}

	// Variable konfigurasi aplikasi
	App.Env = os.Getenv("ENV")
	App.Host = os.Getenv("APP_HOST")
	App.ListenPort = os.Getenv("APP_PORT")
	Kafka.Host = os.Getenv("KAFKA_HOST")
	Kafka.TopicTesting = os.Getenv("KAFKA_TEST_TOPIC")
	Kafka.ConsumerGroupId = os.Getenv("KAFKA_GROUP_ID")
}
