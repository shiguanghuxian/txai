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

// 人脸识别

// FaceDetectfaceResponse 人脸检测与分析
type FaceDetectfaceResponse struct {
	BaseResponse
	Data struct {
		ImageWidth  int `json:"image_width"`
		ImageHeight int `json:"image_height"`
		FaceList    []struct {
			FaceID     string `json:"face_id"`
			X          int    `json:"x"`
			Y          int    `json:"y"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			Gender     int    `json:"gender"`
			Age        int    `json:"age"`
			Expression int    `json:"expression"`
			Beauty     int    `json:"beauty"`
			Glass      int    `json:"glass"`
			Pitch      int    `json:"pitch"`
			Yaw        int64  `json:"yaw"`
			Roll       int    `json:"roll"`
			FaceShape  struct {
				FaceProfile []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"face_profile"`
				LeftEye []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"left_eye"`
				RightEye []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"right_eye"`
				LeftEyebrow []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"left_eyebrow"`
				RightEyebrow []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"right_eyebrow"`
				Mouth []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"mouth"`
				Nose []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"nose"`
			} `json:"face_shape"`
		} `json:"face_list"`
	} `json:"data"`
}

// FaceDetectmultifaceResponse 多人脸检测
type FaceDetectmultifaceResponse struct {
	BaseResponse
	Data struct {
		FaceList []struct {
			X1 float64 `json:"x1"`
			Y1 float64 `json:"y1"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"face_list"`
	} `json:"data"`
}

// FaceFacecompareResponse 人脸对比
type FaceFacecompareResponse struct {
	BaseResponse
	Data struct {
		Similarity int `json:"similarity"`
		FailFlag   int `json:"fail_flag"`
	} `json:"data"`
}

// FaceDetectcrossagefaceResponse 跨年龄人脸识别
type FaceDetectcrossagefaceResponse struct {
	BaseResponse
	Data struct {
		SourceFace struct {
			X1 float64 `json:"x1"`
			Y1 float64 `json:"y1"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"source_face"`
		TargetFace struct {
			X1 float64 `json:"x1"`
			Y1 float64 `json:"y1"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"target_face"`
		Score    float64 `json:"score"`
		FailFlag int     `json:"fail_flag"`
	} `json:"data"`
}

// FaceFaceshapeResponse 五官定位
type FaceFaceshapeResponse struct {
	BaseResponse
	Data struct {
		ImageWidth  int `json:"image_width"`
		ImageHeight int `json:"image_height"`
		FaceList    []struct {
			FaceID     string `json:"face_id"`
			X          int    `json:"x"`
			Y          int    `json:"y"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			Gender     int    `json:"gender"`
			Age        int    `json:"age"`
			Expression int    `json:"expression"`
			Beauty     int    `json:"beauty"`
			Glass      int    `json:"glass"`
			Pitch      int    `json:"pitch"`
			Yaw        int    `json:"yaw"`
			Roll       int    `json:"roll"`
			FaceShape  struct {
				FaceProfile []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"face_profile"`
				LeftEye []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"left_eye"`
				RightEye []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"right_eye"`
				LeftEyebrow []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"left_eyebrow"`
				RightEyebrow []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"right_eyebrow"`
				Mouth []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"mouth"`
				Nose []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"nose"`
			} `json:"face_shape"`
		} `json:"face_list"`
	} `json:"data"`
}

// FaceFaceidentifyResponse 人脸识别
type FaceFaceidentifyResponse struct {
	BaseResponse
	Data struct {
		TimeMs     int `json:"time_ms"`
		GroupSize  int `json:"group_size"`
		Candidates []struct {
			PersonID   string `json:"person_id"`
			FaceID     string `json:"face_id"`
			Confidence int    `json:"confidence"`
			Tag        string `json:"tag"`
		} `json:"candidates"`
	} `json:"data"`
}

// FaceFaceverifyResponse 人脸验证
type FaceFaceverifyResponse struct {
	BaseResponse
	Data struct {
		Ismatch    int `json:"ismatch"`
		Confidence int `json:"confidence"`
	} `json:"data"`
}

// FaceNewpersonResponse 个体创建
type FaceNewpersonResponse struct {
	BaseResponse
	Data struct {
		SucGroup int      `json:"suc_group"`
		SucFace  int      `json:"suc_face"`
		PersonID string   `json:"person_id"`
		FaceID   string   `json:"face_id"`
		GroupIds []string `json:"group_ids"`
	} `json:"data"`
}

// FaceDelpersonResponse 删除个体
type FaceDelpersonResponse struct {
	BaseResponse
	Data struct {
		Deleted  int    `json:"deleted"`
		PersonID string `json:"person_id"`
	} `json:"data"`
}

// FaceAddfaceResponse 增加人脸
type FaceAddfaceResponse struct {
	BaseResponse
	Data struct {
		Added    int      `json:"added"`
		FaceIds  []string `json:"face_ids"`
		RetCodes []int    `json:"ret_codes"`
	} `json:"data"`
}

// FaceDelfaceResponse 删除人脸
type FaceDelfaceResponse struct {
	BaseResponse
	Data struct {
		Deleted int      `json:"deleted"`
		FaceIds []string `json:"face_ids"`
	} `json:"data"`
}

// FaceSetinfoResponse 设置信息
type FaceSetinfoResponse struct {
	BaseResponse
	Data struct {
		PersonID string `json:"person_id"`
	} `json:"data"`
}

// FaceGetinfoResponse 获取信息
type FaceGetinfoResponse struct {
	BaseResponse
	Data struct {
		PersonID   string   `json:"person_id"`
		PersonName string   `json:"person_name"`
		Tag        string   `json:"tag"`
		FaceIds    []string `json:"face_ids"`
		GroupIds   []string `json:"group_ids"`
	} `json:"data"`
}

// FaceGetgroupidsResponse 获取组列表
type FaceGetgroupidsResponse struct {
	BaseResponse
	Data struct {
		GroupIds []string `json:"group_ids"`
	} `json:"data"`
}

// FaceGetpersonidsResponse 获取个体列表
type FaceGetpersonidsResponse struct {
	BaseResponse
	Data struct {
		PersonIds []string `json:"person_ids"`
	} `json:"data"`
}

// FaceGetfaceidsResponse 获取人脸列表 - 根据个体（Person）ID 获取人脸（Face）ID列表
type FaceGetfaceidsResponse struct {
	BaseResponse
	Data struct {
		FaceIds []string `json:"face_ids"`
	} `json:"data"`
}

// FaceGetfaceinfoResponse 获取人脸信息 - 根据人脸（Face）ID 获取人脸（Face）信息
type FaceGetfaceinfoResponse struct {
	BaseResponse
	Data struct {
		FaceInfo struct {
			FaceID     string `json:"face_id"`
			X          int    `json:"x"`
			Y          int    `json:"y"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			Gender     int    `json:"gender"`
			Age        int    `json:"age"`
			Expression int    `json:"expression"`
			Beauty     int    `json:"beauty"`
			Glass      int    `json:"glass"`
			Pitch      int    `json:"pitch"`
			Yaw        int64  `json:"yaw"`
			Roll       int64  `json:"roll"`
		} `json:"face_info"`
	} `json:"data"`
}

/* 图片识别 */

// VisionScenerResponse 场景识别
type VisionScenerResponse struct {
	BaseResponse
	Data struct {
		Topk      int `json:"topk"`
		SceneList []*struct {
			LabelID    int     `json:"label_id"`
			LabelName  string  `json:"label_name"` // 场景名
			LabelConfd float64 `json:"label_confd"`
		} `json:"scene_list"`
	} `json:"data"`
}

// VisionObjectrResponse 物体识别
type VisionObjectrResponse struct {
	BaseResponse
	Data struct {
		Topk       int `json:"topk"`
		ObjectList []*struct {
			LabelID    int     `json:"label_id"`
			LabelName  string  `json:"label_name"` // 物体名
			LabelConfd float64 `json:"label_confd"`
		} `json:"object_list"`
	} `json:"data"`
}

// ImageTagResponse 图像标签识别
type ImageTagResponse struct {
	BaseResponse
	Data struct {
		TagList []struct {
			TagConfidence int    `json:"tag_confidence"`
			TagName       string `json:"tag_name"`
		} `json:"tag_list"`
	} `json:"data"`
}

// VisionImgtotextResponse  看图说话
type VisionImgtotextResponse struct {
	BaseResponse
	Data struct {
		Text string `json:"text"`
	} `json:"data"`
}

// ImageFuzzyResponse 模糊图片检测
type ImageFuzzyResponse struct {
	BaseResponse
	Data struct {
		Fuzzy      bool    `json:"fuzzy"`
		Confidence float64 `json:"confidence"`
	} `json:"data"`
}

// ImageFoodResponse 美食图片识别
type ImageFoodResponse struct {
	BaseResponse
	Data struct {
		Food       bool    `json:"food"`
		Confidence float64 `json:"confidence"`
	} `json:"data"`
}
