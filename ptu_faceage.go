package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	ptuFaceageURI = "/ptu/ptu_faceage"
)

// PtuFaceageForBase64 颜龄检测 - 图片为base64格式
// image 图片base64后字符串
func (ai *TxAi) PtuFaceageForBase64(image string) (*PtuResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ptuResponse := new(PtuResponse)
	err := ai.RequestAPI(ptuFaceageURI, params, ptuResponse)
	if err != nil {
		return nil, err
	}
	return ptuResponse, nil
}

// PtuFaceageForPath 颜龄检测 - 图片为本地路径
// image 图片base64后字符串
// sticker 变妆编码，定义见 https://ai.qq.com/doc/facesticker.shtml
func (ai *TxAi) PtuFaceageForPath(imagePath string) (*PtuResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.PtuFaceageForBase64(img)
}
