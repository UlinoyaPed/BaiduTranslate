# golang的百度在线翻译库

该库用于调用百度翻译API,实现翻译效果。使用前,需要注册 [百度翻译开放平台信息](http://api.fanyi.baidu.com/api/trans/product/index)

## 安装

```bash
go get github.com/UlinoyaPed/BaiduTranslate@l
```

## 使用

### 通用翻译

```go
package main

import (
	"fmt"

	"github.com/UlinoyaPed/BaiduTranslate"
)

func main() {
	// 输入基本信息，BaiduInfo结构体记录配置项
	btr := BaiduTranslate.BaiduInfo{AppID: "XXX", SecretKey: "XXX"}

	// 传入：(原文, 原文语言, 译文语言)
	// 完整实例
	s1, err := btr.NormalTr("Hello world!", "en", "zh") // 对原文进行了url encode，原文可带空格
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s1)
	}

	// 忽略错误
	s2, _ := btr.NormalTr("百度翻译", "auto", "de")
	fmt.Printf("%s\n", s2)

	fmt.Println("---以下为错误示范---")
	// 语言不能带空格，否则会报错！
	// w1, err := btr.NormalTr("百度翻译", " zh", "en")
	// fmt.Println(err.Error(), w1)

	// 无"fr"语言（法语为"fra"）
	w2, err := btr.NormalTr("百度翻译", "auto", "fr")
	fmt.Println(err.Error(), w2)

	// 不能缺少参数
	w3, err := btr.NormalTr("百度翻译", "", "en")
	fmt.Println(err.Error(), w3)
}

```

**输出**

```go
你好，世界！
Baidu Übersetzen
---以下为错误示范---
错误码：58001，错误信息：INVALID_TO_PARAM 
错误码：54000，错误信息：PARAM_FROM_TO_OR_Q_EMPTY 
```

### 语种检测

```go
package main

import (
	"fmt"

	"github.com/UlinoyaPed/BaiduTranslate"
)

func main() {
	// 输入基本信息，BaiduInfo结构体记录配置项
	btr := BaiduTranslate.BaiduInfo{AppID: "XXX", SecretKey: "XXX"}

	// 完整实例
	s1, err := btr.Detect("百度翻译")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s1)
	}

	// 忽略错误
	s2, _ := btr.Detect("Hello World!")
	fmt.Println(s2)

	fmt.Println("---以下为错误示范---")
	
	//不能缺少参数
	w1, err := btr.Detect("")
	fmt.Println(err.Error(), w1)

}

```

**输出**

```go
zh
en
---以下为错误示范---
错误码：54000，错误信息：PARAM_FROM_TO_OR_Q_EMPTY
```

## 受支持的翻译语言

 **(源语言语种不确定时可设置为 auto，目标语言语种不可设置为 auto)**

| 语言简写 |     名称     |
| :------: | :----------: |
|   auto   |   自动检测   |
|    zh    |     中文     |
|    en    |     英语     |
|   yue    |     粤语     |
|   wyw    |    文言文    |
|    jp    |     日语     |
|   kor    |     韩语     |
|   fra    |     法语     |
|   spa    |   西班牙语   |
|    th    |     泰语     |
|   ara    |   阿拉伯语   |
|    ru    |     俄语     |
|    pt    |   葡萄牙语   |
|    de    |     德语     |
|    it    |   意大利语   |
|    el    |    希腊语    |
|    nl    |    荷兰语    |
|    pl    |    波兰语    |
|   bul    |  保加利亚语  |
|   est    |  爱沙尼亚语  |
|   dan    |    丹麦语    |
|   fin    |    芬兰语    |
|    cs    |    捷克语    |
|   rom    |  罗马尼亚语  |
|   slo    | 斯洛文尼亚语 |
|   swe    |    瑞典语    |
|    hu    |   匈牙利语   |
|   cht    |   繁体中文   |
|   vie    |    越南语    |
