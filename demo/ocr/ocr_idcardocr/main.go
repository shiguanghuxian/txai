package main

import (
	"log"

	"github.com/shiguanghuxian/txai"
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	txAi := txai.New("appid", "appkey", true)
	val, err := txAi.OcrIdcardocrForPath(0, "../../img/id_test.jpg")
	log.Println(err)
	// js, _ := json.Marshal(val)
	// log.Println(string(js))
	log.Println("姓名：", val.Data.Name, "身份证号：", val.Data.Id)

}
