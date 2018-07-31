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
	go sendAaiWxasrlong()

	select {}
}

// 发送长语音识别
func sendAaiWxasrlong() {
	val, err := txAi.AaiWxasrlongForPath("./wxasrs.wav", txai.AaiAudioWAV, "http://x.x.x.x:9899/wxasrlong", "")
	log.Println(err)
	js, _ := json.Marshal(val)
	log.Println(string(js))
}

// 开启http服务
func startHttp() {
	http.HandleFunc("/wxasrlong", WxasrlongHandler)
	http.ListenAndServe("0.0.0.0:9899", nil)
}

// WxasrlongHandler 处理异步回调
func WxasrlongHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	aaiWxasrlongResponse, baseResponse, err := txai.AaiWxasrlongHandleCallback(body)

	js1, _ := json.Marshal(aaiWxasrlongResponse)
	log.Println(string(js1))

	js, _ := json.Marshal(baseResponse)
	w.Write(js)
}
