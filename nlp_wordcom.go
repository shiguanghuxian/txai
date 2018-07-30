package txai

var (
	nlpWordcomURI = "/nlp/nlp_wordcom"
)

// NlpWordcomForText 语义解析 - 对文本进行意图识别，快速找出意图及上下文成分
// text 待分析文本
func (ai *TxAi) NlpWordcomForText(text string) (*NlpWordcomResponse, error) {
	params := ai.getPublicParams()
	params.Add("text", text)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpWordcomResponse := new(NlpWordcomResponse)
	err := ai.RequestAPI(nlpWordcomURI, params, nlpWordcomResponse)
	if err != nil {
		return nil, err
	}
	nlpWordcomResponse.Data.IntentName = NlpWordcomIntentNames[nlpWordcomResponse.Data.Intent]
	for _, v := range nlpWordcomResponse.Data.ComTokens {
		v.ComTypeName = NlpWordcomComTypeNames[v.ComType]
	}

	return nlpWordcomResponse, nil
}
