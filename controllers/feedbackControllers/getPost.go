package feedbackcontrollers

import (
	"JHETBackend/common/exception"
	"JHETBackend/models"
	feedbackservice "JHETBackend/services/feedbackService"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedbackQueryDTO struct {
	Page          int    `form:"page"`
	Number        int    `form:"number"`
	Size          int    `form:"size"`
	ID            string `form:"id"`
	Keywords      string `form:"keywords"`
	Status        string `form:"status"`
	Urgency       string `form:"urgency"`
	Category      string `form:"category"`
	ShowSpams     bool   `form:"show_spams"`
	CreatedAfter  string `form:"created_after"`
	CreatedBefore string `form:"created_before"`
	UpdatedAfter  string `form:"updated_after"`
	UpdatedBefore string `form:"updated_before"`
}

// 返回给前端的 VO
type feedbackItemVO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Category  string    `json:"category"` // 如需要可从 Status 映射
	Urgency   string    `json:"urgency"`  // 从 Precedence 反向映射
	Status    string    `json:"status"`
	IsSpam    bool      `json:"isSpam"` // 示例：Status == "spam"
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Route 入口 拆分成两个便于权限控制
func GetPublicFeedbackPosts(c *gin.Context) {
	getFeedbackPosts(c, false)
}
func GetAllFeedbackPosts(c *gin.Context) {
	getFeedbackPosts(c, true)
}

// ##### PRIVATE #####

// GetFeedbackPosts 处理 GET /api/feedback
func getFeedbackPosts(c *gin.Context, showPrivates bool) {
	var dto FeedbackQueryDTO
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.Error(exception.ApiParamError)
		return
	}

	// 构造 models.SearchParams
	// 显式留nil是为了更清晰的说明这里还没处理这些数据
	p := models.SearchParams{
		Page:          dto.Page,
		Size:          dto.Size,
		CreaterID:     nil,
		Keywords:      nil,
		Status:        &dto.Status,
		Urgency:       &dto.Urgency,
		ShowSpams:     dto.ShowSpams,
		ShowPrivates:  showPrivates,
		CreatedBefore: nil,
		CreatedAfter:  nil,
		UpdatedBefore: nil,
		UpdatedAfter:  nil,
		MinPrecedence: nil,
		MaxPrecedence: nil,
	}

	// 解析 CreaterID (类型断言)
	if dto.ID != "" {
		if id, err := strconv.ParseUint(dto.ID, 10, 64); err == nil {
			p.CreaterID = &id
		}
	}
	// 关键字
	if dto.Keywords != "" {
		p.Keywords = &dto.Keywords
	}
	// 状态
	if dto.Status != "" {
		p.Status = &dto.Status
	}
	// 紧急程度 => Precedence 范围
	switch dto.Urgency {
	case "low":
		v := uint8(1)
		p.MinPrecedence, p.MaxPrecedence = &v, &v
	case "medium":
		v := uint8(2)
		p.MinPrecedence, p.MaxPrecedence = &v, &v
	case "high":
		v := uint8(3)
		p.MinPrecedence, p.MaxPrecedence = &v, &v
	}
	// 时间范围
	// 整个小局部函数方便重用
	parseTime := func(s string) *time.Time {
		if s == "" {
			return nil
		}
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return nil
		}
		return &t
	}
	p.CreatedAfter = parseTime(dto.CreatedAfter)
	p.CreatedBefore = parseTime(dto.CreatedBefore)
	p.UpdatedAfter = parseTime(dto.UpdatedAfter)
	p.UpdatedBefore = parseTime(dto.UpdatedBefore)

	// 查数据
	list := feedbackservice.GetFbPostsWithSearchParams(p)

	// 拿到后方的数据 构造 VO 列表返给前端
	voList := make([]feedbackItemVO, 0, len(list))
	for _, post := range list {
		urgency := "medium" // 默认值
		switch post.Precedence {
		case 1:
			urgency = "low"
		case 3:
			urgency = "high"
		}
		voList = append(voList, feedbackItemVO{
			ID:        strconv.FormatUint(post.UserID, 10),
			Title:     post.Title,
			Category:  "", // TODO: 后方没做这个 XD
			Urgency:   urgency,
			Status:    "", // TODO: 后方没规定这个 XD
			IsSpam:    post.IsSpam,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	c.Set("data", voList)
}
