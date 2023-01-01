package BaiduTranslate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func (BaiduInfo *BaiduInfo) Detect(Text string) (string, error) {

	type DetectResult struct {
		Data struct {
			Src string `json:"src"`
		} `json:"data"`
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	}

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

	var ts DetectResult
	_ = json.Unmarshal(body, &ts)

	if ts.ErrorCode != "" {
		err := errors.New("错误码：" + ts.ErrorCode + "，错误信息：" + ts.ErrorMsg)
		return "", err
	} else {
		return ts.Data.Src, nil
	}
}
