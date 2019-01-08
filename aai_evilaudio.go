package txai

// 音频鉴黄

var (
	aaiEvilaudioURI = "/aai/aai_evilaudio"
)

// AaiEvilaudioForUrl 音频鉴黄
func (ai *TxAi) AaiEvilaudioForUrl(speechId, speechURL string, pornDetect, keywordDetect bool) (*AaiEvilaudioResponse, error) {
	params := ai.getPublicParams()
	params.Add("speech_id", speechId)
	params.Add("speech_url", speechURL)
	if pornDetect == true {
		params.Add("porn_detect", "1")
	} else {
		params.Add("porn_detect", "0")
	}
	if keywordDetect == true {
		params.Add("keyword_detect", "1")
	} else {
		params.Add("keyword_detect", "0")
	}
	sign := ai.getReqSign(params)
	params.Add("sign", sign)

	// 响应结果
	aaiEvilaudioResponse := new(AaiEvilaudioResponse)
	err := ai.RequestAPI(aaiEvilaudioURI, params, aaiEvilaudioResponse)
	if err != nil {
		return nil, err
	}
	return aaiEvilaudioResponse, nil
}
