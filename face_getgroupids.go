package txai

var (
	faceGetgroupidsURI = "/face/face_getgroupids"
)

// FaceGetgroupidsAll 获取组列表 - 获取应用下所有的组（Group）ID列表
func (ai *TxAi) FaceGetgroupidsAll() (*FaceGetgroupidsResponse, error) {
	params := ai.getPublicParams()
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceGetgroupidsResponse := new(FaceGetgroupidsResponse)
	err := ai.RequestAPI(faceGetgroupidsURI, params, faceGetgroupidsResponse)
	if err != nil {
		return nil, err
	}
	return faceGetgroupidsResponse, nil
}
