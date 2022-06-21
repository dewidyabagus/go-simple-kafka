package migration

import (
	"learn/kafka/modules/db/order"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&order.Order{})
}
