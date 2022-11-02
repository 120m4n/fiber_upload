package route

import (
	"bootmind/pkg/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Expenses route
func Expenses(app *fiber.App, db *gorm.DB) {
    h := handler.ExpenseHandler{DB: db}

    r := app.Group("/expenses")
    r.Get("/", h.Index)
	r.Get("/:id", h.Show)
	r.Post("/", h.Create)
	r.Put("/:id", h.Update)
	r.Delete("/:id", h.Delete)
	r.Post("/:id/upload", h.Upload)
}