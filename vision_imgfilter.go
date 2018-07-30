package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	visionImgfilterURI = "/vision/vision_imgfilter"
)

// VisionImgfilterForBase64 图片滤镜（AI Lab） - 图片为base64格式
// image 图片base64后字符串
// session_id 一次请求ID 尽可能唯一，长度上限64字节
// filter 变妆编码，定义见 https://ai.qq.com/doc/ptuimgfilter.shtml
func (ai *TxAi) VisionImgfilterForBase64(image, sessionId string, filter int) (*PtuResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("session_id", sessionId)
	params.Add("filter", fmt.Sprint(filter))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ptuResponse := new(PtuResponse)
	err := ai.RequestAPI(visionImgfilterURI, params, ptuResponse)
	if err != nil {
		return nil, err
	}
	return ptuResponse, nil
}

// VisionImgfilterForPath 图片滤镜（AI Lab） - 图片为本地路径
// imagePath 本地图片路径
// session_id 一次请求ID 尽可能唯一，长度上限64字节
// filter 变妆编码，定义见 https://ai.qq.com/doc/ptuimgfilter.shtml
func (ai *TxAi) VisionImgfilterForPath(imagePath, sessionId string, filter int) (*PtuResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.VisionImgfilterForBase64(img, sessionId, filter)
}
