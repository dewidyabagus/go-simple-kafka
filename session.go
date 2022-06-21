package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"

	"learn/kafka/api"

	// Module Repository
	orderRepository "learn/kafka/modules/db/order"

	// Module Business
	orderBusiness "learn/kafka/business/order"

	// Module Controller
	orderHandler "learn/kafka/api/v1/order"
	welcomeHandler "learn/kafka/api/v1/welcome"
)

type (
	Session struct {
		db            *gorm.DB
		kafkaProducer *kafka.Producer
	}
)

func (s *Session) App() *api.Routes {
	// Grup untuk instan repository
	orderRepo := orderRepository.NewRepository(s.db)

	// Grup untuk instan business / use case
	orderService := orderBusiness.NewService(orderRepo)

	return &api.Routes{
		Welcome: welcomeHandler.NewController(),
		Order:   orderHandler.NewController(orderService),
	}
}
