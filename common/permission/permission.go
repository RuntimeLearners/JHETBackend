// permission.go
package permission

import (
	//"JHETBackend/common/basics"
	"JHETBackend/common/exception"
	"JHETBackend/dao"
	"JHETBackend/models"
	"fmt"
	"log"
	"sync"

	"github.com/bits-and-blooms/bitset"
)

// 权限枚举，以 Perm_ 开头，与数据库列名一致
type PermissionID uint32

const (
	_                 PermissionID = 0
	Perm_ForTestOnly1 PermissionID = 255
	Perm_ForTestOnly2 PermissionID = 254

	// 登录
	Perm_Login PermissionID = 001 // 登录

	// 个人信息
	Perm_GetProfile    PermissionID = 010 // 获取个人信息
	Perm_UpdateProfile PermissionID = 011 // 修改个人信息
	Perm_UpdateAvatar  PermissionID = 012 // 更改头像

	//附件
	Perm_UploadImage PermissionID = 021 // 上传图片

	// 问题反馈
	Perm_SubmitFeedback   PermissionID = 031 // 提交反馈
	Perm_ViewFeedback     PermissionID = 032 // 查看反馈（反馈详情）
	Perm_AcceptOrder      PermissionID = 033 // 接单
	Perm_MarkAsSpam       PermissionID = 034 // 标记垃圾
	Perm_ReplyFeedback    PermissionID = 035 // 回复反馈
	Perm_RateFeedback     PermissionID = 036 // 评价反馈
	Perm_QueryFeedbackLog PermissionID = 037 // 查询反馈记录

	// 管理面板
	// 超管 - 用户
	Perm_GetUserList        PermissionID = 100 // 获取用户列表
	Perm_AddUser            PermissionID = 101 // 新增用户
	Perm_EditUser           PermissionID = 102 // 编辑用户
	Perm_DeleteUser         PermissionID = 103 // 删除用户
	Perm_GetUserPermission  PermissionID = 104 // 获取用户权限
	Perm_EditUserPermission PermissionID = 105 // 编辑用户权限
	Perm_GetAnyProfile      PermissionID = 106 // 获取个人的所有信息

	// 超管 - 垃圾审核
	Perm_GetPendingSpam PermissionID = 110 // 获取待审核垃圾
	Perm_ReviewSpam     PermissionID = 111 // 审核垃圾

	// 预设
	Perm_EditPreset   PermissionID = 200 // 编辑预设
	Perm_ViewPreset   PermissionID = 210 // 查看预设
	Perm_AddPreset    PermissionID = 220 // 新增预设
	Perm_DeletePreset PermissionID = 230 // 删除预设

	// 往下继续加...
)

var permissionGroups = map[uint32]models.PermissionGroup{} // permGroupID -> 权限组 对应表

// 加载数据库中的权限组权限表
func loadFromDB() {
	var err error
	if permissionGroups, err = dao.GetAllPermissionGroup(); err != nil {
		panic(exception.SysCannotLoadPermGroup)
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

func IsPermSatisfied(permGroupId uint32, needed ...PermissionID) bool {
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

func AddPermissionGroup(name string, permissions ...PermissionID) error {
	var newPG models.PermissionGroup
	newPG.Name = name
	newPG.Permissions = *bitset.New(255)
	for _, perm := range permissions {
		newPG.Permissions.Set(uint(perm))
	}
	dao.AddPermissionGroup(newPG)
	loadFromDB() // 重新从数据库载入权限表
	return nil
}
