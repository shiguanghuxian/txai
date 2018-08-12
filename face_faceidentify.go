package txai

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

var (
	faceFaceidentifyURI = "/face/face_faceidentify"
)

// FaceFaceidentifyForBase64 人脸识别 - 图片为base64格式
// image 图片base64后字符串
// groupId 候选人组ID（个体创建时设定）
// topns 返回的候选人个数（默认9个）
func (ai *TxAi) FaceFaceidentifyForBase64(image string, groupId string, topns ...int) (*FaceFaceidentifyResponse, error) {
	topn := 9
	if len(topns) > 0 {
		if topns[0] > 0 && topns[0] < 11 {
			topn = topns[0]
		}
	}
	params := ai.getPublicParams()
	params.Add("image", image)
	params.Add("group_id", groupId)
	params.Add("topn", fmt.Sprint(topn))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceFaceidentifyResponse := new(FaceFaceidentifyResponse)
	err := ai.RequestAPI(faceFaceidentifyURI, params, faceFaceidentifyResponse)
	if err != nil {
		return nil, err
	}
	return faceFaceidentifyResponse, nil
}

// FaceFaceidentifyForPath 人脸识别 - 图片为本地路径
// imagePath 本地图片路径
// groupId 候选人组ID（个体创建时设定）
// topns 返回的候选人个数（默认9个）
func (ai *TxAi) FaceFaceidentifyForPath(imagePath string, groupId string, topns ...int) (*FaceFaceidentifyResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.FaceFaceidentifyForBase64(img, groupId, topns...)
}
