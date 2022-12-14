package models

type User struct {
	ID     int            `gorm:"primaryKey" json:"id"`
	Name   string         `gorm:"not null" json:"name" form:"name"`
	Locker LockerResponse `json:"locker"`
	Post   []PostResponse `json:"post"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}

type UserWithoutPost struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Locker LockerResponse `gorm:"foreignKey:ID;joinForeignKey:UserID" json:"locker"`
}

func (UserWithoutPost) TableName() string {
	return "users"
}
