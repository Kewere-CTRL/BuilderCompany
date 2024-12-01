package models

type Services struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Label       string `json:"label"`
	Description string `json:"description" gorm:"unique"`
	Price       int    `json:"price"`
}
