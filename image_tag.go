package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	imageTagURI = "/image/image_tag"
)

// ImageTagForBase64 图像标签识别 - 图片为base64格式
// 识别一个图像的标签信息,对图像分类。
// image 图片base64后字符串
func (ai *TxAi) ImageTagForBase64(image string) (*ImageTagResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	imageTagResponse := new(ImageTagResponse)
	err := ai.RequestAPI(imageTagURI, params, imageTagResponse)
	if err != nil {
		return nil, err
	}
	return imageTagResponse, nil
}

// ImageTagForPath 图像标签识别 - 图片为本地路径
// 识别一个图像的标签信息,对图像分类。
// imagePath 本地图片路径
func (ai *TxAi) ImageTagForPath(imagePath string) (*ImageTagResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.ImageTagForBase64(img)
}
