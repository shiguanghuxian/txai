package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	faceDetectfaceURI = "/face/face_detectface"
)

// FaceDetectfaceForBase64 人脸检测与分析 - 图片为base64格式
// image 图片base64后字符串
// modes 检测模式，0-正常，1-大脸模式（默认1）
func (ai *TxAi) FaceDetectfaceForBase64(image string, modes ...int) (*FaceDetectfaceResponse, error) {
	mode := 1
	if len(modes) > 0 {
		mode = modes[0]
	}
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("mode", fmt.Sprint(mode))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceDetectfaceResponseResponse := new(FaceDetectfaceResponse)
	err := ai.RequestAPI(faceDetectfaceURI, params, faceDetectfaceResponseResponse)
	if err != nil {
		return nil, err
	}
	return faceDetectfaceResponseResponse, nil
}

// FaceDetectfaceForPath 人脸检测与分析 - 图片为本地路径
// imagePath 本地图片路径
// modes 检测模式，0-正常，1-大脸模式（默认1）
func (ai *TxAi) FaceDetectfaceForPath(imagePath string, modes ...int) (*FaceDetectfaceResponse, error) {
	mode := 1
	if len(modes) > 0 {
		mode = modes[0]
	}
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.FaceDetectfaceForBase64(img, mode)
}
