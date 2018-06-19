package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	visionPornURI = "/vision/vision_porn"
)

// VisionPornForBase64 智能鉴黄 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) VisionPornForBase64(image string) (*VisionPornResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	visionPornResponse := new(VisionPornResponse)
	err := ai.RequestAPI(visionPornURI, params, visionPornResponse)
	if err != nil {
		return nil, err
	}
	// 计算图片级别
	hot := 0
	normal := 0
	for _, v := range visionPornResponse.Data.TagList {
		if v.TagName == "normal_hot_porn" && v.TagConfidence > 83 {
			visionPornResponse.Level = 2
			break
		} else if v.TagName == "hot" {
			hot = v.TagConfidence
		} else if v.TagName == "normal" {
			normal = v.TagConfidence
		}
	}
	// 是否是性感图片
	if visionPornResponse.Level != 2 && hot > normal {
		visionPornResponse.Level = 1
	}
	return visionPornResponse, nil
}

// VisionPornForPath 智能鉴黄 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) VisionPornForPath(imagePath string) (*VisionPornResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.VisionPornForBase64(img)
}
