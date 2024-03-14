package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Make string `json:"make" validate:"required"`
}