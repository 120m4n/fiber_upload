package entity

import (

	"gorm.io/gorm"
)

// Expense type
type Expense struct {
    gorm.Model
    ID         uint32     `json:"id"`
    Title      string    `json:"title"`
    Total      float64   `json:"total"`
    Attachment string    `json:"attachment"`
}