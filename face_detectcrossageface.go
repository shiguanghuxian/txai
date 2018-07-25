package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	faceDetectcrossagefaceURI = "/face/face_detectcrossageface"
)

// FaceDetectcrossagefaceForBase64 跨年龄人脸识别 - 图片为base64格式
// sourceImage 待比较图片
// targetImage 待比较图片
func (ai *TxAi) FaceDetectcrossagefaceForBase64(sourceImage, targetImage string) (*FaceDetectcrossagefaceResponse, error) {
	params := ai.getPublicParams()
	params.Add("source_image", sourceImage)
	params.Add("target_image", targetImage)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceDetectcrossagefaceResponse := new(FaceDetectcrossagefaceResponse)
	err := ai.RequestAPI(faceDetectcrossagefaceURI, params, faceDetectcrossagefaceResponse)
	if err != nil {
		return nil, err
	}
	return faceDetectcrossagefaceResponse, nil
}

// FaceDetectcrossagefaceForPath 跨年龄人脸识别 - 图片为本地路径
// sourceImage 待比较图片
// targetImage 待比较图片
func (ai *TxAi) FaceDetectcrossagefaceForPath(sourceImagePath, targetImagePath string) (*FaceDetectcrossagefaceResponse, error) {
	// 图片1
	bodyA, err := ioutil.ReadFile(sourceImagePath)
	if err != nil {
		return nil, err
	}
	imgA := base64.StdEncoding.EncodeToString(bodyA)
	// 图片2
	bodyB, err := ioutil.ReadFile(targetImagePath)
	if err != nil {
		return nil, err
	}
	imgB := base64.StdEncoding.EncodeToString(bodyB)
	return ai.FaceDetectcrossagefaceForBase64(imgA, imgB)
}
