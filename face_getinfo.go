package txai

import (
	"errors"
)

var (
	faceGetinfoURI = "/face/face_getinfo"
)

// FaceGetinfoForPersonId 增加人脸
// personId 指定的个体（Person）ID
func (ai *TxAi) FaceGetinfoForPersonId(personId string) (*FaceGetinfoResponse, error) {
	if personId == "" {
		return nil, errors.New("personId can't be empty")
	}
	params := ai.getPublicParams()
	params.Add("person_id", personId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceGetinfoResponse := new(FaceGetinfoResponse)
	err := ai.RequestAPI(faceGetinfoURI, params, faceGetinfoResponse)
	if err != nil {
		return nil, err
	}
	return faceGetinfoResponse, nil
}
