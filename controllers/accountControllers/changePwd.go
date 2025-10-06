package accountControllers

import  (
	"JHETBackend/common/exception"
	"JHETBackend/services/userService"
	"JHETBackend/models"
	"JHETBackend/utils"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ChangePassword(c *gin.Context) {
	// 更改密码
	var accountID uint64
	var err error
	accountID, err = GetAccountIDFromContext(c)
	if err != nil {
		c.Error(exception.ApiParamError)
		return
	}

	type PwdForm struct {
		OldPassword string `json:"oldPassword" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required"`
	}

	var pwdForm PwdForm
	err = c.ShouldBindJSON(&pwdForm)
	if err != nil {
		c.Error(exception.ApiParamError)
		return
	}
	fmt.Println(accountID, "修改密码")

	var accountInfo *models.AccountInfo
	accountInfo,err = userService.GetAccountInfoByUID(accountID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Error(exception.UsrNotExisted)
		return
	}
	if err != nil {
		fmt.Println(err)
		c.Error(exception.SysUknExc)
		return
	}

	err = userService.ChangePwd(
		accountInfo,
		pwdForm.OldPassword,
		pwdForm.NewPassword,
	)
	if err != nil {
		if errors.Is(err, exception.UsrPasswordErr) {
			c.Error(exception.UsrPasswordErr)
			return
		} else {
			c.Error(err)
			return
		}
	} else {
		utils.JsonSuccessResponse(c, "修改成功", nil)
		return
	}
}