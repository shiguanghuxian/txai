package txai

var (
	faceDelpersonURI = "/face/face_delperson"
)

// FaceDelpersonForPersonId 删除个体
// personId 需要删除的个体（Person）ID
func (ai *TxAi) FaceDelpersonForPersonId(personId string) (*FaceDelpersonResponse, error) {
	params := ai.getPublicParams()
	params.Add("person_id", personId)
	sign := ai.getReqSign(params)
	params.Add("sign", sign)
	// 响应结果
	faceDelpersonResponse := new(FaceDelpersonResponse)
	err := ai.RequestAPI(faceDelpersonURI, params, faceDelpersonResponse)
	if err != nil {
		return nil, err
	}
	return faceDelpersonResponse, nil
}
