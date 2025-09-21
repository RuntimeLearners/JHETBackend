package permission_test

import (
	"JHETBackend/internal/common/permission"
	"log"
	"testing"
)

func Test_TryGetPermission(t *testing.T) {
	log.Print(permission.GetPermissionByGroupID(9999))
	log.Print(permission.GetPermissionByGroupID(1))
	log.Printf("%v", permission.GetAllPermissionGroups())
}

func Test_AddPermission(t *testing.T) {
	log.Printf("%v", permission.AddPermissionGroup("TEST", permission.Perm_AnswerCreate, permission.Perm_PostCreate))
	log.Printf("%v", permission.GetAllPermissionGroups())
	log.Printf("%v", permission.IsPermSatisfied(23, permission.Perm_AnswerCreate))
	log.Printf("%v", permission.IsPermSatisfied(23, permission.Perm_PostDelete))
}
