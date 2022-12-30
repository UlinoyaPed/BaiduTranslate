# golang的百度在线翻译库

该库用于调用百度翻译API,实现翻译效果。使用前,需要注册 [百度翻译开放平台信息](http://api.fanyi.baidu.com/api/trans/product/index)

## 安装

```bash
go install github.com/UlinoyaPed/BaiduTranslate
```

## 使用

```go
package main

import (
	"fmt"

	"github.com/UlinoyaPed/BaiduTranslate"
)

func main() {
	//输入基本信息，Salt长度无要求
    // BaiduInfo结构体记录配置项，Salt为数据传送时加盐（库中已有函数实现，可直接调用）
	btr := BaiduTranslate.BaiduInfo{AppID: "", Salt: BaiduTranslate.Salt(5), SecretKey: ""}

	// 通用翻译
	// 传入：(原文, 原文语言, 译文语言)
	fmt.Println(btr.NormalTr("Hello world!", "en", "zh"))
	fmt.Println(btr.NormalTr("百度翻译", "auto", "de"))
}

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
