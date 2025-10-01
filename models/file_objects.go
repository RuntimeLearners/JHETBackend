package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileObject struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;comment:自增主键"`
	UUID      uuid.UUID      `gorm:"type:binary(16);uniqueIndex;not null;comment:业务UUID"`
	Sha256    [32]byte       `gorm:"type:binary(32);uniqueIndex:uk_sha256;not null;comment:文件SHA256"`
	FileSize  int64          `gorm:"not null;comment:文件字节数"`
	FileName  string         `gorm:"type:varchar(64);not null;comment:原始文件名"`
	FileType  string         `gorm:"type:varchar(64);not null;comment:MIME/文件类型"`
	CreatedAt time.Time      `gorm:"not null;comment:创建时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除时间"`
}

func (FileObject) TableName() string {
	return "file_objects"
}
