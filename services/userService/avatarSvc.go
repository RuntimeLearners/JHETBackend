package userService

import (
	"JHETBackend/common/exception"
	"JHETBackend/dao"
	fileservice "JHETBackend/services/fileService"
	"io"
	"log"
)

// 此处调用fileService存盘，然后调用dao更新数据库中用户头像信息
func UploadAvatar(accountID uint64, fileHandler io.Reader) error {
	fileName, err := fileservice.SaveUploadedFile(&fileHandler)
	if err != nil {
		log.Printf("[ERROR][UserSvc/AvatarSvc] 不能保存用户头像文件 %v", err)
		return exception.ApiFileNotSaved // 转换成统一错误返回，原error信息丢失
	}
	err = dao.UpdateAccountAvatar(accountID, fileName)
	if err != nil {
		return err
	}
	return nil
}
