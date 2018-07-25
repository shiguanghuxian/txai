package txai

import (
	"encoding/base64"
	"io/ioutil"
)

var (
	faceFaceverifyURI = "/face/face_faceverify"
)

// FaceFaceverifyForBase64 人脸验证 - 图片为base64格式
// image 图片base64后字符串
// personId 待验证的个体（Person）ID
func (ai *TxAi) FaceFaceverifyForBase64(image, personId string) (*FaceFaceverifyResponse, error) {
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("person_id", personId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceFaceverifyResponse := new(FaceFaceverifyResponse)
	err := ai.RequestAPI(faceFaceverifyURI, params, faceFaceverifyResponse)
	if err != nil {
		return nil, err
	}
	return faceFaceverifyResponse, nil
}

// FaceFaceverifyForPath 人脸验证 - 图片为本地路径
// imagePath 本地图片路径
// personId 待验证的个体（Person）ID
func (ai *TxAi) FaceFaceverifyForPath(imagePath, personId string) (*FaceFaceverifyResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.FaceFaceverifyForBase64(img, personId)
}
