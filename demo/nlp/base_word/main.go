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

	txAi = txai.New("appid", "appkey", false)

	wordseg() // 分词
	wordpos() // 词性标注
	wordner() // 专有名词识别
	wordsyn() // 同义词识别
}

func wordseg() {
	val, err := txAi.NlpWordsegForText("腾讯人工智能")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

func wordpos() {
	val, err := txAi.NlpWordposForText("腾讯人工智能")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

func wordner() {
	val, err := txAi.NlpWordnerForText("最近张学友在深圳开了一场演唱会")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

func wordsyn() {
	val, err := txAi.NlpWordsynForText("今天的天气怎么样")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
