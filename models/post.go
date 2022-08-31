package models

type Post struct {
	ID     int          `gorm:"primaryKey" json:"id"`
	Title  string       `gorm:"not null" form:"title" json:"title"`
	Body   string       `gorm:"not null" form:"body" json:"body"`
	UserID int          `form:"user_id" json:"user_id"`
	User   UserResponse `json:"user"`
	Tags   []Tag        `gorm:"many2many:post_tags" json:"tags"`
	TagsID []int        `gorm:"-" form:"tags_id" json:"tags_id"`
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

type PostResponseWithoutUserID struct {
	ID    int    `json:"id"`
	Title string `form:"title" json:"title"`
	Body  string `form:"body" json:"body"`
}

func (PostResponseWithoutUserID) TableName() string {
	return "posts"
}

type PostResponseWithTag struct {
	ID     int    `json:"id"`
	Title  string `form:"title" json:"title"`
	Body   string `form:"body" json:"body"`
	UserID int    `json:"user_id"`
	User   User   `json:"user"`
	Tags   []Tag  `gorm:"many2many:post_tags;foreignKey:ID;joinForeignKey:PostID;References:ID;joinReferences:TagID" json:"tags"`
}

func (PostResponseWithTag) TableName() string {
	return "posts"
}
