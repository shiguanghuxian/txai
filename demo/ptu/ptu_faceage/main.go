package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/shiguanghuxian/txai"
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	txAi := txai.New("appid", "appkey", true)
	val, err := txAi.PtuFaceageForPath("../../img/face_detectface.jpg")
	log.Println(err)
	ioutil.WriteFile("abc.jpg", val.Data.Image, 0655)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
