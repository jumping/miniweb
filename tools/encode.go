package tools

import (
	"crypto/md5"
	"fmt"
)

// MD5处理函数
func MD5(str string) string {
	data := []byte(str)
	result := md5.Sum(data)
	
	return fmt.Sprintf("%x", result)
}