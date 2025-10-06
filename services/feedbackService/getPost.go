package feedbackservice

import (
	"JHETBackend/dao"
	"JHETBackend/models"
)

// struct __DTO

// GetLatestFbPs (resCount int) -> [] struct

// 获取最近创建的帖子
func GetLatestFeedbackPosts(maxCount int, offset int) []FeedbackPost {
	result := []FeedbackPost{}
	fbPostIndexs := dao.GetLatestCreatedFbIDs(maxCount, offset)
	for _, fbPostIndex := range fbPostIndexs {
		if fbPostDataDTO, err := dao.GetFbPostData(fbPostIndex); err == nil {
			fbPost := FeedbackPost{
				FeedbackBasics: FeedbackBasics{
					UserID:      fbPostDataDTO.UserID,
					Title:       fbPostDataDTO.Title,
					Content:     fbPostDataDTO.Content,
					Attachments: fbPostDataDTO.Attachments,
					IsAnonymous: fbPostDataDTO.IsAnonymous,
				},
				Precedence: fbPostDataDTO.Precedence,
				IsPrivate:  fbPostDataDTO.IsPrivate,
			}
			result = append(result, fbPost)
		}
	}
	return result
}

func GetFbPostsWithSearchParams(params models.SearchParams) []FeedbackPost {
	result := []FeedbackPost{}
	fbPostIndexs := dao.GetFbIDsWithSearchParams(params)
	for _, fbPostIndex := range fbPostIndexs {
		if fbPostDataDTO, err := dao.GetFbPostData(fbPostIndex); err == nil {
			fbPost := FeedbackPost{
				FeedbackBasics: FeedbackBasics{
					UserID:      fbPostDataDTO.UserID,
					Title:       fbPostDataDTO.Title,
					Content:     fbPostDataDTO.Content,
					Attachments: fbPostDataDTO.Attachments,
					IsAnonymous: fbPostDataDTO.IsAnonymous,
					IsSpam:      fbPostDataDTO.IsSpam,
					CreatedAt:   fbPostDataDTO.CreateAt,
					UpdatedAt:   fbPostDataDTO.UpdatedAt,
				},
				Precedence: fbPostDataDTO.Precedence,
				IsPrivate:  fbPostDataDTO.IsPrivate,
			}
			result = append(result, fbPost)
		}
	}
	return result
}

// GetLatestPublicFbPs (resCount int) -> [] struct

// GetFbPByID (postID uint64) -> struct

// GetFbPsByUserID (userID uint64) -> [] struct

// GetDetailedFbPByID (postID uint64) -> struct
