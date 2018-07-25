package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	faceFacecompareURI = "/face/face_facecompare"
)

// FaceFacecompareForBase64 人脸对比 - 图片为base64格式
// imageA 待对比人脸图片A
// imageB 待对比人脸图片B
func (ai *TxAi) FaceFacecompareForBase64(imageA, imageB string) (*FaceFacecompareResponse, error) {
	params := ai.getPublicParams()
	params.Add("image_a", imageA)
	params.Add("image_b", imageB)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceFacecompareResponse := new(FaceFacecompareResponse)
	err := ai.RequestAPI(faceFacecompareURI, params, faceFacecompareResponse)
	if err != nil {
		return nil, err
	}
	return faceFacecompareResponse, nil
}

// FaceFacecompareForPath 人脸对比 - 图片为本地路径
// imageAPath 待对比人脸图片A
// imageBPath 待对比人脸图片B
func (ai *TxAi) FaceFacecompareForPath(imageAPath, imageBPath string) (*FaceFacecompareResponse, error) {
	// 图片1
	bodyA, err := ioutil.ReadFile(imageAPath)
	if err != nil {
		return nil, err
	}
	imgA := base64.StdEncoding.EncodeToString(bodyA)
	// 图片2
	bodyB, err := ioutil.ReadFile(imageBPath)
	if err != nil {
		return nil, err
	}
	imgB := base64.StdEncoding.EncodeToString(bodyB)
	return ai.FaceFacecompareForBase64(imgA, imgB)
}
