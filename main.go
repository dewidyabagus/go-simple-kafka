package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"

	"learn/kafka/config"
	"learn/kafka/modules/events"
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
	// kafkaDelivery := events.NewEventProducer(producer)

	// Membuat session database
	db, err := config.NewDatabase()
	if err != nil {
		panic("new database:" + err.Error())
	}

	session := &Session{
		db:            db,
		kafkaProducer: producer,
	}

	// Instan Echo Web Framework
	e := echo.New()

	session.App().New(e)

	go func() {
		if err := e.Start(fmt.Sprintf("%s:%s", config.App.Host, config.App.ListenPort)); err != nil {
			log.Println("Shutting Down REST Service")
			os.Exit(0)
		}
	}()

	quit := make(chan os.Signal, 10)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Println("Shutting down error:", err.Error())
	}
}
