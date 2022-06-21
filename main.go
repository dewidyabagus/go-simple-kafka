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
	// Cover error from panic function
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic Error: ", r)
		}
	}()

	// Create kafka session for producer
	producer, err := config.NewKafkaProducer()
	if err != nil {
		panic("new kafka producer:" + err.Error())
	}
	defer producer.Close()

	// Membuat layanan dibelakang layar untuk mengambil informasi hasil publish event ke topic
	// apakah berhasil ter-publish atau gagal
	events.AnsyncListenKafkaDelivery(producer)

	// Create database session (PostgreSQL)
	db, err := config.NewDatabase()
	if err != nil {
		panic("new database:" + err.Error())
	}

	// Mengumpulkan semua sesi dan mendistribusikan ke masing-masing domain yang menggunakan
	session := &Session{
		db:            db,
		kafkaProducer: events.NewEventProducer(producer),
	}

	// Instan Echo Web Framework
	e := echo.New()

	// Meregistrasikan hasil dari setup domain modules dan business menjadi suatu layanan REST API
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
