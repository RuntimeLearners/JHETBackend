package dao

import (
	"JHETBackend/common/exception"
	"JHETBackend/configs/database"
	"JHETBackend/models"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FeedbackPostDB 数据库模型 用于service层出入数据打包
type FeedbackPostDTO struct {
	UserID      uint64
	Title       string
	Content     string
	Attachments []uuid.UUID // List of attachment UUIDs
	Precedence  uint8
	IsAnonymous bool
	IsPrivate   bool
	IsSpam      bool
	CreateAt    time.Time
	UpdatedAt   time.Time
	ParentID    uint64
	ReplyDepth  uint8
}

// 数据库层: 创建一条帖子 函数抽象到只要是帖子就接受
func CreateFbPost(postdata FeedbackPostDTO) error {

	var parentID *uint64
	if postdata.ParentID != 0 {
		parentID = &postdata.ParentID
	} else {
		parentID = nil
	}

	newPost := models.FeedbackPost{
		CreaterID:       postdata.UserID,
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

func CheckFbPostExist(postID uint64) bool {
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

func GetFbReplyDepth(postID uint64) uint8 {
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

// 根据帖子 ID 查询帖子主体数据及其附件 UUID 列表，封装成 FeedbackPostDTO 返回
func GetFbPostData(postID uint64) (FeedbackPostDTO, error) {
	var dto FeedbackPostDTO

	// 查 FeedbackPost 主体
	var postModel models.FeedbackPost
	if err := database.DataBase.
		First(&postModel, postID).Error; err != nil {
		return FeedbackPostDTO{}, exception.FbPostNotFount
	}

	// 查附件 UUID 列表
	attachmentUUIDs, err := getFbPostAttachmentsUUID(postID)
	if err != nil {
		log.Printf("[ERROR][FbPostDAO] 获取帖子附件 UUID 列表失败, 帖子ID: %d, 错误: %v", postID, err)
		attachmentUUIDs = []uuid.UUID{} // 出错则视作无附件
	}

	// 组装 DTO
	dto = FeedbackPostDTO{
		UserID:      postModel.CreaterID,
		Title:       postModel.Title,
		Content:     postModel.Content,
		Attachments: attachmentUUIDs,
		Precedence:  postModel.Precedence,
		IsAnonymous: postModel.IsAnonymous,
		IsPrivate:   postModel.IsPrivate,
		ParentID:    0, // 如果 ParentID 为 nil 则默认为 0
		IsSpam:      postModel.IsSpam,
		CreateAt:    postModel.CreatedAt,
		UpdatedAt:   postModel.UpdatedAt,
		ReplyDepth:  postModel.ReplyDepth,
	}
	if postModel.ParentID != nil {
		dto.ParentID = *postModel.ParentID
	}
	return dto, nil
}

func GetLatestCreatedFbIDs(maxCount int, offset int) []uint64 {
	var ids []uint64
	database.DataBase.Model(&models.FeedbackPost{}). // 指定模型
								Order("created_at DESC"). // 按创建时间倒序
								Limit(maxCount).          // 取前 N 条
								Offset(offset).           // 偏移量
								Pluck("id", &ids)         // 只查 id 字段并写入 ids 切片
	return ids
}

func GetLatestUpdatedFbIDs(maxCount int, offset int) []uint64 {
	var ids []uint64
	database.DataBase.Model(&models.FeedbackPost{}). // 指定模型
								Order("updated_at DESC"). // 按创建时间倒序
								Limit(maxCount).          // 取前 N 条
								Offset(offset).           // 偏移量
								Pluck("id", &ids)         // 只查 id 字段并写入 ids 切片
	return ids
}

func GetFbIDsWithSearchParams(searchParams models.SearchParams) []uint64 {
	// 生成函数指针组 用于统一调用
	opts := []func(*gorm.DB) *gorm.DB{
		addSearchCond("creater_id = ?", searchParams.CreaterID),
		addSearchCond("show_privates = ?", searchParams.ShowPrivates),
		addSearchCond("created_at <= ?", searchParams.CreatedBefore),
		addSearchCond("created_at >= ?", searchParams.CreatedAfter),
		addSearchCond("updated_at <= ?", searchParams.CreatedBefore),
		addSearchCond("updated_at >= ?", searchParams.CreatedAfter),
		addSearchCond("precedence >= ?", searchParams.MinPrecedence),
		addSearchCond("precedence <= ?", searchParams.MaxPrecedence),
	}

	dbReader := database.DataBase.Model(&models.FeedbackPost{}) // 指定数据库模型
	for _, filter := range opts {                               // 遍历所有筛选器
		dbReader = filter(dbReader)
	}
	// 特殊条件单独处理
	if !searchParams.ShowSpams { // 隐藏垃圾帖逻辑
		dbReader.Where("is_spam", false)
	}
	if !searchParams.ShowPrivates { // 隐藏私密帖逻辑
		dbReader.Where("is_private", false)
	}
	var ids []uint64
	dbReader.
		Order("updated_at DESC").                        // 按创建时间倒序
		Limit(searchParams.Size).                        // 取前 N 条
		Offset((searchParams.Page-1)*searchParams.Size). // 偏移量 溢出应该直接返回空
		Pluck("id", &ids)                                // 直接选 id 列
	return ids
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

func getFbPostAttachmentsUUID(postID uint64) ([]uuid.UUID, error) {
	// 查附件引用记录, 并按 BizIndex 排序
	var attachRefs []models.AttachmentRef
	if err := database.DataBase.
		Where("biz_type = ? AND biz_id = ?", "feedback_post", postID).
		Order("biz_index ASC").
		Find(&attachRefs).Error; err != nil {
		return []uuid.UUID{}, err
	}

	// 将 []byte(uuid) 转成 uuid.UUID
	attachments := make([]uuid.UUID, 0, len(attachRefs))
	for _, ref := range attachRefs {
		if len(ref.ObjUUID) == 16 { // 只有 16 位的才是合法 UUID
			attachments = append(attachments, uuid.Must(uuid.FromBytes(ref.ObjUUID)))
		} else {
			log.Printf("[WARN][FbPostDAO] 发现非法附件 UUID 长度: %d, 跳过", len(ref.ObjUUID))
		}
	}
	return attachments, nil
}

// 判断传入是否为空并自动给 GORM 查询添加条件
// 传回的是一个函数指针 务必注意使用时还需统一调用
func addSearchCond(sqlFrag string, ptr any) func(*gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if ptr == nil {
			return tx // 空条件判断
		}
		return tx.Where(sqlFrag, ptr)
	}
}
