package txai

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// http 请求

// HttpPost 发起POST请求
func (ai *TxAi) HttpPost(url string, u url.Values) (respBody []byte, err error) {
	contentType := "application/x-www-form-urlencoded"
	resp, err := http.Post(url, contentType, strings.NewReader(u.Encode()))
	if err != nil {
		return nil, err
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	return respBody, err
}

// RequestAPI 发起请求，并解析响应结果
func (ai *TxAi) RequestAPI(uri string, u url.Values, response BaseResponseInterface) error {
	apiURL := BaseURL + uri
	// 发起请求
	body, err := ai.HttpPost(apiURL, u)
	if err != nil {
		return err
	}
	// 解析json数据
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	ret := response.GetRet()
	if response.GetRet() != 0 {
		return GetError(ret)
	}
	return nil
}
