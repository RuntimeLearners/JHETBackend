package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Post struct {
	ID          uint64         `json:"uid" gorm:"column:uid"` //默认为主键
	Title       string         `json:"title" gorm:"column:title"`
	Content     string         `json:"content" gorm:"column:content"`
	Category    string         `json:"category" gorm:"column:category"`
	Urgency     string         `json:"urgency" gorm:"index"`
	Status      string         `json:"status" gorm:"index"` //处理状态
	IsAnonymous bool           `json:"isAnonymous" gorm:"index"`
	Iamges      datatypes.JSON //储存的是JSON，返回格式为：{{"id":1,"url":"http://xxxx"},{图片2}} 储存格式为：{{"id":1,"url":"文件名"},{图片2}}
	Creator     datatypes.JSON //{"id":1,"name":"昵称","avatar":"http://xxxx"}
	Handler     datatypes.JSON //{"id":1,"name":"姓名","phoneNumber":"","handleAt":"time.Time"}
	Replies     datatypes.JSON //{"id":1,"content":"xxxx","creator":{"content":1,"images":{和上面的格式一样},"createdAt":"time.Time"}
	Evaluation  datatypes.JSON //{"score":5,"comment":"xxxx","createdAt":"time.Time"}
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"` //估计用不到
}
