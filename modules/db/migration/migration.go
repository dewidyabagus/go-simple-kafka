package migration

import (
	"gorm.io/gorm"

	"learn/kafka/modules/db/order"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&order.Order{})
}
