package main

import (
	"encoding/json"
	"log"

	"github.com/shiguanghuxian/txai"
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	txAi := txai.New("1106736025", "9ea4yNLi2jrSc66y", true)
	val, err := txAi.NlpImagetranslateForPath("../../img/nlp_imagetranslate.jpg", "110", txai.SceneDoc, txai.En, txai.Zh)
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
