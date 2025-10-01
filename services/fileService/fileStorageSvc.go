package fileservice

import (
	"JHETBackend/common/exception"
	configreader "JHETBackend/configs/configReader"
	"JHETBackend/dao"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/google/uuid"
)

// #####PUBLIC#####

// 统一从io读文件存盘操作
func SaveUploadedFile(ior *io.Reader, fileSHA256 []byte) (uuid.UUID, error) {

	if len(fileSHA256) == 32 { // 前端传入正确 SHA256, 尝试秒传
		flashTransUUID, ok := dao.GetFileUUIDBySHA256(fileSHA256)
		if ok { // 实现秒传
			log.Printf("[INFO][FileCtrl] Flash upload succeeded, uuid: %v", flashTransUUID)
			return flashTransUUID, nil
		}
	}

	dir := configreader.GetConfig().FileObject.Dir

	fileName := randStrGenerater(64)
	filePath := dir + "/" + fileName
	dst, err := os.Create(filePath)

	if err != nil {
		return uuid.UUID{}, err
	}
	defer dst.Close()
	_, err = io.Copy(dst, *ior)
	if err != nil {
		return uuid.UUID{}, err
	}

	// 计算SHA256 (不信任前端传入的SHA256 后端再计算一次)
	calFileSHA256 := calcFileSHA256(filePath)
	if calFileSHA256 == nil { // 无法计算 SHA256
		os.Remove(filePath) // 回滚操作 防止产生脏文件
		return uuid.UUID{}, exception.ApiFileNotSaved
	}

	// 检查这个文件是否已经注册
	dumpChkUUID, ok := dao.GetFileUUIDBySHA256(calFileSHA256)
	if ok { // 测试文件是否已经存在
		log.Printf("[INFO][FileCtrl] File already exists, uuid: %v", dumpChkUUID)
		os.Remove(filePath) // 刚刚传上来的文件直接丢掉 (浪费啦)
		return dumpChkUUID, nil
	}

	// 注册到数据库
	fileInfo, err := dst.Stat()
	if err != nil { // 获取文件属性出错 ...真的有这种情况吗?
		os.Remove(filePath) // 回滚操作 防止产生脏文件
		return uuid.UUID{}, exception.ApiFileNotSaved
	}
	objuuid, err := dao.RegFileObject(fileInfo.Name(), fileSHA256, fileInfo.Size())
	if err != nil {
		os.Remove(filePath) // 回滚操作 防止产生脏文件
		return uuid.UUID{}, exception.ApiFileNotSaved
	}

	log.Printf("[INFO][FileCtrl] New file uploaded, file: %v, uuid: %v", dst.Name(), objuuid)
	return objuuid, nil
}

// #####PRIVATE#####

// 生成随机字符串 用于临时文件名
func randStrGenerater(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

func calcFileSHA256(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil
	}

	return h.Sum(nil)
}
