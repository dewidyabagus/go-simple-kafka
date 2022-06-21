package order

import "context"

type (
	Service interface {
		CreateNewOrder(orders *NewOrder) error
	}

	Repository interface {
		CreateNewOrder(ctx context.Context, orders []Order) error
		CheckExistingTransNo(transNo string) (exists bool, err error)
	}

	Events interface {
		SendEventAsync(topic string, key []byte, value []byte) error
	}
)
