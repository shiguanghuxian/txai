package txai

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/goroom/rand"
	"github.com/parnurzeal/gorequest"
)

var (
	// BaseURL 腾讯ai接口根地址
	BaseURL = "https://api.ai.qq.com/fcgi-bin"
)

// TxAi 腾讯ai sqk
type TxAi struct {
	AppID   string                // AppID 从腾讯AI控制台获取
	AppKey  string                // AppKey 从腾讯AI控制台获取
	request *gorequest.SuperAgent // http请求对象
	debug   bool                  // 是否调试
}

// New 创建一个sdk操作对象 := gorequest.New()
func New(appID, appKey string, debug bool) *TxAi {
	request := gorequest.New()
	request.Debug = debug
	txAi := &TxAi{
		AppID:   appID,
		AppKey:  appKey,
		request: request,
		debug:   debug,
	}
	return txAi
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
