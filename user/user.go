package user

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"github.com/MwlLj/go-wechat/private"
	"net/http"
)

type CUser struct {
	m_token common.IToken
}

func (this *CUser) CreateTag(request *common.CCreateTagRequest, timeoutMS int64) (*common.CCreateTagResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &CreateTagUrl, &method, b)
	if err != nil {
		return nil, err
	}
	response := common.CCreateTagResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CUser) GetTagList(timeoutMS int64) (*common.CGetTagListResponse, error) {
	method := http.MethodGet
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTagListUrl, &method, nil)
	if err != nil {
		return nil, err
	}
	response := common.CGetTagListResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CUser) UpdateTag(request *common.CUpdateTagRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UpdateTagUrl, &method, b)
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

func (this *CUser) DeleteTag(request *common.CDeleteTagRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteTagUrl, &method, b)
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

func (this *CUser) GetTagUserList(request *common.CGetTagUserListRequest, timeoutMS int64) (*common.CGetTagUserListResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &CreateTagUrl, &method, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetTagUserListResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CUser) AddTagToUsers(request *common.CAddTagToUsersRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &AddTagToUsersUrl, &method, b)
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

func (this *CUser) DeleteTagToUsers(request *common.CDeleteTagToUsersRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteTagToUsersUrl, &method, b)
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

func (this *CUser) GetTagsByUser(request *common.CGetTagsByUserRequest, timeoutMS int64) (*common.CGetTagsByUserResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTagsByUserUrl, &method, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetTagsByUserResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func New(token common.IToken) common.IUser {
	u := CUser{
		m_token: token,
	}
	return &u
}
