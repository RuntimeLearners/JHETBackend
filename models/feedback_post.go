package models

import (
	"time"

	"gorm.io/gorm"
)

// FeedbackPost 对应表 feedback_posts
type FeedbackPost struct {
	ID              uint64 `gorm:"primaryKey;autoIncrement;comment:唯一帖子ID"`
	CreaterID       uint64 `gorm:"not null;index;comment:发帖用户ID"`
	Title           string `gorm:"size:255;not null;comment:标题"`
	Content         string `gorm:"type:text;comment:正文"`
	Precedence      uint8  `gorm:"not null;default:0;comment:优先级"`
	HaveAttachments bool   `gorm:"not null;default:0;comment:是否有附件"`
	IsAnonymous     bool   `gorm:"not null;default:0;comment:是否匿名"`
	IsPrivate       bool   `gorm:"not null;default:0;comment:是否公开"`
	IsClosed        bool   `gorm:"not null;default:0;comment:是否已关闭"`
	IsSpam          bool   `gorm:"not null;default:0;comment:是否为垃圾帖"`
	// 下面使用指针以区分 NULL 和 0
	// ParentID 指回复的目标帖子
	ParentID   *uint64        `gorm:"index:idx_parent,priority:1;comment:根帖ID"`
	ReplyDepth uint8          `gorm:"not null;comment:回复嵌套深度"`
	CreatedAt  time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt  time.Time      `gorm:"not null;comment:更新时间"`
	DeletedAt  gorm.DeletedAt `gorm:"index;comment:软删除时间"`

	// 自引用关联
	Replys []FeedbackPost `gorm:"foreignKey:ParentID;references:ID"`
}

// TableName 设置表名
func (FeedbackPost) TableName() string {
	return "feedback_posts"
}
