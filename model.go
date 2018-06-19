package txai

// 使用到的模型

// BaseResponseInterface 基础模型接口，用于统一请求使用
type BaseResponseInterface interface {
	GetRet() int
}

// BaseResponse 基本响应结构
type BaseResponse struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

// GetRet 获取状态码
func (resp *BaseResponse) GetRet() int {
	return resp.Ret
}

/* OCR */

// OcrIDCardResponse 身份证识别
type OcrIDCardResponse struct {
	BaseResponse
	Data struct {
		Name       string `json:"name"`
		Sex        string `json:"sex"`
		Nation     string `json:"nation"`
		Birth      string `json:"birth"`
		Address    string `json:"address"`
		Id         string `json:"id"`
		Frontimage string `json:"frontimage"`
		Authority  string `json:"authority"`
		ValidDate  string `json:"valid_date"`
		Backimage  string `json:"backimage"`
	} `json:"data"`
}

// OcrResponse OCR 通用响应结构
type OcrResponse struct {
	BaseResponse
	Data struct {
		Angle    string     `json:"angle"`
		ItemList []ItemList `json:"item_list"`
	} `json:"data"`
}

// ItemList OCR 通用ItemList
type ItemList struct {
	Item       string `json:"item"`
	Itemstring string `json:"itemstring"`
	Itemcoord  []struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"itemcoord"`
	Itemconf float64 `json:"itemconf,omitempty"`
	Words    []struct {
		Character  string  `json:"character"`
		Confidence float64 `json:"confidence"`
	} `json:"words,omitempty"`
}

/* 智能鉴黄 */

// VisionPornResponse 智能鉴黄响应结构
type VisionPornResponse struct {
	BaseResponse
	Level int // 图片级别 0正常、1性感、2色情
	Data  struct {
		TagList []struct {
			TagConfidence  int     `json:"tag_confidence"`
			TagConfidenceF float64 `json:"tag_confidence_f"`
			TagName        string  `json:"tag_name"`
		} `json:"tag_list"`
	} `json:"data"`
}

// ImageTerrorismResponse 暴恐图片识别
type ImageTerrorismResponse struct {
	BaseResponse
	Terrorists int // 恐怖分子
	Knife      int // 刀
	Guns       int // 枪支
	Blood      int // 血液
	Fire       int // 火
	Flag       int // 旗帜
	Crowd      int // 人群
	Other      int // 其他
	Data       struct {
		TagList []struct {
			TagConfidence  int     `json:"tag_confidence"`
			TagConfidenceF float64 `json:"tag_confidence_f"`
			TagName        string  `json:"tag_name"`
		} `json:"tag_list"`
	} `json:"data"`
}
