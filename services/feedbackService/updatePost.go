package feedbackservice

import (
	"JHETBackend/common/exception"
	"JHETBackend/dao"
)

// 用自定义类型在代码层面限制状态值
type PostStatus string

const (
	PostStatusCreated    PostStatus = "created"     // 用户已创建
	PostStatusReviewed   PostStatus = "reviewed"    // 管理员已审查
	PostStatusInProgress PostStatus = "in_progress" // 反馈正在处理中
	PostStatusClosed     PostStatus = "closed"      // 反馈帖关闭 (未解决)
	PostStatusResolved   PostStatus = "resolved"    // 反馈已解决
)

func SetFbPostStatus(fbPostID uint64, status PostStatus) error {
	if dao.SetFbPostStatus(fbPostID, string(status)) != nil {
		return exception.FbPostUpdateFailed
	}
	return nil
}

func SetFbPostSpam(fbPostID uint64, isSpam bool) error {
	if dao.SetFbPostSpam(fbPostID, isSpam) != nil {
		return exception.FbPostUpdateFailed
	}
	if isSpam { // 被标记为垃圾信息的同时关闭帖子
		if dao.SetFbPostStatus(fbPostID, string(PostStatusClosed)) != nil {
			return exception.FbPostUpdateFailed
		}
	}
	return nil
}

// 调用这个函数一定是具权者发起了审核
// 如果acknowledge设置为false, 说明审核不通过, 重新启帖并取消垃圾标记
func SetFbPostSpamChecked(fbPostID uint64, acknowledge bool) error {
	fbSpam, _, fbGetErr := dao.GetFbSpamStatus(fbPostID)
	if fbGetErr != nil {
		return exception.FbPostNotFound // 获取状态失败, 视为帖子不存在(大概率)
	}
	if !fbSpam { // 没被标记查**呢
		return exception.FbNotSpamDontChecked
	}
	if dao.SetFbPostSpamChecked(fbPostID, acknowledge) != nil {
		return exception.FbPostUpdateFailed
	}
	if !acknowledge { // 标记不被承认
		SetFbPostSpam(fbPostID, false)                // 取消标记
		SetFbPostStatus(fbPostID, PostStatusReviewed) // 启帖并设置"已阅"状态
	}
	return nil
}

func SetFbPostPrecedence(fbPostID uint64, precedence uint8) error {
	return nil
}

func RatingFbPost(fbPostID uint64, score uint8) error {
	fbStatus, fbGetErr := dao.GetFbStatus(fbPostID)
	if fbGetErr != nil {
		return exception.FbPostNotFound // 获取状态失败, 视为帖子不存在(大概率)
	}
	if fbStatus != string(PostStatusResolved) { // 对的对的 只有解决的帖子才能评分
		return exception.FbRatingUnslovedPost
	}
	if dao.RatingFbPost(fbPostID, score) != nil {
		return exception.FbPostUpdateFailed
	}
	return nil
}
