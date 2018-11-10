package txai

// 音频鉴黄

var (
	aaiEvilaudioURI = "/aai/aai_evilaudio"
)

// AaiEvilaudioForUrl 音频鉴黄
func (ai *TxAi) AaiEvilaudioForUrl(speechId, speechURL string) (*AaiEvilaudioResponse, error) {
	params := ai.getPublicParams()
	params.Add("speech_id", speechId)
	params.Add("speech_url", speechURL)
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
