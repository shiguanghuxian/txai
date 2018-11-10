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

// PtuResponse 图片特效统一响应格式
type PtuResponse struct {
	BaseResponse
	Data struct {
		Image []byte `json:"image"`
	} `json:"data"`
}

/* 自然语言处理 */

// NlpWordsegResponse 分词
type NlpWordsegResponse struct {
	BaseResponse
	Data struct {
		Text       string `json:"text"`
		BaseTokens []struct {
			Word   string `json:"word"`
			Offset int    `json:"offset"`
			Length int    `json:"length"`
		} `json:"base_tokens"`
		MixTokens []struct {
			Word   string `json:"word"`
			Offset int    `json:"offset"`
			Length int    `json:"length"`
		} `json:"mix_tokens"`
	} `json:"data"`
}

// NlpWordcomResponse 语义解析
type NlpWordcomResponse struct {
	BaseResponse
	Data struct {
		Text       string `json:"text"`
		Intent     int    `json:"intent"`
		IntentName string `json:"intent_name"`
		ComTokens  []*struct {
			ComType     int    `json:"com_type"`
			ComTypeName string `json:"com_type_name"`
			ComWord     string `json:"com_word"`
		} `json:"com_tokens"`
	} `json:"data"`
}

// NlpWordnerResponse 专有名词识别接口
type NlpWordnerResponse struct {
	BaseResponse
	Data struct {
		Text      string `json:"text"`
		NerTokens []*struct {
			Word      string   `json:"word"`
			Offset    int      `json:"offset"`
			Length    int      `json:"length"`
			Types     []int    `json:"types"`
			Weights   []int    `json:"weights"`
			TypeNames []string `json:"type_names"`
		} `json:"ner_tokens"`
	} `json:"data"`
}

// NlpWordsynResponse 同义词识别
type NlpWordsynResponse struct {
	BaseResponse
	Data struct {
		Text      string `json:"text"`
		SynTokens []struct {
			OriWord struct {
				Word   string `json:"word"`
				Offset int    `json:"offset"`
				Length int    `json:"length"`
			} `json:"ori_word"`
			SynWords []struct {
				Word   string  `json:"word"`
				Weight float64 `json:"weight"`
			} `json:"syn_words"`
		} `json:"syn_tokens"`
	} `json:"data"`
}

// NlpWordposResponse 词性标注
type NlpWordposResponse struct {
	BaseResponse
	Data struct {
		Text       string `json:"text"`
		BaseTokens []*struct {
			Word        string `json:"word"`
			Offset      int    `json:"offset"`
			Length      int    `json:"length"`
			PosCode     int    `json:"pos_code"`
			PosCodeName string `json:"pos_code_name"`
		} `json:"base_tokens"`
		MixTokens []*struct {
			Word        string `json:"word"`
			Offset      int    `json:"offset"`
			Length      int    `json:"length"`
			PosCode     int    `json:"pos_code"`
			PosCodeName string `json:"pos_code_name"`
		} `json:"mix_tokens"`
	} `json:"data"`
}

// NlpTextpolarResponse 情感分析
type NlpTextpolarResponse struct {
	BaseResponse
	Data struct {
		Text      string  `json:"text"`
		Polar     int     `json:"polar"`
		PolarName string  `json:"polar_name"`
		Confd     float64 `json:"confd"`
	} `json:"data"`
}

// NlpTextchatResponse 智能闲聊
type NlpTextchatResponse struct {
	BaseResponse
	Data struct {
		Session string `json:"session"`
		Answer  string `json:"answer"`
	} `json:"data"`
}

// NlpTexttransResponse 文本翻译（AI Lab）
type NlpTexttransResponse struct {
	BaseResponse
	Data struct {
		Type      LangType `json:"type"`
		OrgText   string   `json:"org_text"`
		TransText string   `json:"trans_text"`
	} `json:"data"`
}

// NlpTexttranslateResponse 文本翻译（翻译君）
type NlpTexttranslateResponse struct {
	BaseResponse
	Data struct {
		SourceText string `json:"source_text"`
		TargetText string `json:"target_text"`
	} `json:"data"`
}

