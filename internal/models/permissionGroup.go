package models

// PermissionGroup 数据库模型
type PermissionGroup struct {
	ID             uint32 `gorm:"column:id;primaryKey"`
	Name           string `gorm:"column:name"`
	PermCreateUser bool   `gorm:"column:perm_create_user"`
	PermDeleteUser bool   `gorm:"column:perm_delete_user"`
	PermReadData   bool   `gorm:"column:perm_read_data"`
	PermWriteData  bool   `gorm:"column:perm_write_data"`
}

func (PermissionGroup) TableName() string {
	return "PermissionGroup"
}
