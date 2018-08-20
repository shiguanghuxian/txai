package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	faceFaceshapeURI = "/face/face_faceshape"
)

// FaceFaceshapeForBase64 五官定位 - 图片为base64格式
// image 图片base64后字符串
// modes 检测模式，0-正常，1-大脸模式（默认1）
func (ai *TxAi) FaceFaceshapeForBase64(image string, modes ...int) (*FaceFaceshapeResponse, error) {
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
	faceFaceshapeResponse := new(FaceFaceshapeResponse)
	err := ai.RequestAPI(faceFaceshapeURI, params, faceFaceshapeResponse)
	if err != nil {
		return nil, err
	}
	return faceFaceshapeResponse, nil
}

// FaceFaceshapeForPath 五官定位 - 图片为本地路径
// imagePath 本地图片路径
// modes 检测模式，0-正常，1-大脸模式（默认1）
func (ai *TxAi) FaceFaceshapeForPath(imagePath string, modes ...int) (*FaceFaceshapeResponse, error) {
	mode := 1
	if len(modes) > 0 {
		mode = modes[0]
	}
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.FaceFaceshapeForBase64(img, mode)
}
