package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	nlpImagetranslateURI = "/nlp/nlp_imagetranslate"
)

// SceneType 图片翻译识别类型
type SceneType = string

const (
	// SceneWord 单词识别
	SceneWord SceneType = "word"
	// SceneDoc 文档识别
	SceneDoc SceneType = "doc"
)

// NlpImagetranslateForBase64 图片翻译 - 图片为base64格式
// image 图片base64后字符串
// sessionId 一次请求ID 一次请求ID（尽可能唯一，长度上限64字节）
// scene 识别类型（word-单词识别，doc-文档识别）
// source 源语言缩写，详细见下文
// target 目标语言缩写，详细见下文
func (ai *TxAi) NlpImagetranslateForBase64(image, sessionId string, scene SceneType, source, target Lang) (*NlpImagetranslateResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("session_id", sessionId)
	params.Add("scene", scene)
	params.Add("source", fmt.Sprint(source))
	params.Add("target", fmt.Sprint(target))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	nlpImagetranslateResponse := new(NlpImagetranslateResponse)
	err := ai.RequestAPI(nlpImagetranslateURI, params, nlpImagetranslateResponse)
	if err != nil {
		return nil, err
	}
	return nlpImagetranslateResponse, nil
}

// NlpImagetranslateForPath 图片翻译 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) NlpImagetranslateForPath(imagePath, sessionId, scene string, source, target Lang) (*NlpImagetranslateResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.NlpImagetranslateForBase64(img, sessionId, scene, source, target)
}
