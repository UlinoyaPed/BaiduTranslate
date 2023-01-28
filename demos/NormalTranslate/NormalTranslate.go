package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"translate/BaiduTranslate"
)

func main() {

	s := `如果您第一次使用，请在本目录下创建profile.txt，并填入您的AppID和SecretKey
一行一个，不加前后缀！
请填入正确的AppID和SecretKey，否则输出为空！
例：
20220517XXXX1XXXX
jXXEOXyXXXfXXlXXsXXX

CTRL+C退出
`
	fmt.Println(s)

	deploy, err := os.OpenFile("profile.txt", os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}

	d := bufio.NewReader(deploy)
	appID, _ := d.ReadString('\n')
	appID = strings.TrimSpace(appID)
	secretKey, _ := d.ReadString('\n')
	secretKey = strings.TrimSpace(secretKey)
	btr := BaiduTranslate.BaiduInfo{AppID: appID, SecretKey: secretKey}

	for {
		fmt.Printf(">")
		q := BaiduTranslate.Input() // 从标准输入读取整行，请不要用fmt的Scanner
		to := "zh"
		lang := btr.Detect(q).Lang
		if lang == to {
			to = "en"
		}
		a1 := btr.NormalTr(q, "auto", to)
		if a1.Err() != nil {
			fmt.Println(a1.Err().Error())
		} else {
			fmt.Println(a1.Dst)
		}
	}

}
