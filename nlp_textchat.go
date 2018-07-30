package txai

var (
	nlpTextchatURI = "/nlp/nlp_textchat"
)

// NlpTextchatForText 智能闲聊 - 基础闲聊
// session 会话标识（应用内唯一）
// question 用户输入的聊天内容
func (ai *TxAi) NlpTextchatForText(session, question string) (*NlpTextchatResponse, error) {
	params := ai.getPublicParams()
	params.Add("session", session)
	params.Add("question", question)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpTextchatResponse := new(NlpTextchatResponse)
	err := ai.RequestAPI(nlpTextchatURI, params, nlpTextchatResponse)
	if err != nil {
		return nil, err
	}
	return nlpTextchatResponse, nil
}
