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
func (ai *TxAi) AaiAsrForPath(speechPath string, format AaiAudioType, rate ...int) (*AaiAsrResponse, error) {
	body, err := ioutil.ReadFile(speechPath)
	if err != nil {
		return nil, err
	}
	speech := base64.StdEncoding.EncodeToString(body)
	return ai.AaiAsrForBase64(speech, format, rate...)
}

// AaiAsrsForBase64 语音识别-流式版（AI Lab） - 音频为base64格式
// speechChunk 待识别语音分片
// format 音频压缩格式编码，定义见 https://ai.qq.com/doc/aaiasr.shtml
// seq 语音分片所在语音流的偏移量（字节）
// length 语音分片长度（字节）
// end 是否结束分片标识，定义见 https://ai.qq.com/doc/aaiasr.shtml
// speechId 语音唯一标识
// rate 音频采样率编码，定义见下文描述
func (ai *TxAi) AaiAsrsForBase64(speechChunk string, format AaiAudioType, seq, length, end int, speechId string, rate ...int) (*AaiAsrsResponse, error) {
	params := ai.getPublicParams()
	params.Add("speech_chunk", speechChunk)
	params.Add("format", fmt.Sprint(format))
	params.Add("len", fmt.Sprint(length))
	params.Add("end", fmt.Sprint(end))
	params.Add("speech_id", speechId)
	params.Add("seq", fmt.Sprint(seq))
	if len(rate) > 0 {
		params.Add("rate", fmt.Sprint(rate[0]))
	} else {
		params.Add("rate", "16000")
	}
	sign := ai.getReqSign(params)
	params.Add("sign", sign)

	// 响应结果
	aaiAsrsResponse := new(AaiAsrsResponse)
	err := ai.RequestAPI(aaiAsrsURI, params, aaiAsrsResponse)
	if err != nil {
		return nil, err
	}
	return aaiAsrsResponse, nil
}

// AaiAsrsForPath 语音识别-流式版（AI Lab） - 音频为本地路径
// speechChunkPath 本地待识别语音分片路径
// format 音频压缩格式编码，定义见 https://ai.qq.com/doc/aaiasr.shtml
// seq 语音分片所在语音流的偏移量（字节）
// end 是否结束分片标识，定义见 https://ai.qq.com/doc/aaiasr.shtml
// speechId 语音唯一标识
// rate 音频采样率编码，定义见下文描述
func (ai *TxAi) AaiAsrsForPath(speechChunkPath string, format AaiAudioType, seq, end int, speechId string, rate ...int) (*AaiAsrsResponse, error) {
	body, err := ioutil.ReadFile(speechChunkPath)
	if err != nil {
		return nil, err
	}
	length := len(body)
	speechChunk := base64.StdEncoding.EncodeToString(body)
	return ai.AaiAsrsForBase64(speechChunk, format, seq, length, end, speechId, rate...)
}

// AaiWxasrsForBase64 语音识别-流式版（AI Lab） - 音频为base64格式
// speechChunk 待识别语音分片
// format 音频压缩格式编码，定义见 https://ai.qq.com/doc/aaiasr.shtml
// rate 音频采样率编码，定义见 https://ai.qq.com/doc/aaiasr.shtml
// bits 音频采样位数，定义见 https://ai.qq.com/doc/aaiasr.shtml
// seq 语音分片所在语音流的偏移量（字节）
// length 语音分片长度（字节）
// end 是否结束分片标识，定义见下文描述
// speechId 语音唯一标识
// contRes 是否获取中间识别结果，定义见 https://ai.qq.com/doc/aaiasr.shtml
func (ai *TxAi) AaiWxasrsForBase64(speechChunk string, format AaiAudioType, rate, bits, seq, length, end int, speechId string, contRes int) (*AaiWxasrsResponse, error) {
	params := ai.getPublicParams()
	params.Add("speech_chunk", speechChunk)
	params.Add("format", fmt.Sprint(format))
	params.Add("rate", fmt.Sprint(rate))
	params.Add("bits", fmt.Sprint(bits))
	params.Add("seq", fmt.Sprint(seq))
	params.Add("len", fmt.Sprint(length))
	params.Add("end", fmt.Sprint(end))
	params.Add("cont_res", fmt.Sprint(contRes))
	params.Add("speech_id", speechId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)

	// 响应结果
	aaiWxasrsResponse := new(AaiWxasrsResponse)
	err := ai.RequestAPI(aaiWxasrsURI, params, aaiWxasrsResponse)
	if err != nil {
		return nil, err
	}
	return aaiWxasrsResponse, nil
}

// AaiWxasrsForPath 语音识别-流式版（AI Lab） - 音频为本地路径
// speechChunkPath 本地待识别语音分片路径
// format 音频压缩格式编码，定义见 https://ai.qq.com/doc/aaiasr.shtml
// rate 音频采样率编码，定义见 https://ai.qq.com/doc/aaiasr.shtml
// bits 音频采样位数，定义见 https://ai.qq.com/doc/aaiasr.shtml
// seq 语音分片所在语音流的偏移量（字节）
// end 是否结束分片标识，定义见下文描述
// speechId 语音唯一标识
// contRes 是否获取中间识别结果，定义见 https://ai.qq.com/doc/aaiasr.shtml
func (ai *TxAi) AaiWxasrsForPath(speechChunkPath string, format AaiAudioType, rate, bits, seq, end int, speechId string, contRes int) (*AaiWxasrsResponse, error) {
	body, err := ioutil.ReadFile(speechChunkPath)
	if err != nil {
		return nil, err
	}
	length := len(body)
	speechChunk := base64.StdEncoding.EncodeToString(body)
	return ai.AaiWxasrsForBase64(speechChunk, format, rate, bits, seq, length, end, speechId, contRes)
}
