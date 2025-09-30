package models

import (
	"time"

	"gorm.io/gorm"
)

// FeedbackPost 对应表 feedback_posts
type FeedbackPost struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement;comment:主键"`
	UserID      uint64         `gorm:"not null;index;comment:发帖用户ID"`
	PostTitle   string         `gorm:"size:255;not null;comment:标题"`
	PostBody    string         `gorm:"type:text;comment:正文"`
	Precedence  uint8          `gorm:"not null;default:0;comment:优先级"`
	IsAnonymous bool           `gorm:"not null;default:0;comment:是否匿名"`
	RootID      *uint64        `gorm:"index:idx_root,priority:1;comment:根帖ID"`
	ReplyPostID *uint64        `gorm:"index:idx_reply,priority:1;comment:被回复的帖子ID"`
	CreatedAt   time.Time      `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"not null;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index;comment:软删除时间"`

	// 自引用关联（可选，按需启用）
	Root      *FeedbackPost  `gorm:"foreignKey:RootID;references:ID;constraint:OnDelete:CASCADE"`
	ReplyPost *FeedbackPost  `gorm:"foreignKey:ReplyPostID;references:ID;constraint:OnDelete:CASCADE"`
	Children  []FeedbackPost `gorm:"foreignKey:RootID;references:ID"`
	Replys    []FeedbackPost `gorm:"foreignKey:ReplyPostID;references:ID"`
}

// TableName 设置表名
func (FeedbackPost) TableName() string {
	return "feedback_posts"
}
