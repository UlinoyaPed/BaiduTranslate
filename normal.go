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

type NormalResult struct {
	Dst string

	errCode string
	errMsg  string
}

func (BaiduInfo *BaiduInfo) NormalTr(Text string, From string, To string) NormalResult {

	//合并字符串，计算sign
	salt := Salt(10)
	montage := BaiduInfo.AppID + Text + salt + BaiduInfo.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(montage))
	sign := hex.EncodeToString(ctx.Sum(nil))

	// 传入需要翻译的语句
	urlstr := "http://fanyi-api.baidu.com/api/trans/vip/translate?q=" + url.QueryEscape(Text) + "&from=" + From + "&to=" + To + "&appid=" + BaiduInfo.AppID + "&salt=" + salt + "&sign=" + sign

	// 发送GET请求
	resp, err := http.Get(urlstr)
	if err != nil {
		log.Println("HTTP GET出现错误！")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	bodyJson := string(body)
	var res NormalResult
	res.errCode = gjson.Get(bodyJson, "error_code").String()
	res.errMsg = gjson.Get(bodyJson, "error_msg").String()
	if res.errCode == "" {
		trans := gjson.Get(bodyJson, "trans_result").Array()[0]
		res.Dst = trans.Get("dst").String()
	}
	return res
}

func (j NormalResult) Err() error {
	if j.errCode != "" {
		err := errors.New("通用翻译错误，错误码：" + j.errCode + "，错误信息：" + j.errMsg)
		return err
	} else {
		return nil
	}
}

/*
// json解析
	trans := gjson.Get(string(body), "trans_result").Array()[0]
	result := trans.Get("dst").String()
	errorCode := gjson.Get(string(body), "error_code").String()
	errorMsg := gjson.Get(string(body), "error_msg").String()
	if errorCode != "" {
		err := errors.New("错误码：" + errorCode + "，错误信息：" + errorMsg)
		return "", err
	} else {
		return result, nil
	}
*/
