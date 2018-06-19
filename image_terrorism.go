package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	imageTerrorismURI = "/image/image_terrorism"
)

// ImageTerrorismForBase64 营业执照OCR识别 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) ImageTerrorismForBase64(image string) (*ImageTerrorismResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	imageTerrorism := new(ImageTerrorismResponse)
	err := ai.RequestAPI(imageTerrorismURI, params, imageTerrorism)
	if err != nil {
		return nil, err
	}
	// 各种属性赋值
	for _, v := range imageTerrorism.Data.TagList {
		switch v.TagName {
		case "terrorists":
			imageTerrorism.Terrorists = v.TagConfidence
			break
		case "knife":
			imageTerrorism.Knife = v.TagConfidence
			break
		case "guns":
			imageTerrorism.Guns = v.TagConfidence
			break
		case "blood":
			imageTerrorism.Blood = v.TagConfidence
			break
		case "fire":
			imageTerrorism.Fire = v.TagConfidence
			break
		case "flag":
			imageTerrorism.Flag = v.TagConfidence
			break
		case "crowd":
			imageTerrorism.Crowd = v.TagConfidence
			break
		case "other":
			imageTerrorism.Other = v.TagConfidence
			break
		}
	}
	return imageTerrorism, nil
}

// ImageTerrorismForPath 营业执照OCR识别 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) ImageTerrorismForPath(imagePath string) (*ImageTerrorismResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.ImageTerrorismForBase64(img)
}
