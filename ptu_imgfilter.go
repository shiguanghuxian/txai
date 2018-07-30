package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	ptuImgfilterURI = "/ptu/ptu_imgfilter"
)

// PtuImgfilterForBase64 图片滤镜（天天P图） - 图片为base64格式
// image 图片base64后字符串
// filter 变妆编码，定义见 https://ai.qq.com/doc/ptuimgfilter.shtml
func (ai *TxAi) PtuImgfilterForBase64(image string, filter int) (*PtuResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("filter", fmt.Sprint(filter))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ptuResponse := new(PtuResponse)
	err := ai.RequestAPI(ptuImgfilterURI, params, ptuResponse)
	if err != nil {
		return nil, err
	}
	return ptuResponse, nil
}

// PtuImgfilterForPath 图片滤镜（天天P图） - 图片为本地路径
// imagePath 本地图片路径
// filter 变妆编码，定义见 https://ai.qq.com/doc/ptuimgfilter.shtml
func (ai *TxAi) PtuImgfilterForPath(imagePath string, filter int) (*PtuResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.PtuImgfilterForBase64(img, filter)
}
