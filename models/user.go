package models

type User struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name" form:"name"`
}
