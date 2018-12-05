package user

import (
	"github.com/MwlLj/go-wechat/common"
	// "github.com/MwlLj/go-wechat/communicate"
)

type CUser struct {
	m_token common.IToken
}

func New(token common.IToken) common.IUser {
	u := CUser{
		m_token: token,
	}
	return &u
}
