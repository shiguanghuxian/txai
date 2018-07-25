package txai

import "errors"

var (
	faceGetpersonidsURI = "/face/face_getpersonids"
)

// FaceGetpersonidsAllForGroupId 获取个体列表 - 获取一个组（Group）中的所有个体（Person）ID
func (ai *TxAi) FaceGetpersonidsAllForGroupId(groupId string) (*FaceGetpersonidsResponse, error) {
	if groupId == "" {
		return nil, errors.New("groupId can't be empty")
	}
	params := ai.getPublicParams()
	params.Add("group_id", groupId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceGetpersonidsResponse := new(FaceGetpersonidsResponse)
	err := ai.RequestAPI(faceGetpersonidsURI, params, faceGetpersonidsResponse)
	if err != nil {
		return nil, err
	}
	return faceGetpersonidsResponse, nil
}
