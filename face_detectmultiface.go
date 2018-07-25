package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	faceDetectmultifaceURI = "/face/face_detectmultiface"
)

// FaceDetectmultifaceForBase64 多人脸检测 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) FaceDetectmultifaceForBase64(image string) (*FaceDetectmultifaceResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceDetectmultifaceResponse := new(FaceDetectmultifaceResponse)
	err := ai.RequestAPI(faceDetectmultifaceURI, params, faceDetectmultifaceResponse)
	if err != nil {
		return nil, err
	}
	return faceDetectmultifaceResponse, nil
}

// FaceDetectmultifaceForPath 多人脸检测 - 图片为本地路径
// imagePath 本地图片路径
func (ai *TxAi) FaceDetectmultifaceForPath(imagePath string) (*FaceDetectmultifaceResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.FaceDetectmultifaceForBase64(img)
}
