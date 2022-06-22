package order

import "context"

type (
	Service interface {
		CreateNewOrder(orders *NewOrder) error
		GetOrderByTransNo(transNo string) (orders []Order, err error)
	}

	Repository interface {
		CreateNewOrder(ctx context.Context, orders []Order) error
		CheckExistingTransNo(transNo string) (exists bool, err error)
		GetOrderByTransNo(ctx context.Context, transNo string) (orders []Order, err error)
	}

	Events interface {
		SendEventAsync(topic string, key []byte, value []byte) error
	}
)
