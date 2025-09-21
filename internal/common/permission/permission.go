// permission.go
package permission

import (
	//"JHETBackend/internal/common/basics"
	"JHETBackend/internal/configs/database"
	"JHETBackend/internal/models"
	"fmt"
	"log"
	"sync"

	"github.com/bits-and-blooms/bitset"
)

// 权限枚举，以 Perm_ 开头，与数据库列名一致
const (
	_                 = iota // 0 空出来
	Perm_PostCreate          // 1
	Perm_PostDelete          // 2
	Perm_RepairCreate        // 3
	Perm_AnswerCreate        // 4
	// 往下继续加...
)

type PermissionID uint32

var permissionGroups = map[uint32]models.PermissionGroup{} // permGroupID -> 权限组 对应表

// 加载数据库中的权限组权限表
func loadFromDB() {
	var tmpPermG []models.PermissionGroup
	if err := database.DataBase.Model(&models.PermissionGroup{}).Find(&tmpPermG).Error; err != nil {
		log.Panic("[FATAL][PERM] 无法读取权限列表")
		return
	}
	// 将数据输入map中，索引使用gpid
	for _, g := range tmpPermG {
		if err := g.Permissions.UnmarshalBinary(g.PermissionData); err != nil {
			log.Panicf("[FATAL][PERM] 权限数据不符合规则 错误: %v", err)
			return
		}
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

func IsPermSatisfied(permGroupId uint32, needed ...uint32) bool {
	tocheck, err := GetPermissionByGroupID(permGroupId)
	if err != nil {
		log.Print("[ERROR][PERM] 该权限组不存在 视为无权")
		return false
	}
	for _, perm := range needed {
		if !tocheck.Permissions.Test(uint(perm)) {
			return false
		}
	}
	return true
}

func AddPermissionGroup(name string, permissions ...uint32) error {
	var newPG models.PermissionGroup
	newPG.Name = name
	newPG.Permissions = *bitset.New(255)
	for _, perm := range permissions {
		newPG.Permissions.Set(uint(perm))
	}
	pgdata, err := newPG.Permissions.MarshalBinary()
	if err != nil {
		log.Printf("[ERROR][PERM] 无法转换权限位图到字节型 保存权限表失败 错误: %v", err)
	}
	newPG.PermissionData = pgdata
	dbnp := database.DataBase.Create(&newPG)
	if dbnp.Error != nil {
		return dbnp.Error
	}
	loadFromDB() // 重新从数据库载入权限表
	return nil
}
