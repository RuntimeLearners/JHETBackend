package webtoken

import (
	configreader "JHETBackend/configs/configReader"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"log"
	"sync"
	"time"
)

// 签名密钥 第一次使用时从配置读取 48位是为了适配base64
var wtSigkey [48]byte

var readKeyOnce sync.Once

/* wtoken 格式(48位)
** ######## #### ######## #### ########################
** 0-------8----12-------20---24----------------------48
** 用户ID--权限组--过期时间-预留-24字节签名
** 签名方式: SHA256(Metadata(即前24位字节)+Key(48字节))取前24字节
** 全部占据48字节, 转换到base64刚好64字符
 */

// Generate 生成 token
func GenerateWt(uid uint64, permGroupID uint32, validSecs uint64) string {

	readKeyOnce.Do(readSigkey) // 读入签名密钥 (只执行一次)

	// 打包 24B 元信息
	var metadata [24]byte
	binary.LittleEndian.PutUint64(metadata[0:8], uid)
	binary.LittleEndian.PutUint32(metadata[8:12], permGroupID)
	var expireTime = time.Now().Unix() + int64(validSecs) // 计算过期时间
	binary.LittleEndian.PutUint64(metadata[12:20], uint64(expireTime))

	// 计算签名
	// append 只接受切片+元素，因此拆分wtSigKey
	hashResult := sha256.Sum256(append(metadata[:], wtSigkey[:]...))
	sig := hashResult[:24] // 取前 24 B

	// 拼接并 base64
	var tokenResult [48]byte
	copy(tokenResult[:24], metadata[:])
	copy(tokenResult[24:], sig)

	log.Print("[INFO][wtService] A wtoken generated.")
	return base64.RawURLEncoding.EncodeToString(tokenResult[:])
}

// Verify 校验 token，返回载荷与是否有效
func VerifyWt(webtoken string) (isValid bool) {

	readKeyOnce.Do(readSigkey)

	// 验证长度与base64规则
	raw, err := base64.RawURLEncoding.DecodeString(webtoken)
	if err != nil || len(raw) != 48 {
		log.Print("[WARN][wtService] A token has invalid format!")
		return false
	}

	content, sig := raw[:24], raw[24:]
	// 重算签名
	hashResult := sha256.Sum256(append(content, wtSigkey[:]...))
	if !bytes.Equal(hashResult[:24], sig) {
		log.Print("[WARN][wtService] A token is invalid!")
		return false
	}

	// 检查是否过期，过期token按无效处理
	var validBefore = binary.LittleEndian.Uint64(content[12:20])
	if time.Now().Unix() >= int64(validBefore) {
		log.Print("[INFO][wtService] A token verified but expired.")
		return false
	}
	log.Printf("[INFO][wtService] A token verified. It will expire in %vsec", int64(validBefore)-time.Now().Unix())
	return true
}

func GetWtPayload(webtoken string) (uid uint64, permGroupID uint32, err error) {
	tokdec, err := base64.RawURLEncoding.DecodeString(webtoken)
	if err != nil || len(tokdec) != 48 {
		return 0, 0, err
	}
	return binary.LittleEndian.Uint64(tokdec[0:8]),
		binary.LittleEndian.Uint32(tokdec[8:12]),
		nil
}

func readSigkey() {
	sigkeytmp, err := base64.StdEncoding.DecodeString(configreader.GetConfig().WebtokenSigkey)
	if err != nil {
		log.Panicln("[FATAL] 无法载入 wt 签名密钥")
	}
	if len(sigkeytmp) != 48 {
		log.Panicf("[FATAL] wt签名密钥长度错误 期望32位 实际%v位", len(sigkeytmp))
	}
	copy(wtSigkey[:], sigkeytmp)
}
