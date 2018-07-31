package txai

// 语音识别

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	aaiWxasrlongURI = "/aai/aai_wxasrlong"
)

// AaiWxasrlongForBase64 长语音识别 - 音频为base64格式
// speech 待识别语音（时长上限15s）
// format 语音压缩格式编码，定义见 https://ai.qq.com/doc/wxasrlong.shtml
// callbackURL 用户回调url，需用户提供，用于平台向用户通知识别结果，详见下文描述
// speechURL 待识别语音下载地址（时长上限15min）
func (ai *TxAi) AaiWxasrlongForBase64(speech string, format AaiAudioType, callbackURL, speechURL string) (*AaiWxasrlongResponse, error) {
	params := ai.getPublicParams()
	if speech != "" {
		params.Add("speech", speech)
	}
	if speechURL != "" {
		params.Add("speech_url", speechURL)
	}
	params.Add("format", fmt.Sprint(format))
	params.Add("callback_url", callbackURL)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)

	// 响应结果
	aaiWxasrlongResponse := new(AaiWxasrlongResponse)
	err := ai.RequestAPI(aaiWxasrlongURI, params, aaiWxasrlongResponse)
	if err != nil {
		return nil, err
	}
	return aaiWxasrlongResponse, nil
}

// AaiWxasrlongForPath 长语音识别 - 音频为本地路径
// speechPath 本地待识别语音路径（时长上限15s）
// format 语音压缩格式编码，定义见 https://ai.qq.com/doc/wxasrlong.shtml
// callbackURL 用户回调url，需用户提供，用于平台向用户通知识别结果，详见下文描述
// speechURL 待识别语音下载地址（时长上限15min）
func (ai *TxAi) AaiWxasrlongForPath(speechPath string, format AaiAudioType, callbackURL, speechURL string) (*AaiWxasrlongResponse, error) {
	body, err := ioutil.ReadFile(speechPath)
	if err != nil {
		return nil, err
	}
	speech := base64.StdEncoding.EncodeToString(body)
	return ai.AaiWxasrlongForBase64(speech, format, callbackURL, speechURL)
}

// AaiWxasrlongHandleCallback 处理回调数据
func AaiWxasrlongHandleCallback(body []byte) (*AaiWxasrlongResponse, *BaseResponse, error) {
	aaiWxasrlongResponse := new(AaiWxasrlongResponse)
	baseResponse := &BaseResponse{
		Ret: 0,
		Msg: "ok",
	}
	// 解析json数据
	err := json.Unmarshal(body, aaiWxasrlongResponse)
	if err != nil {
		return nil, baseResponse, err
	}
	ret := aaiWxasrlongResponse.GetRet()
	if ret != 0 {
		return nil, baseResponse, GetError(ret)
	}
	return aaiWxasrlongResponse, baseResponse, nil
}
