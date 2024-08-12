package share

import (
	"crypto/md5"
	"hash"

	"github.com/anaskhan96/go-password-encoder"
)

var PasswordOption password.Options

func newPasswordOption() {
	PasswordOption = password.Options{
		SaltLen:      8,
		Iterations:   0,
		KeyLen:       32,
		HashFunction: func() hash.Hash { return md5.New() },
	}
}

func init() {
	newPasswordOption()
}
