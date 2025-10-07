package permission_test

import (
	"JHETBackend/common/permission"
	"log"
	"testing"
)

func Test_TryGetPermission(t *testing.T) {
	log.Print(permission.GetPermissionByGroupID(9999))
	log.Print(permission.GetPermissionByGroupID(1))
	log.Printf("%v", permission.GetAllPermissionGroups())
}

func Test_AddSUPermission(t *testing.T) {
	log.Printf("%v", permission.AddPermissionGroup("SU", permission.Perm_All_SUONLY))
}

func Test_AddUSERPermission(t *testing.T) {
	log.Printf("%v", permission.AddPermissionGroup("USER",
		permission.Perm_Login,
		permission.Perm_SubmitFeedback,
		permission.Perm_UpdateAvatar,
		permission.Perm_UpdateProfile,
		permission.Perm_UpdatePwd,
		permission.Perm_UploadImage))
}

func Test_AddADMINPermission(t *testing.T) {
	log.Printf("%v", permission.AddPermissionGroup("ADMIN",
		permission.Perm_Login,
		permission.Perm_SubmitFeedback,
		permission.Perm_UpdateAvatar,
		permission.Perm_UpdateProfile,
		permission.Perm_UpdatePwd,
		permission.Perm_UploadImage,
		permission.Perm_AcceptOrder,
		permission.Perm_GetAnyProfile,
		permission.Perm_MarkAsSpam,
		permission.Perm_ReplyFeedback,
		permission.Perm_ViewFeedback,
		permission.Perm_ViewPreset))
}

func Test_AddUserPG(t *testing.T) {
	permission.AddPermissionGroup("USER", permission.Perm_GetProfile,
		permission.Perm_Login,
		permission.Perm_UploadImage,
		permission.Perm_UpdateAvatar,
		permission.Perm_SubmitFeedback)
	log.Printf("%v", permission.GetAllPermissionGroups())
}
