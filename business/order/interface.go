package order

import "context"

type (
	Service interface {
		CreateNewOrder(orders *NewOrder) error
	}

	Repository interface {
		CreateNewOrder(ctx context.Context, orders []Order) error
	}
)
