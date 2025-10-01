package models

type AttachmentRef struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement;comment:自增主键"`
	ObjUUID  []byte `gorm:"type:binary(16);not null;index;comment:附件 → file_objects:uuid"`
	BizType  string `gorm:"type:varchar(32);not null;comment:业务类型 avatar/postCover/comment …"`
	BizID    uint64 `gorm:"not null;comment:业务主键"`
	BizIndex int    `gorm:"not null;comment:业务内顺序"`
}

func (AttachmentRef) TableName() string {
	return "attachment_refs"
}
