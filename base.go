package BaiduTranslate

import (
	"log"
	"math/rand"
	"time"
)

// 百度翻译开放平台信息
type BaiduInfo struct {
	AppID     string
	Salt      string
	SecretKey string
}

// 自动生盐
// 入口参数为盐的长度
func Salt(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 错误检测，如果有错误输出msg
func checkErr(e error, msg string) {
	if e != nil {
		log.Println(msg)
	}
	return
}
