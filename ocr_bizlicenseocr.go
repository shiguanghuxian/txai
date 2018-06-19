package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	ocrBizlicenseocrURI = "/ocr/ocr_bizlicenseocr"
)

// OcrBizlicenseocrForBase64 营业执照OCR识别 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) OcrBizlicenseocrForBase64(image string) (*OcrResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ocrResponse := new(OcrResponse)
	err := ai.RequestAPI(ocrBizlicenseocrURI, params, ocrResponse)
	if err != nil {
		return nil, err
	}
	return ocrResponse, nil
}

// OcrBizlicenseocrForPath 营业执照OCR识别 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) OcrBizlicenseocrForPath(imagePath string) (*OcrResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.OcrBizlicenseocrForBase64(img)
}
