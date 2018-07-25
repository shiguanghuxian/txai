package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	ocrHandwritingocrURI = "/ocr/ocr_handwritingocr"
)

// OcrHandwritingocrForBase64 手写体OCR - 检测和识别图像上面手写体的字段信息 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) OcrHandwritingocrForBase64(image string) (*OcrResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ocrResponse := new(OcrResponse)
	err := ai.RequestAPI(ocrHandwritingocrURI, params, ocrResponse)
	if err != nil {
		return nil, err
	}
	return ocrResponse, nil
}

// OcrHandwritingocrForPath 手写体OCR - 检测和识别图像上面手写体的字段信息 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) OcrHandwritingocrForPath(imagePath string) (*OcrResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.OcrHandwritingocrForBase64(img)
}

// OcrHandwritingocrForURL 手写体OCR - 检测和识别图像上面手写体的字段信息 - 图片为url地址
// imageURL 待识别图片url
func (ai *TxAi) OcrHandwritingocrForURL(imageURL string) (*OcrResponse, error) {
	params := ai.getPublicParams()
	params.Add("image_url", imageURL)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ocrResponse := new(OcrResponse)
	err := ai.RequestAPI(ocrHandwritingocrURI, params, ocrResponse)
	if err != nil {
		return nil, err
	}
	return ocrResponse, nil
}
