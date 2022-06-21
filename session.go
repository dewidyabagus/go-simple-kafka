package main

import (
	"gorm.io/gorm"

	"learn/kafka/api"
	"learn/kafka/config"

	// Module Repository
	orderRepository "learn/kafka/modules/db/order"
	"learn/kafka/modules/events"

	// Module Business
	orderBusiness "learn/kafka/business/order"

	// Module Controller
	orderHandler "learn/kafka/api/v1/order"
	welcomeHandler "learn/kafka/api/v1/welcome"
)

type (
	Session struct {
		db            *gorm.DB
		kafkaProducer *events.Producer
	}
)

func (s *Session) App() *api.Routes {
	// Repository Domain
	orderRepo := orderRepository.NewRepository(s.db)

	// Business Domain
	orderService := orderBusiness.NewService(orderRepo, s.kafkaProducer, config.Kafka)

	// Return List Routes REST API
	return &api.Routes{
		Welcome: welcomeHandler.NewController(),
		Order:   orderHandler.NewController(orderService),
	}
}
