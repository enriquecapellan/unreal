package models

import "gorm.io/gorm"


type CarModel struct {
	gorm.Model
}

type Car struct {
	CarModel
	Make        string `json:"make" validate:"required"`
	Year        int8   `json:"year" validate:"required"`
	Color       string `json:"color" validate:"required"`
	Engine      string `json:"engine"`
	Fuel        string `json:"fuel"`
	Mileage     int    `json:"mileage"`
	Price       int    `json:"price" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required"`
	Transmition string `json:"transmission" validate:"required"`
	Model string `json:"model" validate:"required"`
}
