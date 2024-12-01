package models

import "time"

type Project struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	Label             string    `json:"label"`
	Description       string    `json:"description" gorm:"unique"`
	Date              time.Time `json:"date"`
	EngineeringSystem string    `json:"engineeringSystem"`
}
