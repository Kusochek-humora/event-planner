package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Инициализация подключения к базе данных
func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("event_planner.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автомиграция моделей (создание таблиц в базе данных)
	db.AutoMigrate(&User{}, &Event{})
	return db, nil
}
