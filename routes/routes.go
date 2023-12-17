package routes

import (
	"github.com/gofiber/fiber/v2"
	"event-planner/controllers"
	"gorm.io/gorm"
)

// SetupRoutes регистрирует все маршруты приложения
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	SetupUserRoutes(app, db)
	SetupEventRoutes(app, db)
}

// SetupUserRoutes регистрирует маршруты, связанные с пользователями
func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/users/register", controllers.RegisterUser)
	app.Post("/users/login", controllers.LoginUser)
	app.Post("/users/logout", controllers.LogoutUser)
	app.Get("/users/profile/:userId", controllers.UserProfile)
}

// SetupEventRoutes регистрирует маршруты, связанные с событиями
func SetupEventRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/events/create", controllers.CreateEvent)
	app.Get("/events/view/:eventId", controllers.ViewEvent)
	app.Put("/events/update/:eventId", controllers.UpdateEvent)
	app.Delete("/events/delete/:eventId", controllers.DeleteEvent)
	app.Get("/events/all", controllers.ListAllEvents)
	app.Get("/events/category/:categoryName", controllers.ListEventsByCategory)
	app.Get("/events/user/:userId", controllers.ListUserEvents)
	app.Get("/events/search", controllers.SearchEvents)
}
