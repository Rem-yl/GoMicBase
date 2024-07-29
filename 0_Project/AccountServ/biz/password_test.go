package biz

import (
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/anaskhan96/go-password-encoder"
)

func TestGetMd5(t *testing.T) {
	s := "rem"
	fmt.Println(GetMd5(s))
	s = "ram"
	fmt.Println(GetMd5(s))
	s = "rea"
	fmt.Println(GetMd5(s))

	salt, encode := password.Encode("rem", nil)
	fmt.Println(salt)
	fmt.Println(encode)

	check := password.Verify("rem", salt, encode, nil)
	fmt.Println(check)
	check1 := password.Verify("rem1", salt, encode, nil)
	fmt.Println(check1)

	options := password.Options{
		SaltLen:      16,
		Iterations:   100,
		KeyLen:       32,
		HashFunction: md5.New,
	}

	salt, encode = password.Encode("rem", &options)
	fmt.Println(salt)
	fmt.Println(encode)

}
