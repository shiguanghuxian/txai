package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	visionImgtotextURI = "/vision/vision_imgtotext"
)

// VisionImgtotextForBase64 看图说话 - 图片为base64格式
// image 图片base64后字符串
// sessionId 一次请求ID 尽可能唯一，长度上限64字节
func (ai *TxAi) VisionImgtotextForBase64(image, sessionId string) (*VisionImgtotextResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("session_id", sessionId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	visionImgtotextResponse := new(VisionImgtotextResponse)
	err := ai.RequestAPI(visionImgtotextURI, params, visionImgtotextResponse)
	if err != nil {
		return nil, err
	}
	return visionImgtotextResponse, nil
}

// VisionImgtotextForPath 行驶证驾驶证OCR识别 - 图片为本地路径
// imagePath 本地图片路径
// sessionId 一次请求ID 尽可能唯一，长度上限64字节
func (ai *TxAi) VisionImgtotextForPath(imagePath, sessionId string) (*VisionImgtotextResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.VisionImgtotextForBase64(img, sessionId)
}
