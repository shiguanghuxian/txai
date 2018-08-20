package txai

import "github.com/axgle/mahonia"

var (
	nlpWordsegURI = "/nlp/nlp_wordseg"
	nlpWordposURI = "/nlp/nlp_wordpos"
	nlpWordnerURI = "/nlp/nlp_wordner"
	nlpWordsynURI = "/nlp/nlp_wordsyn"
)

// NlpWordsegForText 分词 - 对文本进行智能分词识别，支持基础词与混排词粒度
// text 待分析文本
func (ai *TxAi) NlpWordsegForText(text string) (*NlpWordsegResponse, error) {
	// 转gbk
	enc := mahonia.NewEncoder("gbk")
	text = enc.ConvertString(text)
	// 组织参数
	params := ai.getPublicParams()
	params.Add("text", text)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)

	// 响应结果
	nlpWordsegResponse := new(NlpWordsegResponse)
	err := ai.RequestAPI(nlpWordsegURI, params, nlpWordsegResponse)
	if err != nil {
		return nil, err
	}
	return nlpWordsegResponse, nil
}

// NlpWordposForText 词性标注 - 对文本进行分词，同时为每个分词标注正确的词性
// text 待分析文本
func (ai *TxAi) NlpWordposForText(text string) (*NlpWordposResponse, error) {
	// 转gbk
	enc := mahonia.NewEncoder("gbk")
	text = enc.ConvertString(text)
	// 组织参数
	params := ai.getPublicParams()
	params.Add("text", text)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpWordposResponse := new(NlpWordposResponse)
	err := ai.RequestAPI(nlpWordposURI, params, nlpWordposResponse)
	if err != nil {
		return nil, err
	}
	// 处理词性名
	for _, v := range nlpWordposResponse.Data.BaseTokens {
		v.PosCodeName = NlpWordposNames[v.PosCode]
	}
	for _, v := range nlpWordposResponse.Data.MixTokens {
		v.PosCodeName = NlpWordposNames[v.PosCode]
	}
	return nlpWordposResponse, nil
}

// NlpWordnerForText 专有名词识别 - 对文本进行专有名词的分词识别，找出文本中的专有名词
// text 待分析文本
func (ai *TxAi) NlpWordnerForText(text string) (*NlpWordnerResponse, error) {
	// 转gbk
	enc := mahonia.NewEncoder("gbk")
	text = enc.ConvertString(text)
	// 组织参数
	params := ai.getPublicParams()
	params.Add("text", text)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpWordnerResponse := new(NlpWordnerResponse)
	err := ai.RequestAPI(nlpWordnerURI, params, nlpWordnerResponse)
	if err != nil {
		return nil, err
	}
	// 获取专有名词编码
	for _, v := range nlpWordnerResponse.Data.NerTokens {
		v.TypeNames = make([]string, 0)
		for _, vv := range v.Types {
			v.TypeNames = append(v.TypeNames, NlpWordnerNames[vv])
		}
	}
	return nlpWordnerResponse, nil
}

// NlpWordsynForText 同义词识别 - 识别文本中存在同义词的分词，并返回相应的同义词
// text 待分析文本
func (ai *TxAi) NlpWordsynForText(text string) (*NlpWordsynResponse, error) {
	// 转gbk
	enc := mahonia.NewEncoder("gbk")
	text = enc.ConvertString(text)
	// 组织参数
	params := ai.getPublicParams()
	params.Add("text", text)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpWordsynResponse := new(NlpWordsynResponse)
	err := ai.RequestAPI(nlpWordsynURI, params, nlpWordsynResponse)
	if err != nil {
		return nil, err
	}
	return nlpWordsynResponse, nil
}
