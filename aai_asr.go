package txai

// 语音识别

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	aaiAsrURI    = "/aai/aai_asr"
	aaiAsrsURI   = "/aai/aai_asrs"
	aaiWxasrsURI = "/aai/aai_wxasrs"
)

// AaiAudioType 语言识别 语音压缩格式编码
type AaiAudioType = int

const (
	AaiAudioPCM   AaiAudioType = 1
	AaiAudioWAV   AaiAudioType = 2
	AaiAudioAMR   AaiAudioType = 3
	AaiAudioSILK  AaiAudioType = 4
	AaiAudioSPEEX AaiAudioType = 5
	AaiAudioMP3   AaiAudioType = 8
)

// AaiAsrForBase64 语音识别-echo版 - 音频为base64格式
// speech 待识别语音（时长上限15s）
// format 语音压缩格式编码，定义见下文描述
// rate 语音采样率编码，定义见下文描述，（不传）默认即16KHz
func (ai *TxAi) AaiAsrForBase64(speech string, format AaiAudioType, rate ...int) (*AaiAsrResponse, error) {
	params := ai.getPublicParams()
	params.Add("speech", speech)
	params.Add("format", fmt.Sprint(format))
	if len(rate) > 0 {
		params.Add("rate", fmt.Sprint(rate[0]))
	} else {
		params.Add("rate", "16000")
	}
	sign := ai.getReqSign(params)
	params.Add("sign", sign)

	// 响应结果
	aaiAsrResponse := new(AaiAsrResponse)
	err := ai.RequestAPI(aaiAsrURI, params, aaiAsrResponse)
	if err != nil {
		return nil, err
	}
	return aaiAsrResponse, nil
}

// AaiAsrForPath 语音识别-echo版 - 音频为本地路径
// imagePath 本地图片路径
// speech 本地语音路径（时长上限15s）
// format 语音压缩格式编码，定义见下文描述
// rate 语音采样率编码，定义见下文描述，（不传）默认即16KHz
func (ai *TxAi) AaiAsrForPath(speechPath string, format int, rate ...int) (*AaiAsrResponse, error) {
	body, err := ioutil.ReadFile(speechPath)
	if err != nil {
		return nil, err
	}
	speech := base64.StdEncoding.EncodeToString(body)
	return ai.AaiAsrForBase64(speech, format, rate...)
}
