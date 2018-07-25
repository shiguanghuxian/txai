package txai

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"strings"
)

var (
	faceAddfaceURI = "/face/face_addface"
)

// FaceAddfaceForBase64 增加人脸 - 图片为base64格式
// image 图片base64后字符串
// personId 指定的个体（Person）ID
// tag 备注信息
func (ai *TxAi) FaceAddfaceForBase64(images, personId, tag string) (*FaceAddfaceResponse, error) {
	if personId == "" {
		return nil, errors.New("personId can't be empty")
	}
	params := ai.getPublicParams()
	params.Add("images", images)
	params.Add("person_id", personId)
	params.Add("tag", tag)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceAddfaceResponse := new(FaceAddfaceResponse)
	err := ai.RequestAPI(faceAddfaceURI, params, faceAddfaceResponse)
	if err != nil {
		return nil, err
	}
	return faceAddfaceResponse, nil
}

// FaceAddfaceForPath 增加人脸 - 图片为本地路径
// imagePath 本地图片路径
// personId 指定的个体（Person）ID
// tag 备注信息
func (ai *TxAi) FaceAddfaceForPath(imagePath, personId, tag string) (*FaceAddfaceResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.FaceAddfaceForBase64(img, personId, tag)
}

// FaceAddMultiplefaceForPath 增加人脸 - 图片为本地路径
// imagePath 本地图片路径
// personId 指定的个体（Person）ID
// tag 备注信息
func (ai *TxAi) FaceAddMultiplefaceForPath(personId, tag string, imagePaths ...string) (*FaceAddfaceResponse, error) {
	imgs := make([]string, 0)
	for _, imagePath := range imagePaths {
		body, err := ioutil.ReadFile(imagePath)
		if err != nil {
			return nil, err
		}
		img := base64.StdEncoding.EncodeToString(body)
		imgs = append(imgs, img)
	}

	return ai.FaceAddfaceForBase64(strings.Join(imgs, "|"), personId, tag)
}
