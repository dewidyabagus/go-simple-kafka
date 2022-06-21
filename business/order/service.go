package order

import (
	"context"
	"fmt"
	"time"

	"learn/kafka/utils/validator"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) CreateNewOrder(newOrder *NewOrder) error {
	if err := validator.GetValidator().Struct(newOrder); err != nil {
		return fmt.Errorf("validate new order: %s", err.Error())
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

	return s.repository.CreateNewOrder(context.Background(), orders)
}
