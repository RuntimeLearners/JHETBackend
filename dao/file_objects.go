package dao

import (
	"JHETBackend/configs/database"
	"JHETBackend/models"

	"github.com/google/uuid"
)

// 将下载的临时文件注册到数据库并落盘
func RegFileObject(filename string, sha256 []byte, filesize int64) (uuid.UUID, error) {
	newUUID, _ := uuid.New().MarshalBinary()
	newFileObj := models.FileObject{
		UUID:     newUUID,
		Sha256:   sha256,
		FileSize: filesize,
		FileName: filename,
		FileType: "general", // TODO: 文件类型检查
	}
	dbnp := database.DataBase.Create(&newFileObj)
	if dbnp.Error != nil {
		return uuid.UUID{}, dbnp.Error
	}
	return uuid.UUID(newUUID), nil
}

func GetFileUUIDBySHA256(target []byte) (uuid.UUID, bool) {
	var fileObj models.FileObject
	database.DataBase.Model(&models.FileObject{}).Where("sha256 = ?", target).First(&fileObj)
	return uuid.UUID(fileObj.UUID), fileObj.ID != 0 // 直接通过 ID 判断是否存在
}
