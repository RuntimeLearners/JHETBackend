package feedbackservice

import (
	"JHETBackend/common/exception"
	"JHETBackend/dao"

	"github.com/google/uuid"
)

type FeedbackBasics struct {
	UserID      uint64
	Title       string
	Content     string
	Attachments []uuid.UUID // List of attachment UUIDs
	IsAnonymous bool
}

type FeedbackPost struct {
	FeedbackBasics
	Precedence uint8
	IsPrivate  bool
}

type FeedbackReplyPost struct {
	FeedbackBasics
	ParentID uint64
}

func CreateFeedbackPost(postdata FeedbackPost) error {
	daoReq := dao.FeedbackPostDTO{
		UserID:      postdata.UserID,
		Title:       postdata.Title,
		Content:     postdata.Content,
		Attachments: postdata.Attachments,
		Precedence:  postdata.Precedence,
		IsAnonymous: postdata.IsAnonymous,
		IsPrivate:   postdata.IsPrivate,
		ParentID:    0,
		ReplyDepth:  0,
	}
	if err := dao.CreateFbPost(daoReq); err != nil {
		return err
	}
	return nil
}

func ReplyFeedbackPost(replayPostdata FeedbackReplyPost) error {
	if !dao.CheckFbPostExist(replayPostdata.ParentID) {
		return exception.FbReplyPostNotFound
	}
	parentReplyDepth := dao.GetFbReplyDepth(replayPostdata.ParentID)
	if parentReplyDepth >= 1 {
		return exception.FbReplyNestTooDeep
	}
	daoReq := dao.FeedbackPostDTO{
		UserID:      replayPostdata.UserID,
		Title:       replayPostdata.Title,
		Content:     replayPostdata.Content,
		Attachments: replayPostdata.Attachments,
		Precedence:  0, // 回复帖没有优先级
		IsAnonymous: replayPostdata.IsAnonymous,
		IsPrivate:   false, // 回复帖当然是公开的
		ParentID:    replayPostdata.ParentID,
		ReplyDepth:  parentReplyDepth + 1,
	}
	if err := dao.CreateFbPost(daoReq); err != nil {
		return err
	}
	return nil
}
