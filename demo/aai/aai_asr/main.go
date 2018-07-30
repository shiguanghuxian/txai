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

	txAi = txai.New("1106736025", "9ea4yNLi2jrSc66y", true)

	asr() // 语音识别-echo版
}

func asr() {
	val, err := txAi.AaiAsrForPath("../../img/wxasrs.wav", 2, 16000)
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
