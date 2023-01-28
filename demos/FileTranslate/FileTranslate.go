package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"translate/BaiduTranslate"

	"github.com/dablelv/go-huge-util/file"
)

func init() {
	now := time.Now()
	file := "./" + now.Format("2006-01-02") + ".log" // 使用2006年1月2号15点04分
	logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Println(err.Error())
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[Translate]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return
}

func FindFile(dirname string, format string) ([]string, error) {
	names, err := file.GetDirAllEntryPaths(dirname, true)
	if err != nil {
		log.Println(err.Error())
	}
	var result []string
	ext := ""
	for i := 0; i < len(names); i++ {
		ext = path.Ext(names[i])
		if ext == format {
			result[i] = names[i]
		}
	}
	return result, err
}

func main() {
	s := `如果您第一次使用，请在本目录下创建profile.txt，并填入您的AppID和SecretKey
一行一个，不加前后缀！
请填入正确的AppID和SecretKey，否则输出为空！
例：
20220517XXXX1XXXX
jXXEOXyXXXfXXXXXsXXX
`
	fmt.Println(s)

	deploy, err := os.OpenFile("profile.txt", os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err.Error())
	}

	d := bufio.NewReader(deploy)
	appID, _ := d.ReadString('\n')
	appID = strings.TrimSpace(appID)
	secretKey, _ := d.ReadString('\n')
	secretKey = strings.TrimSpace(secretKey)
	btr := BaiduTranslate.BaiduInfo{AppID: appID, SecretKey: secretKey}

	in, err := os.OpenFile("input.txt", os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err.Error())
	}
	// 创建 Reader
	r := bufio.NewReader(in)

	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	i := 0
	n := 0
	from := "en"
	to := "zh"
	for {
		line, ierr := r.ReadString('\n')
		line = strings.TrimSpace(line) // 去掉字符串首尾空白字符
		if ierr != nil && ierr != io.EOF {
			log.Println(err.Error())
		}

		i++

		if line == "" {
			log.Printf("第%d行是空行", i)
			fmt.Println()

			_, err = writer.WriteString("\n")
			if err != nil {
				log.Println(err.Error())
			}

		} else {
			if n != 1 {
				lang := btr.Detect(line)
				if lang.Err() != nil {
					log.Println(lang.Err().Error())
				} else {
					fmt.Printf("语言：%s\n", lang.Lang)
					if lang.Lang == to {
						from, to = to, from
					}
					from = lang.Lang
					n++
				}
			}

			result := btr.NormalTr(line, from, to)
			if result.Err() != nil {
				log.Printf("第%d行出现错误，错误：%s", i, result.Err().Error())
				_, err = writer.WriteString("\n")
			} else {
				fmt.Printf("%s\n", result.Dst)
				log.Printf("第%d行翻译完成！", i)

				_, err = writer.WriteString(result.Dst + "\n")
				if err != nil {
					log.Println(err.Error())
				}
			}

		}

		if ierr == io.EOF {
			break
		}

	}

	fmt.Println("---翻译已完成，按回车键退出---")
	log.Println("---翻译已完成---")

	writer.Flush()
	BaiduTranslate.Input()
	return
}
