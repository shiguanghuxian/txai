package txai

import "strings"

var (
	nlpTextdetectURI = "/nlp/nlp_textdetect"
)

// NlpTextdetectForText 语种识别 - 对文本进行翻译，支持多种语言之间互译
// text 待翻译文本
// force 是否强制从候选语言中选择（只对二选一有效）
// candidateLangs 备选语言缩写，详细见 https://ai.qq.com/doc/textdetect.shtml
func (ai *TxAi) NlpTextdetectForText(text string, force bool, candidateLangs ...Lang) (*NlpTextdetectResponse, error) {
	params := ai.getPublicParams()
	params.Add("text", text)
	if force == true {
		params.Add("force", "1")
	} else {
		params.Add("force", "0")
	}
	params.Add("candidate_langs", strings.Join(candidateLangs, "|"))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpTextdetectResponse := new(NlpTextdetectResponse)
	err := ai.RequestAPI(nlpTextdetectURI, params, nlpTextdetectResponse)
	if err != nil {
		return nil, err
	}
	return nlpTextdetectResponse, nil
}
