package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	ptuFacemergeURI = "/ptu/ptu_facemerge"
)

// PtuFacemergeForBase64 人脸融合 - 图片为base64格式
// image 图片base64后字符串
// model 变妆编码，定义见 https://ai.qq.com/doc/facemerge.shtml
func (ai *TxAi) PtuFacemergeForBase64(image string, model int) (*PtuResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("model", fmt.Sprint(model))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	ptuResponse := new(PtuResponse)
	err := ai.RequestAPI(ptuFacemergeURI, params, ptuResponse)
	if err != nil {
		return nil, err
	}
	return ptuResponse, nil
}

// PtuFacemergeForPath 人脸融合 - 图片为本地路径
// image 图片base64后字符串
// model 变妆编码，定义见 https://ai.qq.com/doc/facemerge.shtml
func (ai *TxAi) PtuFacemergeForPath(imagePath string, model int) (*PtuResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.PtuFacemergeForBase64(img, model)
}
