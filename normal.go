package BaiduTranslate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (BaiduInfo *BaiduInfo) NormalTr(Text string, From string, To string) string {

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
	montage := BaiduInfo.AppID + Text + BaiduInfo.Salt + BaiduInfo.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(montage))
	sign := hex.EncodeToString(ctx.Sum(nil))

	// 翻译 传入需要翻译的语句
	urlstr := "http://fanyi-api.baidu.com/api/trans/vip/translate?q=" + url.QueryEscape(Text) + "&from=" + From + "&to=" + To + "&appid=" + BaiduInfo.AppID + "&salt=" + BaiduInfo.Salt + "&sign=" + sign

	// 发送GET请求
	resp, err := http.Get(urlstr)
	checkErr(err, "HTTP GET出现错误！")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ts TransResult
	_ = json.Unmarshal(body, &ts)
	if ts.ErrorCode != "" {
		errmsg := "错误码：" + ts.ErrorCode + "，错误信息：" + ts.ErrorMsg
		return errmsg
	} else {
		return ts.Result[0].Dst
	}
}

/*

// 返回结果
type TransResult struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Result    [1]Result `json:"trans_result"`
	ErrorCode string    `json:"error_code"`
	ErrorMsg  string    `json:"error_msg"`

	Data [1]Data `json:"data"`
}
type Result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
type Data struct {
	Src string `json:"src"`
}

// 生成32位MD5
func Sign(bi *BaiduInfo) string {
	text := bi.AppID + bi.Text + bi.Salt + bi.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 翻译 传入需要翻译的语句
func (bi *BaiduInfo) Translate() string {
	url := "http://api.fanyi.baidu.com/api/trans/vip/translate?q=" + bi.Text + "&from=" + bi.From + "&to=" + bi.To + "&appid=" + bi.AppID + "&salt=" + bi.Salt + "&sign=" + Sign(bi)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("网络异常")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var ts TransResult
	_ = json.Unmarshal(body, &ts)
	if ts.ErrorCode != "" {
		return ts.ErrorMsg
	} else {
		return ts.Result[0].Dst
	}
}

func (bi *BaiduInfo) Detect() string {
	url := "http://api.fanyi.baidu.com/api/trans/vip/language?q=" + bi.Text + "&appid=" + bi.AppID + "&salt=" + bi.Salt + "&sign=" + Sign(bi)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("网络异常")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var ts TransResult
	_ = json.Unmarshal(body, &ts)
	if ts.ErrorCode != "" {
		return ts.ErrorCode
	} else {
		return ts.Data[0].Src
	}
}

*/
