package models

type Locker struct {
	ID     int          `gorm:"primaryKey" json:"id" form:"id"`
	Code   string       `gorm:"unique, not null" json:"code" form:"code"`
	UserID int          `gorm:"not null" json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
}

type LockerResponse struct {
	ID     int    `json:"id" form:"id"`
	Code   string `json:"code" form:"code"`
	UserID int    `json:"user_id" form:"user_id"`
}

func (LockerResponse) TableName() string {
	return "lockers"
}
