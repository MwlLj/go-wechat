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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &CreateTagUrl, &method, nil, nil, b)
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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTagListUrl, &method, nil, nil, nil)
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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UpdateTagUrl, &method, nil, nil, b)
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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteTagUrl, &method, nil, nil, b)
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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTagUserListUrl, &method, nil, nil, b)
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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &AddTagToUsersUrl, &method, nil, nil, b)
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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteTagToUsersUrl, &method, nil, nil, b)
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
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTagsByUserUrl, &method, nil, nil, b)
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

func (this *CUser) UpdateUserRemark(request *common.CUpdateUserRemarkRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UpdateUserRemarkUrl, &method, nil, nil, b)
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

func (this *CUser) GetUserBaseInfo(request *common.CGetUserBaseInfoRequest, timeoutMS int64) (*common.CGetUserBaseInfoResponse, error) {
	method := http.MethodGet
	params := make(map[string]string)
	params[GetUserBaseInfoParamOpenId] = request.OpenId
	params[GetUserBaseInfoParamLang] = request.Lang
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetUserBaseInfoUrl, &method, &params, nil, nil)
	if err != nil {
		return nil, err
	}
	response := common.CGetUserBaseInfoResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CUser) GetUserBaseInfoMulti(request *common.CGetUserBaseInfoMultiRequest, timeoutMS int64) (*common.CGetUserBaseInfoMultiResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetUserBaseInfoMultiUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetUserBaseInfoMultiResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CUser) getFollowUserOnce(nextOpenId *string, openIds *[]string, timeoutMS int64) error {
	method := http.MethodGet
	params := make(map[string]string)
	params[GetFollowUsersParamNextOpenId] = *nextOpenId
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetFollowUsersUrl, &method, &params, nil, nil)
	if err != nil {
		return err
	}
	response := CSingleUsers{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	for _, v := range response.Data.OpenId {
		*openIds = append(*openIds, v)
	}
	if response.NextOpenId != "" {
		this.getFollowUserOnce(&response.NextOpenId, openIds, timeoutMS)
	}
	return nil
}

func (this *CUser) GetFollowUsers(timeoutMS int64) (*common.CGetFollowUsersResponse, error) {
	method := http.MethodGet
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetFollowUsersUrl, &method, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	response := CSingleUsers{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	if response.NextOpenId != "" {
		this.getFollowUserOnce(&response.NextOpenId, &response.Data.OpenId, timeoutMS)
	}
	var follows common.CGetFollowUsersResponse
	follows.OpenIds = response.Data.OpenId
	return &follows, nil
}

func (this *CUser) getBlackListUserOnce(nextOpenId *string, openIds *[]string, timeoutMS int64) error {
	request := CGetBlackListSingle{}
	request.BeginOpenId = *nextOpenId
	b, err := json.Marshal(&request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetBlackListUsersUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := CSingleUsers{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	for _, v := range response.Data.OpenId {
		*openIds = append(*openIds, v)
	}
	if response.NextOpenId != "" {
		this.getBlackListUserOnce(&response.NextOpenId, openIds, timeoutMS)
	}
	return nil
}

func (this *CUser) GetBlackListUsers(timeoutMS int64) (*common.CGetBlackListUsersResponse, error) {
	request := CGetBlackListSingle{}
	request.BeginOpenId = ""
	b, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetBlackListUsersUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := CSingleUsers{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	if response.NextOpenId != "" {
		this.getBlackListUserOnce(&response.NextOpenId, &response.Data.OpenId, timeoutMS)
	}
	var blackList common.CGetBlackListUsersResponse
	blackList.OpenIds = response.Data.OpenId
	return &blackList, nil
}

func (this *CUser) TakeUsersToBlackList(request *common.CTakeUsersToBlackListRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &TakeUsersToBlackListUrl, &method, nil, nil, b)
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

func (this *CUser) UnTakeUsersToBlackList(request *common.CUnTakeUsersToBlackListRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UnTakeUsersToBlackListUrl, &method, nil, nil, b)
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

func New(token common.IToken) common.IUser {
	u := CUser{
		m_token: token,
	}
	return &u
}
