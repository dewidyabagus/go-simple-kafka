package config

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Membuat session untuk kafka producer
func NewKafkaProducer() (producer *kafka.Producer, err error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": Kafka.Host,
		"acks":              1,
		"retries":           1,
		"retry.backoff.ms":  200,
		"batch.size":        16384 * 4, // 64 KB
		"linger.ms":         100,
	}

	return kafka.NewProducer(config)
}

// Membuat session unyuk kafka consumer
func NewKafkaConsumer(groupID string) (consumer *kafka.Consumer, err error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": Kafka.Host,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}

	return kafka.NewConsumer(config)
}
