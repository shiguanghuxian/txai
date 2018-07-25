package txai

import "errors"

var (
	faceGetfaceidsURI = "/face/face_getfaceids"
)

// FaceGetfaceidsAllForPersonId 获取人脸列表 - 根据个体（Person）ID 获取人脸（Face）ID列表
func (ai *TxAi) FaceGetfaceidsAllForPersonId(personId string) (*FaceGetfaceidsResponse, error) {
	if personId == "" {
		return nil, errors.New("personId can't be empty")
	}
	params := ai.getPublicParams()
	params.Add("person_id", personId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceGetfaceidsResponse := new(FaceGetfaceidsResponse)
	err := ai.RequestAPI(faceGetfaceidsURI, params, faceGetfaceidsResponse)
	if err != nil {
		return nil, err
	}
	return faceGetfaceidsResponse, nil
}
