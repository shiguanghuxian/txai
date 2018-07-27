package txai

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
)

var (
	visionObjectrURI = "/vision/vision_objectr"
)

// VisionObjectrForBase64 物体识别 - 图片为base64格式
// 对图片进行物体识别，快速找出图片中包含的物体信息
// image 图片base64后字符串
// format 图片格式，暂时只支持1 JPG格式（image/jpeg）
// topk 返回结果个数（已按置信度倒排）
func (ai *TxAi) VisionObjectrForBase64(image string, format, topk int) (*VisionObjectrResponse, error) {
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
	visionObjectrResponse := new(VisionObjectrResponse)
	err := ai.RequestAPI(visionObjectrURI, params, visionObjectrResponse)
	if err != nil {
		return nil, err
	}
	// 添加场景名
	for _, v := range visionObjectrResponse.Data.ObjectList {
		v.LabelName = VisionObjectrNames[v.LabelID]
	}
	return visionObjectrResponse, nil
}

// VisionObjectrForPath 物体识别 - 图片为本地路径
// 对图片进行物体识别，快速找出图片中包含的物体信息
// imagePath 本地图片路径
// format 图片格式，暂时只支持1 JPG格式（image/jpeg）
// topk 返回结果个数（已按置信度倒排）
func (ai *TxAi) VisionObjectrForPath(imagePath string, format, topk int) (*VisionObjectrResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.VisionObjectrForBase64(img, format, topk)
}
