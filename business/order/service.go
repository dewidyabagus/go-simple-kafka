package order

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"learn/kafka/utils/configuration"
	"learn/kafka/utils/validator"
)

type service struct {
	repository Repository
	events     Events
	kafkaInfo  *configuration.Kafka
}

func NewService(repository Repository, events Events, kafkaInfo *configuration.Kafka) Service {
	return &service{
		repository: repository,
		events:     events,
		kafkaInfo:  kafkaInfo,
	}
}

func (s *service) CreateNewOrder(newOrder *NewOrder) error {
	if err := validator.GetValidator().Struct(newOrder); err != nil {
		return fmt.Errorf("validate new order: %s", err.Error())
	}

	if found, err := s.repository.CheckExistingTransNo(newOrder.TransactionNo); err != nil {
		return fmt.Errorf("error validate trans no: %s", err.Error())

	} else if found {
		return fmt.Errorf("transaction no already exists")

	}

	payload, err := json.Marshal(newOrder)
	if err != nil {
		return fmt.Errorf("marshal payload: %s", err.Error())
	}

	date, _ := time.ParseInLocation("2006-01-02 15:04:05", newOrder.Date, time.Local)

	orders := make([]Order, len(newOrder.Items))
	for i, item := range newOrder.Items {
		orders[i] = Order{
			TransactionNo: newOrder.TransactionNo,
			ItemID:        item.ItemID,
			ItemPrice:     item.ItemPrice,
			Qty:           item.Qty,
			Date:          date,
		}
	}

	if err := s.repository.CreateNewOrder(context.Background(), orders); err != nil {
		return fmt.Errorf("SQL Error: %s", err.Error())
	}

	// alternative menggunakan os.Getenv() dengan mendistribusikan file config yang dibutuhkan ke business
	return s.events.SendEventAsync(s.kafkaInfo.TopicOrder, []byte(newOrder.TransactionNo), payload)
}

func (s *service) GetOrderByTransNo(transNo string) (orders []Order, err error) {
	return s.repository.GetOrderByTransNo(context.Background(), transNo)
}
