package txai

import (
	"errors"
)

var (
	faceSetinfoURI = "/face/face_setinfo"
)

// FaceSetinfoForPersonId 设置信息
// personId 指定的个体（Person）ID
// personName 新的名字
// tag 备注信息
func (ai *TxAi) FaceSetinfoForPersonId(personId, personName, tag string) (*FaceSetinfoResponse, error) {
	if personId == "" {
		return nil, errors.New("personId can't be empty")
	}
	if personName == "" {
		return nil, errors.New("personName can't be empty")
	}
	if tag == "" {
		return nil, errors.New("tag can't be empty")
	}
	params := ai.getPublicParams()
	params.Add("person_id", personId)
	params.Add("person_name", personName)
	params.Add("tag", tag)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceSetinfoResponse := new(FaceSetinfoResponse)
	err := ai.RequestAPI(faceSetinfoURI, params, faceSetinfoResponse)
	if err != nil {
		return nil, err
	}
	return faceSetinfoResponse, nil
}
