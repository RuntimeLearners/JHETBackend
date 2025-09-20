// permission.go
package permission

import (
	"JHETBackend/internal/configs/database"
	"JHETBackend/internal/models"
	"fmt"
	"log"
	"sync"
)

// 权限枚举，以 Perm_ 开头，与数据库列名一致

type PermissionID uint32

var permissionGroups = map[uint32]models.PermissionGroup{} // permGroupID -> 权限组 对应表

// 加载数据库中的权限组权限表
func loadFromDB() {
	var tmpPermG []models.PermissionGroup
	if err := database.DataBase.Model(&models.PermissionGroup{}).Find(&tmpPermG).Error; err != nil {
		log.Panic("[FATAL][PERM] 无法读取权限列表")
		return
	}
	// 写入 map，key 用 ID
	for _, g := range tmpPermG {
		permissionGroups[g.ID] = g
	}
}

var loadDBOnce sync.Once

func GetPermissionByGroupID(permGroupId uint32) (models.PermissionGroup, error) {
	loadDBOnce.Do(loadFromDB) // 懒加载：从数据库获取权限组权限表

	premGroupResult, ok := permissionGroups[permGroupId]
	if !ok {
		log.Print("[ERROR][PERM] 尝试获取一个不存在的权限组权限")
		return models.PermissionGroup{}, fmt.Errorf("permission group with ID %d not found", permGroupId)
	}
	return premGroupResult, nil
}

func GetAllPermissionGroups() *map[uint32]models.PermissionGroup {
	loadDBOnce.Do(loadFromDB) // 懒加载：从数据库获取权限组权限表
	return &permissionGroups
}

func AddPermissionGroup(permissionGroup *models.PermissionGroup) error {
	return database.DataBase.Create(permissionGroup).Error
}
