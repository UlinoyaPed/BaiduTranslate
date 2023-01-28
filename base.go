package BaiduTranslate

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

// 百度翻译开放平台信息
type BaiduInfo struct {
	AppID     string
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

// 输入工具，可直接输入整行，不必为fmt.Scan的空格分词所困扰
func Input() string {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	return s
}
