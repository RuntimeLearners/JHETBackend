package webtoken_test

import (
	webtokenservice "JHETBackend/internal/common/webtoken"
	"encoding/base64"
	"encoding/binary"
	"log"
	"testing"
)

func TestWebtokenService(t *testing.T) {
	uid := uint64(123456789)
	pgid := uint32(88)
	expt := uint64(120)

	tok := webtokenservice.GenerateWt(uid, pgid, expt)
	log.Println("token:", tok)

	tokdec, err := base64.RawURLEncoding.DecodeString(tok)
	if err != nil || len(tokdec) != 48 {
		t.Error("not base64")
	}

	log.Printf("uid: %v pgid: %v expt: %v",
		binary.LittleEndian.Uint64(tokdec[0:8]),
		binary.LittleEndian.Uint32(tokdec[8:12]),
		binary.LittleEndian.Uint64(tokdec[12:20]))
	if uid == binary.LittleEndian.Uint64(tokdec[0:8]) &&
		pgid == binary.LittleEndian.Uint32(tokdec[8:12]) {
	} else {
		t.Error("metadata not match!")
	}

	// 校验
	isValid := webtokenservice.VerifyWt(tok)
	if !isValid {
		t.Error("not valid!")
		return
	}
}

func TestKey(t *testing.T) {
	// 手动复制生成的tok到这里测试过期
	tok := "Fc1bBwAAAABYAAAAxFvNaAAAAAAAAAAA8XvXmLW9pn1gT2VPmcqqtrtNNC-jKGYa"
	isValid := webtokenservice.VerifyWt(tok)
	if !isValid {
		t.Error("not valid!..")
		return
	}
}
