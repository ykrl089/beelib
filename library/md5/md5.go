package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}
