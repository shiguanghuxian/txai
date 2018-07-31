package txai

import "fmt"

var (
	aaiTtsURI = "/aai/aai_tts"
	aaiTtaURI = "/aai/aai_tta"
)

// AaiTtsSpeakerType 语音合成（AI Lab） 语音发音人编码
type AaiTtsSpeakerType = int

const (
	// AaiTtsSpeaker1 普通话男声
	AaiTtsSpeaker1 AaiTtsSpeakerType = 1
	// AaiTtsSpeaker5 静琪女声
	AaiTtsSpeaker5 AaiTtsSpeakerType = 5
	// AaiTtsSpeaker6 欢馨女声
	AaiTtsSpeaker6 AaiTtsSpeakerType = 6
	// AaiTtsSpeaker7 碧萱女声
	AaiTtsSpeaker7 AaiTtsSpeakerType = 7
)

// AaiTtsFormatType 语音合成（AI Lab） 合成语音格式编码
type AaiTtsFormatType = int

const (
	// AaiTtsFormatPCM PCM
	AaiTtsFormatPCM AaiTtsFormatType = 1
	// AaiTtsFormatWAV WAV
	AaiTtsFormatWAV AaiTtsFormatType = 2
	// AaiTtsFormatMP3 MP3
	AaiTtsFormatMP3 AaiTtsFormatType = 3
)

// AaiTtsForText 语音合成（AI Lab）
// text 待合成文本
// speaker 语音发音人编码，定义见 https://ai.qq.com/doc/aaitts.shtml
// format 合成语音格式编码
// volume 合成语音音量，取值范围[-10, 10]，如-10表示音量相对默认值小10dB，0表示默认音量，10表示音量相对默认值大10dB
// speed 合成语音语速，默认100
// aht 合成语音降低/升高半音个数，即改变音高，默认0
// apc 控制频谱翘曲的程度，改变说话人的音色，默认58
func (ai *TxAi) AaiTtsForText(text string, speaker AaiTtsSpeakerType, format AaiTtsFormatType, volume, speed, aht, apc int) (*AaiTtsResponse, error) {
	if volume < -10 || volume > 10 {
		volume = 0
	}
	if speed < 50 || speed > 100 {
		speed = 100
	}
	if aht < -24 || aht > 24 {
		aht = 0
	}
	if apc < 0 || apc > 100 {
		apc = 58
	}
	params := ai.getPublicParams()
	params.Add("text", text)
	params.Add("speaker", fmt.Sprint(speaker))
	params.Add("format", fmt.Sprint(format))
	params.Add("volume", fmt.Sprint(volume))
	params.Add("speed", fmt.Sprint(speed))
	params.Add("aht", fmt.Sprint(aht))
	params.Add("apc", fmt.Sprint(apc))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	aaiTtsResponse := new(AaiTtsResponse)
	err := ai.RequestAPI(aaiTtsURI, params, aaiTtsResponse)
	if err != nil {
		return nil, err
	}
	return aaiTtsResponse, nil
}

// AaiTtaModelType 语音合成（优图） 模型编码
type AaiTtaModelType = int

const (
	// AaiTtaModelType0 女生
	AaiTtaModelType0 AaiTtaModelType = 0
	// AaiTtaModelType1 女生纯英文
	AaiTtaModelType1 AaiTtaModelType = 1
	// AaiTtaModelType2 男生
	AaiTtaModelType2 AaiTtaModelType = 2
	// AaiTtaModelType6 喜道公子
	AaiTtaModelType6 AaiTtaModelType = 6
)

// AaiTtaForText 语音合成（优图）
// text 待合成文本
// modelType 发音模型，默认为0，定义见 https://ai.qq.com/doc/aaitts.shtml
// speed 语速，默认为0 0.6倍速-2,0.8倍速-1,正常速度0,1.2倍速1,1.5倍速2
func (ai *TxAi) AaiTtaForText(text string, modelType AaiTtaModelType, speed int) (*AaiTtaResponse, error) {
	if speed < -2 || speed > 2 {
		speed = 0
	}
	params := ai.getPublicParams()
	params.Add("text", text)
	params.Add("model_type", fmt.Sprint(modelType))
	params.Add("speed", fmt.Sprint(speed))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	aaiTtaResponse := new(AaiTtaResponse)
	err := ai.RequestAPI(aaiTtaURI, params, aaiTtaResponse)
	if err != nil {
		return nil, err
	}
	return aaiTtaResponse, nil
}
