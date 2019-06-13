package main

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
)

func Md5(v string) string {
	sum := md5.Sum([]byte(v))
	return hex.EncodeToString(sum[:])
}

func UUID() string {
	return uuid.NewV4().String()
}
