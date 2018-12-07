package menu

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"net/http"
)

type CCommonResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type CMenu struct {
	m_token common.IToken
}

func (this *CMenu) Create(data *[]common.CButton, timeoutMS int64) error {
	menu := CMenuData{Buttons: *data}
	b, err := json.Marshal(&menu)
	if err != nil {
		return err
	}
	method := http.MethodPost
	body, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &CreateMenuUrl, &method, nil, nil, b)
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
	body, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetMenuUrl, &method, nil, nil, nil)
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
	body, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteMenuUrl, &method, nil, nil, nil)
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
