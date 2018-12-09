package shop

import (
	"github.com/MwlLj/go-wechat/common"
)

type CShop struct {
	m_token common.IToken
}

func New(token common.IToken) common.IShop {
	shop := CShop{m_token: token}
	return &shop
}
