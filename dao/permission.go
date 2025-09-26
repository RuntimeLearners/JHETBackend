package dao

import (
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"log"

	"github.com/bits-and-blooms/bitset"
)

// PermissionGroup 数据库模型 仅供dao内部使用
type permissionGroupDB struct {
	// ID不应手动输入，交给数据库自增管理
	ID uint32 `gorm:"column:PGID;primaryKey"`
	// 权限组名称
	Name string `gorm:"column:PGName"`
	// 权限组数据 即权限表的数据库内存储形式
	PermissionData []byte `gorm:"column:PermissionData"`
	// 权限组权限表表
	Permissions bitset.BitSet `gorm:"-"`
}

func (permissionGroupDB) TableName() string {
	return "permission_groups" // 数据库表名：由于名称后缀有DB所以需要手动指定
}

func GetAllPermissionGroup() (result map[uint32]models.PermissionGroup, err error) {
	var tmpPermG []permissionGroupDB
	if err := database.DataBase.Model(&permissionGroupDB{}).Find(&tmpPermG).Error; err != nil {
		log.Printf("[ERROR][PERM] 无法读取权限列表 错误: %v", err)
		return map[uint32]models.PermissionGroup{}, err
	}
	// 将数据输入map中，索引使用gpid
	result = make(map[uint32]models.PermissionGroup)
	for _, g := range tmpPermG {
		if err := g.Permissions.UnmarshalBinary(g.PermissionData); err != nil {
			log.Printf("[ERROR][PERM] 权限数据不符合规则 错误: %v", err)
			return map[uint32]models.PermissionGroup{}, err
		}
		// real permgroup unit
		rpgu := models.PermissionGroup{
			Name:        g.Name,
			Permissions: g.Permissions,
		}
		result[g.ID] = rpgu
	}
	return result, nil
}

func AddPermissionGroup(target models.PermissionGroup) error {
	//database permission group unit
	dbpgu := permissionGroupDB{
		Name:        target.Name,
		Permissions: target.Permissions,
	}
	// change bitset to []bytes and input to DB
	pgdata, err := dbpgu.Permissions.MarshalBinary()
	if err != nil {
		log.Printf("[ERROR][DAO-PERM] 无法转换权限位图到字节型 保存权限表失败 错误: %v", err)
		return err
	}
	dbpgu.PermissionData = pgdata
	dbnp := database.DataBase.Create(&dbpgu)
	if dbnp.Error != nil {
		return dbnp.Error
	}
	return nil
}
