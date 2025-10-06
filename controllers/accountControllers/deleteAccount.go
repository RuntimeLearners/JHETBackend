package accountControllers

import (
	"JHETBackend/common/exception"
	"JHETBackend/dao"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	accountIDStr := c.Param("id")
	if accountIDStr == "" { //如果没有传入id参数, 用qurey时有用, 用path时空直接返回404
		c.Error(exception.ApiParamError)
		fmt.Println("err1")
		//utils.JsonSuccessResponse(c, "查询失败", models.AccountInfo{})
		return
	}

	var err error
	accountID, err := strconv.ParseUint(accountIDStr, 10, 64) //将传入str id转换为uint, 小转换就不写service了
	if err != nil {
		fmt.Println("err1", err)
		c.Error(exception.ApiParamError)
		return
	}
	// 调用服务层删除用户
	err = dao.DeleteAccount(accountID)
	if err != nil {
		if errors.Is(err, exception.UsrNotExisted) {
			c.Error(exception.UsrNotExisted)
		} else {
			c.Error(err)
		}
		return
	}

	//utils.JsonSuccessResponse(c, "成功删除用户", nil)
}
