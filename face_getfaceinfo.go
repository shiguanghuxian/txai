package txai

import "errors"

var (
	faceGetfaceinfoURI = "/face/face_getfaceinfo"
)

// FaceGetfaceinfoForFaceId 获取人脸信息 - 根据人脸（Face）ID 获取人脸（Face）信息
func (ai *TxAi) FaceGetfaceinfoForFaceId(faceId string) (*FaceGetfaceinfoResponse, error) {
	if faceId == "" {
		return nil, errors.New("faceId can't be empty")
	}
	params := ai.getPublicParams()
	params.Add("face_id", faceId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceGetfaceinfoResponse := new(FaceGetfaceinfoResponse)
	err := ai.RequestAPI(faceGetfaceinfoURI, params, faceGetfaceinfoResponse)
	if err != nil {
		return nil, err
	}
	return faceGetfaceinfoResponse, nil
}
