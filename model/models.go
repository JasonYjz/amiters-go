package model

import "time"

type Model struct {
	Id uint `gorm:"primaryKey;autoIncrement" json:"id"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Model
	Name string `gorm:"type:varchar(50)" json:"name"`
	Mail string `gorm:"type:varchar(50)" json:"mail"`
	Wxid string `gorm:"type:varchar(12)" json:"wxid"`
	Img string `gorm:"type:varchar(32)" json:"img"`
	QQ string `gorm:"type:varchar(12)" json:"qq"`
	Signature string `gorm:"type:varchar(100)" json:"signature"`
}

type Post struct {
	Model
	UserId uint `gorm:"index:idx_post_user_id" json:"userId"`
	SectionId uint `gorm:"index:idx_post_section_id" json:"sectionId"`
	Title string `gorm:"type:varchar(50)" json:"title"`
	Content string `gorm:"type:type:longtext;not null;" json:"content"`
	ViewCount uint `gorm:"type:int(10) unsigned" json:"viewCount"`
	LikedCount uint `gorm:"type:int(10) unsigned" json:"likedCount"`
	CreateTime time.Time `gorm:"type:datetime" json:"createTime"`
	UpdateTime time.Time `gorm:"type:datetime" json:"updateTime"`
}

type Section struct {
	Model
	Name string `gorm:"type:varchar(50)" json:"name"`
	Description string `gorm:"type:varchar(100)" json:"description"`
}

type Tag struct {
	Model
	Name string `gorm:"type:varchar(20)" json:"name"`
}


type PostTag struct {
	PostId uint `gorm:"type:int(10) unsigned" json:"postId"`
	TagId uint `gorm:"type:int(10) unsigned" json:"tagId"`
}