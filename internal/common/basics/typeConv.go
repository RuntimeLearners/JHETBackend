package basics

import (
	"encoding/binary"
	"errors"
)

func Uint64sToBytes(u []uint64) ([]byte, error) {
	if len(u) == 0 {
		return nil, nil
	}
	buf := make([]byte, 8*len(u))
	for i, v := range u {
		binary.LittleEndian.PutUint64(buf[i*8:], v)
	}
	return buf, nil
}

// Scan 把 []byte -> []uint64
func BytesToUint64s(b []byte) ([]uint64, error) {
	if len(b)%8 != 0 {
		return nil, errors.New("invalid byte length: not a multiple of 8")
	}
	out := make([]uint64, len(b)/8)
	for i := 0; i < len(out); i++ {
		out[i] = binary.LittleEndian.Uint64(b[i*8:])
	}
	return out, nil
}
