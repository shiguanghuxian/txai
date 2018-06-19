package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	ocrIdcardocrURL = "/ocr/ocr_idcardocr"
)

// OcrIdcardocrForBase64 身份证OCR识别 - 图片为base64格式
// cardType 身份证图片类型，0-正面，1-反面
// image 图片base64后字符串
func (ai *TxAi) OcrIdcardocrForBase64(cardType int, image string) (*OcrIDCardResponse, error) {
	params := ai.getPublicParams()
	params.Add("card_type", fmt.Sprint(cardType))
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	idCardResponse := new(OcrIDCardResponse)
	err := ai.RequestAPI(ocrIdcardocrURL, params, idCardResponse)
	if err != nil {
		return nil, err
	}
	return idCardResponse, nil
}

// OcrIdcardocrForPath 身份证OCR识别 - 图片为本地路径
// cardType 身份证图片类型，0-正面，1-反面
// imagePath 本地图片路径
func (ai *TxAi) OcrIdcardocrForPath(cardType int, imagePath string) (*OcrIDCardResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.OcrIdcardocrForBase64(cardType, img)
}
