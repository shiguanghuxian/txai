package main

import (
	"encoding/json"
	"log"

	"github.com/shiguanghuxian/txai"
)

var (
	txAi *txai.TxAi
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	txAi = txai.New("appid", "appkey", true)

	val, err := txAi.AaiEvilaudioForUrl("demo", "https://ai.qq.com/cgi-bin/appdemo_ttsecho?text=欢迎使用腾讯AI。&speaker=1&volume=0&speed=100&format=3&aht=0&apc=58&download=1")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
