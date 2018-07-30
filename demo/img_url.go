package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/shiguanghuxian/txai"
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	txAi := txai.New("appid", "appkey", true)
	txAi.SetDebug(false)
	img, err := txAi.ImgURLToBase64("https://yyb.gtimg.com/aiplat/static/ai-demo/large/odemo-pic-2.jpg")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	txAi.SetDebug(true)
	val, err := txAi.OcrBcocrForBase64(img)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
