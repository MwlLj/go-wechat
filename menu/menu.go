package menu

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"io/ioutil"
	"net/http"
	"strings"
)

type CCommonResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type CMenu struct {
	m_token common.IToken
}

func (this *CMenu) send(timeoutMS int64, url *string, method *string, payload []byte) (resBody []byte, e error) {
	var err error = nil
	token, err := this.m_token.GetToken(timeoutMS)
	if err != nil {
		return nil, err
	}
	var req *http.Request = nil
	if payload == nil {
		req, err = http.NewRequest(*method, *url, nil)
	} else {
		reader := strings.NewReader(string(payload))
		req, err = http.NewRequest(*method, *url, reader)
	}
	if err != nil {
		return nil, err
	}
	values := req.URL.Query()
	values.Add(AccessToken, string(token))
	req.URL.RawQuery = values.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (this *CMenu) Create(data *[]common.CButton, timeoutMS int64) error {
	menu := CMenuData{Buttons: *data}
	b, err := json.Marshal(&menu)
	if err != nil {
		return err
	}
	method := http.MethodPost
	body, err := this.send(timeoutMS, &CreateMenuUrl, &method, b)
	if err != nil {
		return err
	}
	response := CCommonResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CMenu) GetAll(timeoutMS int64) (*common.CGetMenuJson, error) {
	method := http.MethodGet
	body, err := this.send(timeoutMS, &GetMenuUrl, &method, nil)
	if err != nil {
		return nil, err
	}
	response := common.CGetMenuJson{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (this *CMenu) DeleteAll(timeoutMS int64) error {
	method := http.MethodGet
	body, err := this.send(timeoutMS, &DeleteMenuUrl, &method, nil)
	if err != nil {
		return err
	}
	response := CCommonResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func New(token common.IToken) common.IMenu {
	menu := CMenu{m_token: token}
	return &menu
}
