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

	txAi := txai.New("1106736025", "9ea4yNLi2jrSc66y", true)
	val, err := txAi.VisionImgfilterForPath("../../img/face_detectface.jpg", "110", 31)
	log.Println(err)
	ioutil.WriteFile("abc.jpg", val.Data.Image, 0655)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}
