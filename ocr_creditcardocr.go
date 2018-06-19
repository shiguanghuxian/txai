package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	ocrCreditcardocrURI = "/ocr/ocr_creditcardocr"
)

// OcrCreditcardocrForBase64 银行卡OCR识别 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) OcrCreditcardocrForBase64(image string) (*OcrResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ocrResponse := new(OcrResponse)
	err := ai.RequestAPI(ocrCreditcardocrURI, params, ocrResponse)
	if err != nil {
		return nil, err
	}
	return ocrResponse, nil
}

// OcrCreditcardocrForPath 银行卡OCR识别 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) OcrCreditcardocrForPath(imagePath string) (*OcrResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.OcrCreditcardocrForBase64(img)
}
