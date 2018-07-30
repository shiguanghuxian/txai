package txai

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/goroom/rand"
)

const (
	// BaseURL 腾讯ai接口根地址
	BaseURL = "https://api.ai.qq.com/fcgi-bin"
)

// TxAi 腾讯ai sqk
type TxAi struct {
	AppID   string        // AppID 从腾讯AI控制台获取
	AppKey  string        // AppKey 从腾讯AI控制台获取
	timeout time.Duration // http请求超时
	debug   bool          // 是否调试
}

// New 创建一个sdk操作对象 := gorequest.New()
func New(appID, appKey string, debug bool) *TxAi {
	txAi := &TxAi{
		AppID:   appID,
		AppKey:  appKey,
		timeout: 30 * time.Second,
		debug:   debug,
	}
	return txAi
}

// SetDebug 设置是否调试模式
func (ai *TxAi) SetDebug(debug bool) {
	ai.debug = debug
}

// SetRequestTimeout 设置超时时间
func (ai *TxAi) SetRequestTimeout(d time.Duration) {
	if d == 0 {
		d = 30 * time.Second
	}
	ai.timeout = d
}

// URLToBase64 读取url返回base64字符串
func (ai *TxAi) URLToBase64(imageURL string) (string, error) {
	client := http.Client{Timeout: ai.timeout}
	resp, err := client.Get(imageURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return img, nil
}

// 计算签名
func (ai *TxAi) getReqSign(u url.Values) (sign string) {
	byte := []byte(fmt.Sprintf("%s&%s=%s", u.Encode(), "app_key", ai.AppKey))
	// 计算md5值
	newmd := md5.New()
	newmd.Write(byte)
	sign_byte := newmd.Sum(nil)
	sign = strings.ToUpper(fmt.Sprintf("%x", sign_byte))
	return sign
}

// 生成随机字符串
func (ai *TxAi) nonceStr() string {
	return rand.GetRand().String(16, rand.RST_NUMBER)
}

// 获取公共参数
func (ai *TxAi) getPublicParams() (u url.Values) {
	u = url.Values{}
	u.Add("app_id", ai.AppID)
	u.Add("time_stamp", fmt.Sprint(time.Now().Unix()))
	u.Add("nonce_str", ai.nonceStr())
	return
}
