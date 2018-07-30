package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	nlpSpeechtranslateURI = "/nlp/nlp_speechtranslate"
)

// FormatAudioType 语音翻译 语音压缩格式编码类型
type FormatAudioType int

const (
	AudioAMR  FormatAudioType = 3
	AudioSILK FormatAudioType = 4
	AudioPCM  FormatAudioType = 6
	AudioMP3  FormatAudioType = 8
	AudioAAC  FormatAudioType = 9
)

// NlpSpeechtranslateForBase64 语音翻译 - 音频为base64格式
// speechChunk 待识别语音分片
// sessionId 一次请求ID 一次请求ID（尽可能唯一，长度上限64字节）
// format 语音压缩格式编码，定义见下文描述 枚举FormatAudioType
// seq 语音分片所在语音流的偏移量（字节）
// end 是否结束分片标识，定义见下文描述
// source 源语言缩写，详细见下文
// target 目标语言缩写，详细见下文
func (ai *TxAi) NlpSpeechtranslateForBase64(speechChunk, sessionId string, format FormatAudioType, seq, end uint, source, target Lang) (*NlpSpeechtranslateResponse, error) {
	params := ai.getPublicParams()
	params.Add("format", fmt.Sprint(format))
	params.Add("seq", fmt.Sprint(seq))
	params.Add("end", fmt.Sprint(end))
	params.Add("session_id", sessionId)
	params.Add("speech_chunk", speechChunk)
	params.Add("source", fmt.Sprint(source))
	params.Add("target", fmt.Sprint(target))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpSpeechtranslateResponse := new(NlpSpeechtranslateResponse)
	err := ai.RequestAPI(nlpSpeechtranslateURI, params, nlpSpeechtranslateResponse)
	if err != nil {
		return nil, err
	}
	return nlpSpeechtranslateResponse, nil
}

// NlpSpeechtranslateForPath 语音翻译 - 音频为本地路径
// speechChunkPath 本地图片路径
// sessionId 一次请求ID 一次请求ID（尽可能唯一，长度上限64字节）
// format 语音压缩格式编码，定义见下文描述 枚举FormatAudioType
// seq 语音分片所在语音流的偏移量（字节）
// end 是否结束分片标识，定义见下文描述
// source 源语言缩写，详细见下文
// target 目标语言缩写，详细见下文
func (ai *TxAi) NlpSpeechtranslateForPath(speechChunkPath, sessionId string, format FormatAudioType, seq, end uint, source, target Lang) (*NlpSpeechtranslateResponse, error) {
	body, err := ioutil.ReadFile(speechChunkPath)
	if err != nil {
		return nil, err
	}
	speech := base64.StdEncoding.EncodeToString(body)
	return ai.NlpSpeechtranslateForBase64(speech, sessionId, format, seq, end, source, target)
}
