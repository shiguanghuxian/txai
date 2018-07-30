package txai

import "fmt"

var (
	nlpTexttransURI = "/nlp/nlp_texttrans"
)

// LangType 翻译类型
type LangType uint

const (
	// LangAuto 自动识别（中英文互转）
	LangAuto LangType = 0
	// LangZhToEn 中文翻译成英文
	LangZhToEn LangType = 1
	// LangEnToZh 英文翻译成中文
	LangEnToZh LangType = 1
	// LangZhToEs 中文翻译成西班牙文
	LangZhToEs LangType = 3
	// LangEsToZh 西班牙文翻译成中文
	LangEsToZh LangType = 4
	// LangZhToFr 中文翻译成法文
	LangZhToFr LangType = 5
	// LangFrToZh 法文翻译成中文
	LangFrToZh LangType = 6
	// LangZhToVi 英文翻译成越南语
	LangZhToVi LangType = 7
	// LangViToZh 越南语翻译成英文
	LangViToZh LangType = 8
	// LangZhToZhYY 中文翻译成粤语
	LangZhToZhYY LangType = 9
	// LangZhYYToZh 中文翻译成粤语
	LangZhYYToZh LangType = 10
	// LangZhToKr 中文翻译成韩文
	LangZhToKr LangType = 11
	// LangEnToDe 英文翻译成德语
	LangEnToDe LangType = 13
	// LangDeToEn 德语翻译成英文
	LangDeToEn LangType = 14
	// LangZhToJp 中文翻译成日文
	LangZhToJp LangType = 15
	// LangJpToZh 日文翻译成中文
	LangJpToZh LangType = 16
)

// NlpTexttransForText 文本翻译（AI Lab） - 对文本进行翻译，支持多种语言之间互译
// text 待翻译文本
// langType 翻译类型，默认为0，详细见下文
func (ai *TxAi) NlpTexttransForText(text string, langTypes ...LangType) (*NlpTexttransResponse, error) {
	langType := LangAuto
	if len(langTypes) > 0 {
		langType = langTypes[0]
	}
	params := ai.getPublicParams()
	params.Add("text", text)
	params.Add("type", fmt.Sprint(langType))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpTexttransResponse := new(NlpTexttransResponse)
	err := ai.RequestAPI(nlpTexttransURI, params, nlpTexttransResponse)
	if err != nil {
		return nil, err
	}
	return nlpTexttransResponse, nil
}
