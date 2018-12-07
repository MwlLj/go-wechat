package store

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"github.com/MwlLj/go-wechat/private"
	"net/http"
)

type CStore struct {
	m_token common.IToken
}

func (this *CStore) CreateStore(request *common.CCreateStoreRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &CreateStoreUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := private.CCommonResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CStore) GetStore(request *common.CGetStoreRequest, timeoutMS int64) (*common.CGetStoreResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetStoreUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetStoreResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CStore) GetStoreList(request *common.CGetStoreListRequest, timeoutMS int64) (*common.CGetStoreListResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetStoreListUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetStoreListResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CStore) ModifyStore(request *common.CModifyStoreRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &ModifyStoreUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := private.CCommonResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CStore) DeleteStore(request *common.CDeleteStoreRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteStoreUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := private.CCommonResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CStore) GetCategory(timeoutMS int64) (*common.CGetCategoryResponse, error) {
	method := http.MethodGet
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetCategoryUrl, &method, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	response := common.CGetCategoryResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func New(token common.IToken) common.IStore {
	s := CStore{
		m_token: token,
	}
	return &s
}
