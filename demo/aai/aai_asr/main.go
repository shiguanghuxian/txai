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

	asr()    // 语音识别-echo版
	asrs()   // 语音识别-流式版（AI Lab）
	wxasrs() // 语音识别-流式版(WeChat AI)
}

func asr() {
	val, err := txAi.AaiAsrForPath("../../img/wxasrs.wav", txai.AaiAudioWAV, 16000)
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

func asrs() {
	val, err := txAi.AaiAsrsForPath("../../img/wxasrs.wav", txai.AaiAudioWAV, 0, 1, "110", 16000)
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

func wxasrs() {
	val, err := txAi.AaiWxasrsForPath("../../img/wxasrs.wav", txai.AaiAudioWAV, 16000, 16, 0, 1, "110", 1)
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
