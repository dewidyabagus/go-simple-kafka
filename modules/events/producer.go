package events

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	session *kafka.Producer
}

// Mengembalikan struct Producer untuk kemudahan penggunaan kafka pada bagian producer
func NewEventProducer(session *kafka.Producer) *Producer {
	return &Producer{session}
}

// Mengirimkan event ke kafka tanpa menunggu hasil pengiriman sukses atau gagal (asynchronous)
func (p *Producer) SendEventAsync(topic string, key []byte, value []byte) error {
	return p.session.Produce(&kafka.Message{
		Key:            key,
		Value:          value,
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	}, nil)
}
