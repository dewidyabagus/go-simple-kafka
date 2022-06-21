package order

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"

	"learn/kafka/business/order"
)

const (
	timeOut time.Duration = 10
)

type Order struct {
	ID            uint         `gorm:"column:id;type:bigserial;autoIncrement;primaryKey"`
	TransactionNo string       `gorm:"column:transaction_no;type:varchar(45);unique;not null"`
	ItemID        uint         `gorm:"column:item_id;type:integer;index:orders_index_outlet_item,sort:asc;not null"`
	ItemPrice     float64      `gorm:"column:item_price;type:numeric;not null"`
	Qty           uint         `gorm:"column:qty;type:smallint;not null"`
	Date          time.Time    `gorm:"column:date;type:timestamp;default:now();not null"`
	CreatedAt     time.Time    `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time    `gorm:"column:updated_at"`
	DeletedAt     sql.NullTime `gorm:"column:deleted_at;index"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateNewOrder(ctx context.Context, orders []order.Order) error {
	ctxWT, cancel := context.WithTimeout(ctx, timeOut*time.Second)
	defer cancel()

	return r.db.WithContext(ctxWT).Create(r.toModelOrder(orders)).Error
}

func (r *Repository) toModelOrder(orders []order.Order) []Order {
	modelOrders := make([]Order, len(orders))

	for i, item := range orders {
		modelOrders[i] = Order{
			TransactionNo: item.TransactionNo,
			ItemID:        item.ItemID,
			ItemPrice:     item.ItemPrice,
			Qty:           item.Qty,
			Date:          item.Date,
		}
	}

	return modelOrders
}