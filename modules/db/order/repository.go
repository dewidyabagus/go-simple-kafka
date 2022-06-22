package order

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (r *Repository) CheckExistingTransNo(transNo string) (exists bool, err error) {
	var count int64

	err = r.db.Model(&Order{}).Where("transaction_no = ?", transNo).Count(&count).Error

	return (count != 0), err
}

func (r *Repository) GetOrderByTransNo(ctx context.Context, transNo string) (orders []order.Order, err error) {
	ctxWT, cancel := context.WithTimeout(ctx, timeOut*time.Second)
	defer cancel()

	modelOrders := []Order{}

	if err := r.db.WithContext(ctxWT).Order("id").Find(&modelOrders, "transaction_no = ?", transNo).Error; err != nil {
		return nil, fmt.Errorf("SQL Error: %s", err.Error())

	} else if len(modelOrders) == 0 {
		return nil, errors.New("error record not found")

	}

	return r.toBusinessOrder(modelOrders), nil
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

func (r *Repository) toBusinessOrder(orders []Order) []order.Order {
	response := make([]order.Order, len(orders))

	for i, item := range orders {
		response[i] = order.Order{
			ID:            item.ID,
			TransactionNo: item.TransactionNo,
			ItemID:        item.ItemID,
			ItemPrice:     item.ItemPrice,
			Qty:           item.Qty,
			Date:          item.Date,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			DeletedAt:     item.DeletedAt.Time,
		}
	}

	return response
}
