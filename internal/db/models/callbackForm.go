package models

type CallbackForm struct {
	UserID          uint   `json:"userID"`
	TelephoneNumber string `json:"telephoneNumber"`
	User            User   `gorm:"foreignKey:UserID;references:ID"`
}
