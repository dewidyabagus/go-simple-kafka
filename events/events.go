package events

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Digunakan untuk menerima pesan status pengiriman event ke topic kafka secara asynchronous melalui chanel
func AnsyncListenKafkaDelivery(producer *kafka.Producer) {
	go func() {
		for event := range producer.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message topic: %s \n", *ev.TopicPartition.Topic)
				} else {
					log.Printf("Produced event to topic %s: key = %-10s value = %s", *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()
}