// NlpImagetranslateResponse 图片翻译
type NlpImagetranslateResponse struct {
	BaseResponse
	Data struct {
		SessionID    string `json:"session_id"`
		ImageRecords []struct {
			SourceText string `json:"source_text"`
			TargetText string `json:"target_text"`
			X          int    `json:"x"`
			Y          int    `json:"y"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
		} `json:"image_records"`
	} `json:"data"`
}

// NlpSpeechtranslateResponse 语音翻译
type NlpSpeechtranslateResponse struct {
	BaseResponse
	Data struct {
		SessionID  string `json:"session_id"`
		End        int    `json:"end"`
		Seq        int    `json:"seq"`
		SourceText string `json:"source_text"`
		TargetText string `json:"target_text"`
	} `json:"data"`
}

// NlpTextdetectResponse 语种识别
type NlpTextdetectResponse struct {
	BaseResponse
	Data struct {
		Lang Lang `json:"lang"`
	} `json:"data"`
}

// AaiAsrResponse 语音识别-echo版
type AaiAsrResponse struct {
	BaseResponse
	Data struct {
		Format AaiAudioType `json:"format"`
		Rate   int          `json:"rate"`
		Text   string       `json:"text"`
	} `json:"data"`
}

// AaiAsrsResponse 语音识别-流式版（AI Lab）
type AaiAsrsResponse struct {
	BaseResponse
	Data struct {
		Format     int    `json:"format"`
		Rate       int    `json:"rate"`
		Seq        int    `json:"seq"`
		Len        int    `json:"len"`
		End        int    `json:"end"`
		SpeechID   string `json:"speech_id"`
		SpeechText string `json:"speech_text"`
	} `json:"data"`
}

// AaiWxasrsResponse 语音识别-流式版（AI Lab）
type AaiWxasrsResponse struct {
	BaseResponse
	Data struct {
		Format     int    `json:"format"`
		Rate       int    `json:"rate"`
		End        int    `json:"end"`
		SpeechID   string `json:"speech_id"`
		SpeechText string `json:"speech_text"`
		IsFinalRes int    `json:"is_final_res"`
		Ack        int    `json:"ack"`
	} `json:"data"`
}

// AaiWxasrlongResponse 长语音识别
type AaiWxasrlongResponse struct {
	BaseResponse
	Data struct {
		TaskID string `json:"task_id"`
		Text   string `json:"text"`
	} `json:"data"`
}

// AaiDetectkeywordResponse 关键词检索
type AaiDetectkeywordResponse struct {
	BaseResponse
	Data struct {
		Bps   int `json:"bps"`
		Eps   int `json:"eps"`
		IsEnd int `json:"is_end"`
		Res   struct {
			KeyWords []struct {
				KeyWord string  `json:"key_word"`
				Mbtm    int     `json:"mbtm"`
				Metm    float64 `json:"metm"`
				Score   int     `json:"score"`
			} `json:"key_words"`
			KeyWordsSize int `json:"key_words_size"`
		} `json:"res"`
		SegIndex int    `json:"seg_index"`
		TaskID   string `json:"task_id"`
	} `json:"data"`
}

// AaiTtsResponse 语音合成（AI Lab）
type AaiTtsResponse struct {
	BaseResponse
	Data struct {
		Format int    `json:"format"`
		Speech []byte `json:"speech"`
		Md5Sum string `json:"md5sum"`
	} `json:"data"`
}

// AaiTtaResponse 语音合成（优图）
type AaiTtaResponse struct {
	BaseResponse
	Data struct {
		Voice []byte `json:"voice"`
	} `json:"data"`
}

// AaiEvilaudioResponse 音频鉴黄
type AaiEvilaudioResponse struct {
	BaseResponse
	Data struct {
		SpeechID  string `json:"speech_id"`
		SpeechURL string `json:"speech_url"`
		PornFlag  int    `json:"porn_flag"`
		PornScore int    `json:"porn_score"`
	} `json:"data"`
}
