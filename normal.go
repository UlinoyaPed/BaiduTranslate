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

func (BaiduInfo *BaiduInfo) NormalTr(Text string, From string, To string) (string, error) {

	type TransResult struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Result [1]struct {
			Src string `json:"src"`
			Dst string `json:"dst"`
		} `json:"trans_result"`
		ErrorCode string `json:"error_code"`
		ErrorMsg  string `json:"error_msg"`
	}

	//合并字符串，计算sign
	salt := Salt(10)
	montage := BaiduInfo.AppID + Text + salt + BaiduInfo.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(montage))
	sign := hex.EncodeToString(ctx.Sum(nil))

	// 翻译 传入需要翻译的语句
	urlstr := "http://fanyi-api.baidu.com/api/trans/vip/translate?q=" + url.QueryEscape(Text) + "&from=" + From + "&to=" + To + "&appid=" + BaiduInfo.AppID + "&salt=" + salt + "&sign=" + sign

	// 发送GET请求
	resp, err := http.Get(urlstr)
	if err != nil {
		log.Println("HTTP GET出现错误！")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ts TransResult
	_ = json.Unmarshal(body, &ts)
	if ts.ErrorCode != "" {
		err := errors.New("错误码：" + ts.ErrorCode + "，错误信息：" + ts.ErrorMsg)
		return "", err
	} else {
		return ts.Result[0].Dst, nil
	}
}
