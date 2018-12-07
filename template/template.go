package template

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"net/http"
)

var _ = fmt.Println

type CTemplate struct {
	m_token common.IToken
}

func (this *CTemplate) SetIndustry(request *common.CSetIndustryRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &SetIndustryUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := CCommonResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CTemplate) GetIndustry(timeoutMS int64) (*common.CGetIndustryResponse, error) {
	method := http.MethodGet
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetIndustryUrl, &method, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	response := common.CGetIndustryResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CTemplate) GetTemplateId(request *common.CGetTemplateIdRequest, timeoutMS int64) (*common.CGetTemplateIdResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTemplateIdByTemplLibIdUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetTemplateIdResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CTemplate) GetTemplateList(timeoutMS int64) (*common.CGetTemplateListResponse, error) {
	method := http.MethodGet
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTemplateListUrl, &method, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	response := common.CGetTemplateListResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CTemplate) DeleteTemplate(request *common.CDeleteTemplateRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteTemplateUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := CCommonResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CTemplate) SendTemplateMsg(request *common.CSendTemplateMsgRequest, timeoutMS int64) (*common.CSendTemplateMsgResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &SendTemplateMsgUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CSendTemplateMsgResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != common.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func New(token common.IToken) common.ITemplate {
	template := CTemplate{m_token: token}
	return &template
}
