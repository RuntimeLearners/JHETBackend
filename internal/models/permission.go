package models

import "github.com/bits-and-blooms/bitset"

// PermissionGroup 数据库模型
type PermissionGroup struct {
	// ID不应手动输入，交给数据库自增管理
	ID uint32 `gorm:"column:PGID;primaryKey"`
	// 权限组名称
	Name           string        `gorm:"column:PGName"`
	PermissionData []byte        `gorm:"column:PermissionData"`
	Permissions    bitset.BitSet `gorm:"-"`
}

func (PermissionGroup) TableName() string {
	return "PermissionGroups"
}
