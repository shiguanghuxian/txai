package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	ocrDriverlicenseocrURI = "/ocr/ocr_driverlicenseocr"
)

// OcrDriverlicenseocrForBase64 行驶证驾驶证OCR识别 - 图片为base64格式
// t 识别类型，0-行驶证识别，1-驾驶证识别
// image 图片base64后字符串
func (ai *TxAi) OcrDriverlicenseocrForBase64(t int, image string) (*OcrResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("type", fmt.Sprint(t))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ocrResponse := new(OcrResponse)
	err := ai.RequestAPI(ocrDriverlicenseocrURI, params, ocrResponse)
	if err != nil {
		return nil, err
	}
	return ocrResponse, nil
}

// OcrDriverlicenseocrForPath 行驶证驾驶证OCR识别 - 图片为本地路径
// t 识别类型，0-行驶证识别，1-驾驶证识别
// imagePath 本地图片路径
func (ai *TxAi) OcrDriverlicenseocrForPath(t int, imagePath string) (*OcrResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.OcrDriverlicenseocrForBase64(t, img)
}
