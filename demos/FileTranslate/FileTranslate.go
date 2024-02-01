package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/UlinoyaPed/BaiduTranslate"
)

func init() {
	now := time.Now()
	file := "./" + now.Format("2006-01-02") + ".log"
	logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("[Translate]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// 打印欢迎消息
	s := `如果您第一次使用，请在本目录下创建profile.txt，并填入您的AppID和SecretKey
一行一个，不加前后缀！
请填入正确的AppID和SecretKey，否则输出为空！
例：
20220517XXXX1XXXX
jXXEOXyXXXfXXXXXsXXX
`
	fmt.Println(s)

	// 读取配置文件中的AppID和SecretKey
	deployData, err := ioutil.ReadFile("profile.txt")
	if err != nil {
		log.Fatal(err)
	}
	deployLines := strings.Split(string(deployData), "\n")
	appID := strings.TrimSpace(deployLines[0])
	secretKey := strings.TrimSpace(deployLines[1])
	btr := BaiduTranslate.BaiduInfo{AppID: appID, SecretKey: secretKey}

	// 读取输入文件内容
	inputData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputLines := strings.Split(string(inputData), "\n")

	// 打开输出文件
	outputFile, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	n := 0
	from := "en"
	to := "zh"
	for i, line := range inputLines {
		line = strings.TrimSpace(line)
		if line == "" {
			log.Printf("第%d行是空行", i+1)
			fmt.Println()
			_, err = writer.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if n != 1 {
				// 检测语言并切换翻译方向
				lang := btr.Detect(line)
				if lang.Err() != nil {
					log.Printf("第%d行出现错误，错误：%s", i+1, lang.Err().Error())
				} else {
					fmt.Printf("语言：%s\n", lang.Lang)
					if lang.Lang == to {
						from, to = to, from
					}
					from = lang.Lang
					n++
				}
			}

			// 进行翻译
			result := btr.NormalTr(line, from, to)
			if result.Err() != nil {
				log.Printf("第%d行出现错误，错误：%s", i+1, result.Err().Error())
				_, err = writer.WriteString("\n")
			} else {
				fmt.Printf("%s\n", result.Dst)
				log.Printf("第%d行翻译完成！", i+1)
				_, err = writer.WriteString(result.Dst + "\n")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	fmt.Println("---翻译已完成，按回车键退出---")
	log.Println("---翻译已完成---")

	// 刷新缓冲区并关闭写入器
	writer.Flush()
	BaiduTranslate.Input()
}
