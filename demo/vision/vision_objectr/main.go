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
	val, err := txAi.VisionObjectrForPath("../../img/vision_porn.jpg", 1, 5)
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
