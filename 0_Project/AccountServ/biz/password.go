package biz

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/anaskhan96/go-password-encoder"
)

var DefaultOptions = password.Options{
	SaltLen:      16,
	Iterations:   100,
	KeyLen:       32,
	HashFunction: md5.New,
}

func GetMd5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
