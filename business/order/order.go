package order

import "time"

type (
	Item struct {
		ItemID    uint    `validate:"required,min=1"`
		ItemPrice float64 `validate:"required,min=0"`
		Qty       uint    `validate:"required,min=1"`
	}

	NewOrder struct {
		TransactionNo string `validate:"required"`
		Items         []Item `validate:"dive"`
		Date          string `validate:"required,datetime=2006-01-02 15:04:05"`
	}

	Order struct {
		ID            uint
		TransactionNo string
		ItemID        uint
		ItemPrice     float64
		Qty           uint
		Date          time.Time
		CreatedAt     time.Time
		UpdatedAt     time.Time
		DeletedAt     time.Time
	}
)
