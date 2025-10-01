package dao

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"log"

	"github.com/google/uuid"
)

// FeedbackPostDB 数据库模型 用于service层传入数据打包
type FeedbackPostDAO struct {
	UserID      uint64
	Title       string
	Content     string
	Attachments []uuid.UUID // List of attachment UUIDs
	Precedence  uint8
	IsAnonymous bool
	IsPrivate   bool
	ParentID    uint64
	ReplyDepth  uint8
}

// 数据库层: 创建一条帖子 函数抽象到只要是帖子就接受
func CreateFeedbackPost(postdata FeedbackPostDAO) error {

	var parentID *uint64
	if postdata.ParentID != 0 {
		parentID = &postdata.ParentID
	} else {
		parentID = nil
	}

	newPost := models.FeedbackPost{
		UserID:          postdata.UserID,
		Title:           postdata.Title,
		Content:         postdata.Content,
		Precedence:      postdata.Precedence,
		HaveAttachments: (len(postdata.Attachments) > 0),
		IsAnonymous:     postdata.IsAnonymous,
		IsPrivate:       postdata.IsPrivate,
		IsClosed:        false, // 新帖默认不关闭
		ParentID:        parentID,
		ReplyDepth:      postdata.ReplyDepth,
	}

	// TODO: 入库和注册附件放在事务中会更合理 但无伤大雅
	// 帖子入库
	if err := database.DataBase.Create(&newPost).Error; err != nil {
		return exception.ApiFeedbackNotCreated
	}

	// 在引用表注册附件 先入库, 才能拿到正确的 id
	if newPost.HaveAttachments {
		if err := regPostAttachment(newPost.ID, postdata.Attachments); err != nil {
			log.Printf("[ERROR][FeedbackPostDAO] 无法注册附件 错误: %v", err)
			return exception.ApiFeedbackNotCreated
		}
	}

	return nil
}

func CheckFeedbackPostExist(postID uint64) bool {
	var cnt int64
	if err := database.DataBase.Model(&models.FeedbackPost{}).
		Where("id = ?", postID).
		Count(&cnt).Error; err != nil {
		return false // 数据库错误视作找不到外键
	}
	if cnt == 0 {
		return false // 找不到外键
	}
	if cnt > 1 {
		panic("[!][FATAL] 帖子ID在数据库发生重复, 请检查数据库完整性")
	}
	return true
}

func GetFeedbackReplyDepth(postID uint64) uint8 {
	var depth uint8
	if err := database.DataBase.Model(&models.FeedbackPost{}).
		Where("id = ?", postID).
		Select("reply_depth").
		Scan(&depth).Error; err != nil {
		// 记录不存在或其他错误，统一返回 0
		return 0
	}
	return depth
}

// ##### PRIVATE #####

func regPostAttachment(postID uint64, obj_uuids []uuid.UUID) error {
	for index, value := range obj_uuids {
		valUUID, err := value.MarshalBinary() // 转换 UUID 到 []byte
		if err != nil {
			return exception.FbPostAttachmentInvalid
		}
		newAttachmentRef := models.AttachmentRef{
			ObjUUID:  valUUID,
			BizType:  "feedback_post", // 这个 dao 只处理这个业务
			BizID:    postID,
			BizIndex: index,
		}
		if err := database.DataBase.Create(&newAttachmentRef).Error; err != nil {
			return err
		}
	}
	return nil
}
