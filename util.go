package main

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(v string) string {
	sum := md5.Sum([]byte(v))
	return hex.EncodeToString(sum[:])
}
