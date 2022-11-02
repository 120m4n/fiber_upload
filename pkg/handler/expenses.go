package handler

import (
	"bootmind/pkg/entity"
	"bootmind/pkg/validation"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ExpenseHandler type
type ExpenseHandler struct {
    DB *gorm.DB
}

// Index to list all expenses
func (h ExpenseHandler) Index(ctx *fiber.Ctx) error {
    // expenses := []entity.Expense{
    //     {
    //         ID:         "1",
    //         Title:      "Lunch at MyFood",
    //         Total:      14.95,
    //         Attachment: "photo.jpg",
    //     },
    // }
    var expenses []entity.Expense
    h.DB.Find(&expenses)
    return ctx.JSON(fiber.Map{"data": expenses})
}
// Show to show a single expense
func (h ExpenseHandler) Show(ctx *fiber.Ctx) error {
    id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
    }

    var expense entity.Expense
    expense.ID = uint32(id)
    h.DB.First(&expense)
    return ctx.JSON(fiber.Map{"data": expense})


    
}    
// Create to create a new expense
func (h ExpenseHandler) Create(ctx *fiber.Ctx) error {
    var expense entity.Expense
    if err := ctx.BodyParser(&expense); err != nil {
        return err
    }

    //validate the data
    if validate := validation.CreateOrUpdate(expense); len(validate) > 0 {
        return ctx.Status(422).JSON(fiber.Map{"message": validate})
    }

    h.DB.Create(&expense)
    return ctx.JSON(fiber.Map{"message": "Expense successfully registered"})
}

// Update to update a expense
func (h ExpenseHandler) Update(ctx *fiber.Ctx) error {
    expense := new(entity.Expense)
    if err := ctx.BodyParser(expense); err != nil {
        return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
    }

    id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

    if err != nil {
        return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your expense"}})
    }

    if validate := validation.CreateOrUpdate(*expense); len(validate) > 0 {
      return ctx.Status(422).JSON(fiber.Map{"errors": validate})
    }

    var expenseDB entity.Expense
    expenseDB.ID = uint32(id)
    h.DB.First(&expenseDB)
    h.DB.Model(&expenseDB).Updates(map[string]interface{}{
        "title": expense.Title,
        "total": expense.Total,
    })

    return ctx.JSON(fiber.Map{"message": "Expense successfully updated"})
}

// Delete to delete a expense
func (h ExpenseHandler) Delete(ctx *fiber.Ctx) error {
    id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
    if err != nil {
        return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
    }

    var expense entity.Expense
    expense.ID = uint32(id)
    h.DB.Delete(&expense)
    return ctx.JSON(fiber.Map{"message": "Expense successfully deleted"})
}

// Upload an attachment
func (h ExpenseHandler) Upload(ctx *fiber.Ctx) error {
    id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
    if err != nil {
        return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"Dont be a fool"}})
    }

    file, err := ctx.FormFile("attachment")

    if err != nil {
        return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
    }

    // save file to root directory
    if err := ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename)); err != nil {
        return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"Error saving file"}})
    }


    var expense entity.Expense
    expense.ID = uint32(id)
    h.DB.First(&expense)
    h.DB.Model(&expense).Update("attachment", file.Filename)

    return ctx.JSON(fiber.Map{"message": "Attachment successfully uploaded"})
}


