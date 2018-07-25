package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	ocrPlateocrURI = "/ocr/ocr_plateocr"
)

// OcrPlateocrForBase64 车牌OCR - 识别车牌上面的字段信息 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) OcrPlateocrForBase64(image string) (*OcrResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ocrResponse := new(OcrResponse)
	err := ai.RequestAPI(ocrPlateocrURI, params, ocrResponse)
	if err != nil {
		return nil, err
	}
	return ocrResponse, nil
}

// OcrPlateocrForPath 车牌OCR - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) OcrPlateocrForPath(imagePath string) (*OcrResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.OcrPlateocrForBase64(img)
}

// OcrPlateocrForURL 车牌OCR - 识别车牌上面的字段信息 - 图片为url地址
// imageURL 待识别图片url
func (ai *TxAi) OcrPlateocrForURL(imageURL string) (*OcrResponse, error) {
	params := ai.getPublicParams()
	params.Add("image_url", imageURL)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ocrResponse := new(OcrResponse)
	err := ai.RequestAPI(ocrPlateocrURI, params, ocrResponse)
	if err != nil {
		return nil, err
	}
	return ocrResponse, nil
}
