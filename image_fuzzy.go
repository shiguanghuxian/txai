package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	imageFuzzyURI = "/image/image_fuzzy"
)

// ImageFuzzyForBase64 模糊图片检测 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) ImageFuzzyForBase64(image string) (*ImageFuzzyResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	imageFuzzyResponse := new(ImageFuzzyResponse)
	err := ai.RequestAPI(imageFuzzyURI, params, imageFuzzyResponse)
	if err != nil {
		return nil, err
	}
	return imageFuzzyResponse, nil
}

// ImageFuzzyForPath 模糊图片检测 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) ImageFuzzyForPath(imagePath string) (*ImageFuzzyResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.ImageFuzzyForBase64(img)
}
