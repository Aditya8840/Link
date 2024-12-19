package utils

import (
	"encoding/base64"
	"encoding/binary"
)


func Base64Encode(num int64) string {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(num))
	return base64.StdEncoding.EncodeToString(buf)
}