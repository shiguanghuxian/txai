package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	ptuFacestickerURI = "/ptu/ptu_facesticker"
)

// PtuFacestickerForBase64 大头贴 - 图片为base64格式
// image 图片base64后字符串
// sticker 变妆编码，定义见 https://ai.qq.com/doc/facesticker.shtml
func (ai *TxAi) PtuFacestickerForBase64(image string, sticker int) (*PtuResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("sticker", fmt.Sprint(sticker))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ptuResponse := new(PtuResponse)
	err := ai.RequestAPI(ptuFacestickerURI, params, ptuResponse)
	if err != nil {
		return nil, err
	}
	return ptuResponse, nil
}

// PtuFacestickerForPath 大头贴 - 图片为本地路径
// image 图片base64后字符串
// sticker 变妆编码，定义见 https://ai.qq.com/doc/facesticker.shtml
func (ai *TxAi) PtuFacestickerForPath(imagePath string, sticker int) (*PtuResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.PtuFacestickerForBase64(img, sticker)
}
