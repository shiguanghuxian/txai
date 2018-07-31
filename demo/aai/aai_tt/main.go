package main

import (
	"encoding/json"
	"io/ioutil"
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

	tts() // 语音合成（AI Lab）
	tta() // 语音合成（优图）
}

func tts() {
	val, err := txAi.AaiTtsForText("你好吗天气好吗英文版", txai.AaiTtsSpeaker5, txai.AaiTtsFormatMP3, 0, 100, 0, 58)
	log.Println(err)
	ioutil.WriteFile("abc1.mp3", val.Data.Speech, 0655)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

func tta() {
	val, err := txAi.AaiTtaForText("你好吗天气好吗英文版", txai.AaiTtaModelType0, 0)
	log.Println(err)
	ioutil.WriteFile("abc2.mp3", val.Data.Voice, 0655)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
