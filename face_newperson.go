package txai

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
)

var (
	faceNewpersonURI = "/face/face_newperson"
)

// FaceNewpersonForBase64 个体创建 - 图片为base64格式
// groupIds 个体组，可以用|分开多个组
// personId 指定的个体（Person）ID
// image 图片base64后字符串
// personName 个体名称
// tags 备注信息
func (ai *TxAi) FaceNewpersonForBase64(groupIds, personId, image, personName string, tags ...string) (*FaceNewpersonResponse, error) {
	if groupIds == "" {
		return nil, errors.New("groupIds can't be empty")
	}
	if personId == "" {
		return nil, errors.New("personId can't be empty")
	}
	if personName == "" {
		return nil, errors.New("personName can't be empty")
	}
	tag := ""
	if len(tags) > 0 {
		tag = tags[0]
	}

	params := ai.getPublicParams()
	params.Add("group_ids", groupIds)
	params.Add("person_id", personId)
	params.Add("image", image)
	params.Add("person_name", personName)
	if tag != "" {
		params.Add("tag", tag)
	}
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceNewpersonResponse := new(FaceNewpersonResponse)
	err := ai.RequestAPI(faceNewpersonURI, params, faceNewpersonResponse)
	if err != nil {
		return nil, err
	}
	return faceNewpersonResponse, nil
}

// FaceNewpersonForPath 个体创建 - 图片为本地路径
// groupIds 个体组，可以用|分开多个组
// personId 指定的个体（Person）ID
// imagePath 本地图片路径
// personName 个体名称
// tags 备注信息
func (ai *TxAi) FaceNewpersonForPath(groupIds, personId, imagePath, personName string, tags ...string) (*FaceNewpersonResponse, error) {
	body, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}
	img := base64.StdEncoding.EncodeToString(body)
	return ai.FaceNewpersonForBase64(groupIds, personId, img, personName, tags...)
}
