package txai

var (
	nlpTextpolarURI = "/nlp/nlp_textpolar"
)

// NlpTextpolarForText 情感分析 - 对文本进行情感分析，快速判断情感倾向（正面或负面）
// text 待分析文本
func (ai *TxAi) NlpTextpolarForText(text string) (*NlpTextpolarResponse, error) {
	params := ai.getPublicParams()
	params.Add("text", text)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpTextpolarResponse := new(NlpTextpolarResponse)
	err := ai.RequestAPI(nlpTextpolarURI, params, nlpTextpolarResponse)
	if err != nil {
		return nil, err
	}
	switch nlpTextpolarResponse.Data.Polar {
	case -1:
		nlpTextpolarResponse.Data.PolarName = "负面"
		break
	case 0:
		nlpTextpolarResponse.Data.PolarName = "中性"
		break
	case 1:
		nlpTextpolarResponse.Data.PolarName = "正面"
		break
	}
	return nlpTextpolarResponse, nil
}
