package BaiduTranslate

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/tidwall/gjson"
)

// NormalResult 包含翻译结果和错误信息的结构体
type NormalResult struct {
	Dst     string // 翻译结果
	ErrCode string // 错误码
	ErrMsg  string // 错误信息
}

// NormalTr 执行普通翻译的方法
func (info *BaiduInfo) NormalTr(text, from, to string) NormalResult {
	// 生成随机盐值
	salt := Salt(10)
	// 构造待翻译的字符串，并计算签名
	sign := generateSign(info.AppID, text, salt, info.SecretKey)

	// 构造请求URL
	urlStr := "http://fanyi-api.baidu.com/api/trans/vip/translate?q=" + url.QueryEscape(text) + "&from=" + from + "&to=" + to + "&appid=" + info.AppID + "&salt=" + salt + "&sign=" + sign

	// POST请求
	method := "POST"
	// 5秒超时
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		fmt.Println(err)
		return NormalResult{ErrCode: "http.post", ErrMsg: "POST错误"}
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return NormalResult{ErrCode: "http.post", ErrMsg: "POST错误"}
	}
	defer res.Body.Close()

	// 解析返回数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return NormalResult{ErrCode: "http.read", ErrMsg: "返回信息读取出错"}
	}
	//fmt.Println(string(body))

	bodyJSON := string(body)
	var resp NormalResult
	resp.ErrCode = gjson.Get(bodyJSON, "error_code").String()
	resp.ErrMsg = gjson.Get(bodyJSON, "error_msg").String()
	if resp.ErrCode == "" {
		trans := gjson.Get(bodyJSON, "trans_result").Array()[0]
		resp.Dst = trans.Get("dst").String()
	}
	return resp
}

// Err 返回翻译错误信息
func (result NormalResult) Err() error {
	if result.ErrCode != "" {
		err := errors.New("通用翻译错误，错误码：" + result.ErrCode + "，错误信息：" + result.ErrMsg)
		return err
	} else {
		return nil
	}
}

// generateSign 根据给定的参数生成签名
func generateSign(appID, text, salt, secretKey string) string {
	// 合并字符串，计算sign
	montage := appID + text + salt + secretKey
	hash := md5.New()
	hash.Write([]byte(montage))
	sign := hex.EncodeToString(hash.Sum(nil))
	return sign
}
