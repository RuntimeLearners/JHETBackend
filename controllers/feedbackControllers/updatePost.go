package feedbackcontrollers

import (
	"JHETBackend/common/exception"
	feedbackservice "JHETBackend/services/feedbackService"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetFbPostStatus(c *gin.Context) {
	var req struct {
		Accepted bool `json:"accept" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.FbPostDataInvalid)
		return
	}
	// 从 URL 参数获取 id，并转换为 uint64
	parentStr := c.Param("id")
	postID, err := strconv.ParseUint(parentStr, 10, 64)
	if err != nil {
		c.Error(exception.FbPostDataInvalid)
		return
	}
	if req.Accepted {
		feedbackservice.SetFbPostStatus(postID, feedbackservice.PostStatusInProgress)
	} else {
		feedbackservice.SetFbPostStatus(postID, feedbackservice.PostStatusReviewed)
	}
}
