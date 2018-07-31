package txai

// 语音识别

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	aaiDetectkeywordURI = "/aai/aai_detectkeyword"
)

// AaiDetectkeywordForBase64 关键词检索 - 音频为base64格式
// speech 待识别语音（时长上限15s）
// format 语音压缩格式编码，定义见 https://ai.qq.com/doc/detectword.shtml
// callbackURL 用户回调url，需用户提供，用于平台向用户通知识别结果
// keyWords 待识别关键词
// speechURL 待识别语音下载地址（时长上限15min）
func (ai *TxAi) AaiDetectkeywordForBase64(speech, speechURL string, format AaiAudioType, callbackURL string, keyWords ...string) (*AaiDetectkeywordResponse, error) {
	params := ai.getPublicParams()
	if speech != "" {
		params.Add("speech", speech)
	}
	if speechURL != "" {
		params.Add("speech_url", speechURL)
	}
	params.Add("format", fmt.Sprint(format))
	params.Add("callback_url", callbackURL)
	params.Add("key_words", strings.Join(keyWords, "|"))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)

	// 响应结果
	aaiDetectkeywordResponse := new(AaiDetectkeywordResponse)
	err := ai.RequestAPI(aaiDetectkeywordURI, params, aaiDetectkeywordResponse)
	if err != nil {
		return nil, err
	}
	return aaiDetectkeywordResponse, nil
}

// AaiDetectkeywordForPath 关键词检索 - 音频为本地路径
// speechPath 本地待识别语音路径（时长上限15s）
// format 语音压缩格式编码，定义见 https://ai.qq.com/doc/detectword.shtml
// callbackURL 用户回调url，需用户提供，用于平台向用户通知识别结果
// keyWords 待识别关键词
// speechURL 待识别语音下载地址（时长上限15min）
func (ai *TxAi) AaiDetectkeywordForPath(speechPath, speechURL string, format AaiAudioType, callbackURL string, keyWords ...string) (*AaiDetectkeywordResponse, error) {
	body, err := ioutil.ReadFile(speechPath)
	if err != nil {
		return nil, err
	}
	speech := base64.StdEncoding.EncodeToString(body)
	return ai.AaiDetectkeywordForBase64(speech, speechURL, format, callbackURL, keyWords...)
}

// AaiDetectkeywordHandleCallback 处理回调数据
func AaiDetectkeywordHandleCallback(body []byte) (*AaiDetectkeywordResponse, *BaseResponse, error) {
	aaiDetectkeywordResponse := new(AaiDetectkeywordResponse)
	baseResponse := &BaseResponse{
		Ret: 0,
		Msg: "ok",
	}
	// 解析json数据
	err := json.Unmarshal(body, aaiDetectkeywordResponse)
	if err != nil {
		return nil, baseResponse, err
	}
	ret := aaiDetectkeywordResponse.GetRet()
	if ret != 0 {
		return nil, baseResponse, GetError(ret)
	}
	return aaiDetectkeywordResponse, baseResponse, nil
}
