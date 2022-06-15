package main

import (
	"learn/kafka/config"
	"learn/kafka/events"
	"log"
	"os"
	"os/signal"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r)
		}
	}()

	// Membuat session kafka bagian producer
	producer, err := config.NewKafkaProducer()
	if err != nil {
		panic("new kafka producer:" + err.Error())
	}
	defer producer.Close()

	events.AnsyncListenKafkaDelivery(producer)
	kafkaDelivery := events.NewEventProducer(producer)

	// dummy data
	// myTopic := os.Getenv("KAFKA_TEST_TOPIC")
	users := []string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := []string{"book", "alarm clock", "t-shirts", "gift card", "batteries", "xxxxx"}

	for i := 0; i < len(users); i++ {
		key := users[i]
		data := items[i]

		kafkaDelivery.SendEventAsync(config.Kafka.TopicTesting, []byte(key), []byte(data))
	}

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)

	<-quit

}
