package txai

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
)

var (
	visionScenerURI = "/vision/vision_scener"
)

// VisionScenerForBase64 场景识别 - 图片为base64格式
// 对图片进行场景识别，快速找出图片中包含的场景信息
// image 图片base64后字符串
// format 图片格式，暂时只支持1 JPG格式（image/jpeg）
// topk 返回结果个数（已按置信度倒排）
func (ai *TxAi) VisionScenerForBase64(image string, format, topk int) (*VisionScenerResponse, error) {
	if topk < 1 || topk > 5 {
		return nil, errors.New("topk range from 1 to 5")
	}
	if format != 1 {
		return nil, errors.New("The picture format can only be jpg [format=1]")
	}

	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("format", fmt.Sprint(format))
	params.Add("topk", fmt.Sprint(topk))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	visionScenerResponse := new(VisionScenerResponse)
	err := ai.RequestAPI(visionScenerURI, params, visionScenerResponse)
	if err != nil {
		return nil, err
	}
	// 添加场景名
	for _, v := range visionScenerResponse.Data.SceneList {
		v.LabelName = VisionScenerNames[v.LabelID]
	}
	return visionScenerResponse, nil
}

// VisionScenerForPath 场景识别 - 图片为本地路径
// 对图片进行场景识别，快速找出图片中包含的场景信息
// imagePath 本地图片路径
// format 图片格式，暂时只支持1 JPG格式（image/jpeg）
// topk 返回结果个数（已按置信度倒排）
func (ai *TxAi) VisionScenerForPath(imagePath string, format, topk int) (*VisionScenerResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.VisionScenerForBase64(img, format, topk)
}
