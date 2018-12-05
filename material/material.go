package material

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"github.com/MwlLj/go-wechat/private"
)

type CMaterial struct {
	m_token common.IToken
}

func (this *CMaterial) UploadImage(path *string, timeoutMS int64) (*common.CUploadImageResponse, error) {
	resBody, err := communicate.UploadFileWithToken(this.m_token, timeoutMS, &UploadImageFormname, path, &UploadImageUrl)
	if err != nil {
		return nil, err
	}
	response := common.CUploadImageResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func New(token common.IToken) common.IMaterial {
	media := CMaterial{
		m_token: token,
	}
	return &media
}
