package entity

import "time"

type Weight struct {
	ID   int        `gorm:"column:id" json:"id"`
	Date *time.Time `gorm:"column:date" json:"date"`
	Max  int        `gorm:"column:max" json:"max"`
	Min  int        `gorm:"column:min" json:"min"`
}

type WeightDto struct {
	Date *time.Time `json:"date" validate:"required"`
	Max  int        `json:"max" validate:"required,numeric"`
	Min  int        `json:"min" validate:"required,numeric"`
}

type CreateUpdateWeightRequest struct {
	Date string `json:"date" validate:"required"`
	Max  int    `json:"max" validate:"required,numeric"`
	Min  int    `json:"min" validate:"required,numeric"`
}

type DeleteWeightRequest struct {
	Date string `json:"date" validate:"required"`
}
