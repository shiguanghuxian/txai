package txai

import "fmt"

var (
	nlpTexttranslateURI = "/nlp/nlp_texttranslate"
)

// Lang 语言
type Lang = string

const (
	// Zh 中文
	Zh Lang = "zh"
	// En 英文
	En Lang = "en"
	// Jp 日文
	Jp Lang = "jp"
	// Kr 韩文
	Kr Lang = "kr"
	// Fr 法文
	Fr Lang = "fr"
	// Es 西班牙文
	Es Lang = "es"
	// It 意大利文
	It Lang = "it"
	// De 德文
	De Lang = "de"
	// Tr 土耳其文
	Tr Lang = "tr"
	// Ru 俄文
	Ru Lang = "ru"
	// Pt 葡萄牙文
	Pt Lang = "pt"
	// Vi 越南文
	Vi Lang = "vi"
	// Id 印度尼西亚文
	Id Lang = "id"
	// Ms 马来西亚文
	Ms Lang = "ms"
	// Th 泰文
	Th Lang = "th"
	// LangAuto 自动识别（中英互译）
	Auto Lang = "auto"
)

// NlpTexttranslateForText 文本翻译（翻译君） - 对文本进行翻译，支持多种语言之间互译
// text 待翻译文本
// langType 翻译类型，默认为0，详细见下文
func (ai *TxAi) NlpTexttranslateForText(text string, source, target Lang) (*NlpTexttranslateResponse, error) {
	params := ai.getPublicParams()
	params.Add("text", text)
	params.Add("source", fmt.Sprint(source))
	params.Add("target", fmt.Sprint(target))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpTexttranslateResponse := new(NlpTexttranslateResponse)
	err := ai.RequestAPI(nlpTexttranslateURI, params, nlpTexttranslateResponse)
	if err != nil {
		return nil, err
	}
	return nlpTexttranslateResponse, nil
}
