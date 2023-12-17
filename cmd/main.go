package main

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Инициализация подключения к базе данных
	db := controllers.InitDatabase()
	defer db.Close()

	// Настройка маршрутов
	routes.SetupRoutes(app, db)

	// Запуск сервера
	app.Listen(":3000")
}
