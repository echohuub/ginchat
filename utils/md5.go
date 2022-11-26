package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

func MD5Encode(input string) string {
	return strings.ToUpper(Md5Encode(input))
}

func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

func ValidPassword(plainpwd, salt string, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
