# 腾讯AI SDK
腾讯AI开发平台sdk

腾讯AI接口调用比较简单，可以通过此库简化调用流程，返回值是结构体方便使用。
开发者可以不去考虑签名、数据请求就像调本地函数一样调用接口

## 示例
```
package main

import (
	"encoding/json"
	"log"

	"github.com/shiguanghuxian/txai" // 引入sdk
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 创建sdk操作对象
	txAi := txai.New("appid", "appkey", true)
	// 调用对应腾讯ai接口的对应函数
	val, err := txAi.ImageFoodForPath("../../img/image_terrorism.jpg")
	// 打印结果
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

```


## 备注
由于腾讯AI部分接口使用gbk格式，本sdk已自动将格式转换，使用时无需考虑编码问题。
