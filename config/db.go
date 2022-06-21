package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"learn/kafka/modules/db/migration"
)

func NewDatabase() (db *gorm.DB, err error) {
	config := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		Database.Host, Database.User, Database.Password, Database.Schema, Database.Port, Database.TimeZone,
	)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migration.AutoMigrate(db)

	return db, nil
}
