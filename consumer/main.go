package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"learn/kafka/config"
)

// Mensimulasikan consume event dari produce
func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r)
		}
	}()

	consumer, err := config.NewKafkaConsumer()
	if err != nil {
		panic("new kafka consumer: " + err.Error())
	}
	defer consumer.Close()

	if err := consumer.SubscribeTopics([]string{config.Kafka.TopicTesting}, nil); err != nil {
		panic("subscribe topic: " + err.Error())
	}

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)

	for {
		select {
		case <-sigChan:
			log.Println("Consumed shuting down . . .")
			return

		default:
			event, err := consumer.ReadMessage(time.Millisecond * 100)
			if err != nil {
				if !strings.Contains(err.Error(), "Local: Timed out") {
					log.Println("Error Listen Event:", err.Error())
				}
			} else {
				log.Printf("Consumed event from topic %s : key = %-10s value = %s \n", *event.TopicPartition.Topic, string(event.Key), string(event.Value))
			}
		}
	}
}