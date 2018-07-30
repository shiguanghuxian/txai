package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	ptuFacecosmeticURI = "/ptu/ptu_facecosmetic"
)

// PtuFacecosmeticForBase64 人脸美妆 - 图片为base64格式
// image 图片base64后字符串
// cosmetic 美妆编码，定义见 https://ai.qq.com/doc/facecosmetic.shtml
func (ai *TxAi) PtuFacecosmeticForBase64(image string, cosmetic int) (*PtuResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("cosmetic", fmt.Sprint(cosmetic))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ptuResponse := new(PtuResponse)
	err := ai.RequestAPI(ptuFacecosmeticURI, params, ptuResponse)
	if err != nil {
		return nil, err
	}
	return ptuResponse, nil
}

// PtuFacecosmeticForPath 人脸美妆 - 图片为本地路径
// imagePath 本地图片路径
// cosmetic 美妆编码，定义见 https://ai.qq.com/doc/facecosmetic.shtml
func (ai *TxAi) PtuFacecosmeticForPath(imagePath string, cosmetic int) (*PtuResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.PtuFacecosmeticForBase64(img, cosmetic)
}
