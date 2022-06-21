package order

import "time"

type (
	Item struct {
		ItemID    uint    `validate:"required,min=1" json:"item_id"`
		ItemPrice float64 `validate:"required,min=0" json:"item_price"`
		Qty       uint    `validate:"required,min=1" json:"qty"`
	}

	NewOrder struct {
		TransactionNo string `validate:"required" json:"transaction_no"`
		Items         []Item `validate:"dive" json:"items"`
		Date          string `validate:"required,datetime=2006-01-02 15:04:05" json:"date"`
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

	Kafka struct {
		Host            string
		TopicOrder      string
		ConsumerGroupId string
	}
)
