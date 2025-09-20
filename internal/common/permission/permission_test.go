package permission_test

import (
	"JHETBackend/internal/common/permission"
	"JHETBackend/internal/models"
	"log"
	"testing"
)

func Test_TryGetPermission(t *testing.T) {
	log.Print(permission.GetPermissionByGroupID(9999))
	log.Print(permission.GetPermissionByGroupID(1))
	log.Printf("%v", permission.GetAllPermissionGroups())
}

func Test_AddPermission(t *testing.T) {
	if err := permission.AddPermissionGroup(&models.PermissionGroup{
		Name:  "Test_UserGroup4",
		Perm1: false,
		Perm2: true,
	}); err != nil {
		t.Errorf("please check dbwerror: %v", err)
	}
}
