package models

type Tag struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type TagResponseWithPost struct {
	ID   int                         `gorm:"primaryKey" json:"id"`
	Name string                      `gorm:"not null" json:"name"`
	Post []PostResponseWithoutUserID `gorm:"many2many:post_tags;ForeignKey:ID;joinForeignKey:TagID;References:ID;joinReferences:PostID" json:"post"`
}

func (TagResponseWithPost) TableName() string {
	return "tags"
}
