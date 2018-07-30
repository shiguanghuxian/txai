package main

import (
	"encoding/json"
	"log"

	"github.com/shiguanghuxian/txai"
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	txAi := txai.New("appid", "appkey", false)
	val, err := txAi.NlpTextpolarForText("今天和明天的天气怎么样？")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
