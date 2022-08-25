package models

type Post struct {
	ID     int          `gorm:"primaryKey" json:"id"`
	Title  string       `gorm:"not nul" form:"title" json:"title"`
	Body   string       `gorm:"not nul" form:"body" json:"body"`
	UserID int          `form:"user_id" json:"user_id"`
	User   UserResponse `json:"user"`
}

type PostResponse struct {
	ID     int    `json:"id"`
	Title  string `form:"title" json:"title"`
	Body   string `form:"body" json:"body"`
	UserID int    `json:"user_id"`
}

func (PostResponse) TableName() string {
	return "posts"
}
