package models

import "gorm.io/gorm"

// Модель пользователя
type User struct {
	gorm.Model
	Username string
	Password string
	// Добавьте другие поля по необходимости
}

// Модель события
type Event struct {
	gorm.Model
	Title       string
	Description string
	Category    string
	Date        string
	// Добавьте другие поля по необходимости
}
