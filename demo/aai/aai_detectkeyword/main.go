package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/shiguanghuxian/txai"
)

var txAi = txai.New("appid", "appkey", true)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	go startHttp()
	time.Sleep(2 * time.Second)
	go sendAaiDetectkeyword()

	select {}
}

// 发送关键词检索请求
func sendAaiDetectkeyword() {
	val, err := txAi.AaiDetectkeywordForPath("./wxasrs.wav", "", txai.AaiAudioWAV, "http://x.x.x.x:9899/detectkeyword", "腾讯", "开放平台")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

// 开启http服务
func startHttp() {
	http.HandleFunc("/detectkeyword", WxasrlongHandler)
	http.ListenAndServe("0.0.0.0:9899", nil)
}

// WxasrlongHandler 处理异步回调
func WxasrlongHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))
	aaiDetectkeywordResponse, baseResponse, err := txai.AaiDetectkeywordHandleCallback(body)

	js1, _ := json.Marshal(aaiDetectkeywordResponse)
	log.Println(string(js1))

	js, _ := json.Marshal(baseResponse)
	w.Write(js)
}
