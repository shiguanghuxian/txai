package txai

import (
	"errors"
	"strings"
)

var (
	faceDelfaceURI = "/face/face_delface"
)

// FaceDelfaceForPersonId 获取信息
// personId 指定的个体（Person）ID
// faceIds 需要删除的人脸（Face）ID（多个之间用"|")
func (ai *TxAi) FaceDelfaceForPersonId(personId string, faceIds ...string) (*FaceDelfaceResponse, error) {
	if personId == "" {
		return nil, errors.New("personId can't be empty")
	}
	if len(faceIds) == 0 {
		return nil, errors.New("faceIds can't be empty")
	}
	params := ai.getPublicParams()
	params.Add("person_id", personId)
	params.Add("face_ids", strings.Join(faceIds, "|"))
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceDelfaceResponse := new(FaceDelfaceResponse)
	err := ai.RequestAPI(faceDelfaceURI, params, faceDelfaceResponse)
	if err != nil {
		return nil, err
	}
	return faceDelfaceResponse, nil
}
