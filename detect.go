package BaiduTranslate

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)

type DetectResult struct {
	Lang string

	errCode string
	errMsg  string
}

func (BaiduInfo *BaiduInfo) Detect(Text string) DetectResult {

	//合并字符串，计算sign
	salt := Salt(10)
	montage := BaiduInfo.AppID + Text + salt + BaiduInfo.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(montage))
	sign := hex.EncodeToString(ctx.Sum(nil))

	// 拼接完整url
	urlstr := "http://fanyi-api.baidu.com/api/trans/vip/language?q=" + url.QueryEscape(Text) + "&salt=" + salt + "&sign=" + sign + "&appid=" + BaiduInfo.AppID

	// 发送GET请求
	resp, err := http.Get(urlstr)
	if err != nil {
		log.Println("HTTP GET出现错误！")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	// json解析
	bodyJson := string(body)
	var res DetectResult
	res.errCode = gjson.Get(bodyJson, "error_code").String()
	res.errMsg = gjson.Get(bodyJson, "error_msg").String()
	if res.errCode == "0" {
		res.Lang = gjson.Get(bodyJson, "data.src").String()
	}

	return res

}

func (j DetectResult) Err() error {
	if j.errCode != "0" {
		err := errors.New("语种识别错误，错误码：" + j.errCode + "，错误信息：" + j.errMsg)
		return err
	} else {
		return nil
	}
}

/*
if errorCode != "0" {
		err := errors.New("错误码：" + errorCode + "，错误信息：" + errorMsg)
		return "", err
	} else {
		return result, nil
	}
*/
