package feedbackservice

import (
	"JHETBackend/common/exception"
	"JHETBackend/dao"
)

type FeedbackPost struct {
	UserID      uint64
	Title       string
	Content     string
	Precedence  uint8
	IsAnonymous bool
	IsPrivate   bool
}

type FeedbackReplyPost struct {
	FeedbackPost
	ParentID uint64
}

func CreateFeedbackPost(postdata FeedbackPost) error {
	daoReq := dao.FeedbackPostDAO{
		UserID:      postdata.UserID,
		Title:       postdata.Title,
		Content:     postdata.Content,
		Precedence:  postdata.Precedence,
		IsAnonymous: postdata.IsAnonymous,
		IsPrivate:   postdata.IsPrivate,
		ParentID:    0,
		ReplyDepth:  0,
	}
	if err := dao.CreateFeedbackPost(daoReq); err != nil {
		return err
	}
	return nil
}

func ReplyFeedbackPost(replayPostdata FeedbackReplyPost) error {
	if !dao.CheckFeedbackPostExist(replayPostdata.ParentID) {
		return exception.FbReplyPostNotFound
	}
	parentReplyDepth := dao.GetFeedbackReplyDepth(replayPostdata.ParentID)
	if parentReplyDepth >= 1 {
		return exception.FbReplyNestTooDeep
	}
	daoReq := dao.FeedbackPostDAO{
		UserID:      replayPostdata.UserID,
		Title:       replayPostdata.Title,
		Content:     replayPostdata.Content,
		Precedence:  replayPostdata.Precedence,
		IsAnonymous: replayPostdata.IsAnonymous,
		IsPrivate:   replayPostdata.IsPrivate,
		ParentID:    replayPostdata.ParentID,
		ReplyDepth:  parentReplyDepth + 1,
	}
	if err := dao.CreateFeedbackPost(daoReq); err != nil {
		return err
	}
	return nil
}
