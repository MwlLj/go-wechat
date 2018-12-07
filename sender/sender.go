package sender

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"github.com/MwlLj/go-wechat/private"
	"net/http"
)

type CSender struct {
	m_token common.IToken
}

func (this *CSender) GroupSendByTag(request *common.CGroupSendByTagRequest, timeoutMS int64) (*common.CGroupSendByTagResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GroupSendByTagUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGroupSendByTagResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CSender) GroupSendByOpenIds(request *common.CGroupSendByOpenIdsRequest, timeoutMS int64) (*common.CGroupSendByTagResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GroupSendByTagUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGroupSendByTagResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CSender) DeleteGroupSend(request *common.CDeleteGroupSendRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteGroupSendUrl, &method, nil, nil, b)
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

func (this *CSender) PreviewMessage(request *common.CPreviewMessageRequest, timeoutMS int64) (*common.CPreviewMessageResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &PreviewMessageUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CPreviewMessageResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func New(token common.IToken) common.ISender {
	sender := CSender{m_token: token}
	return &sender
}
