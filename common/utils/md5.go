package utils

import (
	"crypto/md5"
	"fmt"
)

//随机增强密码
const PassWorldSolt = "aYZCLuGR5cyoIjH1zBdkhkYk1uY6jGoS"

func GetMD5(data [] byte) (result string) {
	// 将字节数组转为md5加密数组
	md5Sum := md5.Sum(data)
	// 将字节数组格式化为字符串
	result = fmt.Sprintf("%x", md5Sum)
	return
}
