package config

import (
	"learn/kafka/utils/configuration"
	"os"

	"github.com/joho/godotenv"
)

var (
	Kafka = &configuration.Kafka{}

	App struct {
		Env        string
		Host       string
		ListenPort string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Schema   string
		TimeZone string
	}
)

func init() {
	// Membuat variable environment dari file .env
	if err := godotenv.Load("./config/.env"); err != nil {
		panic("load file env: " + err.Error())
	}

	// Variable konfigurasi aplikasi

	// Runing Apps
	App.Env = os.Getenv("ENV")
	App.Host = os.Getenv("APP_HOST")
	App.ListenPort = os.Getenv("APP_PORT")

	// Kafka Config
	Kafka.Host = os.Getenv("KAFKA_HOST")
	Kafka.TopicOrder = os.Getenv("KAFKA_ORDER_TOPIC")
	Kafka.ShipmentGroup = os.Getenv("KAFKA_SHIPMENT_GRUOP")
	Kafka.InventoryGroup = os.Getenv("KAFKA_INVENTORY_GROUP")

	// Database Config
	Database.Host = os.Getenv("DATABASE_HOST")
	Database.Port = os.Getenv("DATABASE_PORT")
	Database.User = os.Getenv("DATABASE_USER")
	Database.Password = os.Getenv("DATABASE_PASSWORD")
	Database.Schema = os.Getenv("DATABASE_SCHEMA")
	Database.TimeZone = os.Getenv("DATABASE_TIMEZONE")
}
