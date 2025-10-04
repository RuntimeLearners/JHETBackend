package feedbackcontrollers

import (
	"JHETBackend/common/exception"
	"JHETBackend/controllers/accountControllers"
	feedbackservice "JHETBackend/services/feedbackService"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 前端传入的 JSON 结构
type createFeedbackReq struct {
	Title           string   `json:"title" binding:"required,min=1,max=120"`
	Urgency         string   `json:"urgency" binding:"required,oneof=low medium high"`
	Content         string   `json:"content" binding:"required,min=1,max=5000"`
	IsAnonymous     bool     `json:"isAnonymous"`
	IsPrivate       bool     `json:"isPrivate"`
	AttachmentUUIDs []string `json:"attachment_uuids"`
}

// 创建反馈帖控制解析器
func CreateFeedbackPost(c *gin.Context) {
	var req createFeedbackReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.FbPostDataInvalid)
		return
	}

	//解析优先级 (前端似乎喜欢用字符串)
	var precedence uint8
	switch req.Urgency {
	case "low":
		precedence = 1
	case "medium":
		precedence = 2
	case "high":
		precedence = 3
	default:
		c.Error(exception.FbPostDataInvalid)
		return
	}

	// 解析附件
	attachments := make([]uuid.UUID, 0, len(req.AttachmentUUIDs))
	for _, s := range req.AttachmentUUIDs {
		uid, err := uuid.Parse(s)
		if err != nil {
			c.Error(exception.FbPostAttachmentInvalid)
			return
		}
		attachments = append(attachments, uid)
	}

	// 取当前用户 ID
	accountID, err := accountControllers.GetAccountIDFromContext(c)
	if err != nil {
		c.Error(err)
		return
	}

	// 组装业务对象
	postData := feedbackservice.FeedbackPost{
		FeedbackBasics: feedbackservice.FeedbackBasics{
			UserID:      accountID,
			Title:       req.Title,
			Content:     req.Content,
			Attachments: attachments,
			IsAnonymous: req.IsAnonymous,
		},
		Precedence: precedence,
		IsPrivate:  req.IsPrivate,
	}

	// 给到 service 层
	if err := feedbackservice.CreateFeedbackPost(postData); err != nil {
		c.Error(exception.ApiFeedbackNotCreated)
		return
	}
}
