package models

// PermissionGroup 数据库模型
type PermissionGroup struct {
	// ID不应手动输入，交给数据库自增管理
	ID uint32 `gorm:"column:PGID;primaryKey"`
	// 权限组名称
	Name  string `gorm:"column:PGName"`
	Perm1 bool   `gorm:"column:perm1"`
	Perm2 bool   `gorm:"column:perm2"`
	//PermN bool   `gorm:"column:notfoundperm"` //debug
}

func (PermissionGroup) TableName() string {
	return "PermissionGroups"
}
