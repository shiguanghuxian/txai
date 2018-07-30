package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	ptuFacedecorationURI = "/ptu/ptu_facedecoration"
)

// PtuFacedecorationForBase64 人脸变妆 - 图片为base64格式
// image 图片base64后字符串
// decoration 变妆编码，定义见 https://ai.qq.com/doc/facedecoration.shtml
func (ai *TxAi) PtuFacedecorationForBase64(image string, decoration int) (*PtuResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("decoration", fmt.Sprint(decoration))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ptuResponse := new(PtuResponse)
	err := ai.RequestAPI(ptuFacedecorationURI, params, ptuResponse)
	if err != nil {
		return nil, err
	}
	return ptuResponse, nil
}

// PtuFacedecorationForPath 人脸变妆 - 图片为本地路径
// imagePath 本地图片路径
// decoration 变妆编码，定义见 https://ai.qq.com/doc/facedecoration.shtml
func (ai *TxAi) PtuFacedecorationForPath(imagePath string, decoration int) (*PtuResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.PtuFacedecorationForBase64(img, decoration)
}
