package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	imageFoodURI = "/image/image_food"
)

// ImageFoodForBase64 美食图片识别 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) ImageFoodForBase64(image string) (*ImageFoodResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	imageFoodResponse := new(ImageFoodResponse)
	err := ai.RequestAPI(imageFoodURI, params, imageFoodResponse)
	if err != nil {
		return nil, err
	}
	return imageFoodResponse, nil
}

// ImageFoodForPath 美食图片识别 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) ImageFoodForPath(imagePath string) (*ImageFoodResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.ImageFoodForBase64(img)
}
