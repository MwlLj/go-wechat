package wechat

import (
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/server"
)

type IWeChat interface {
	RegisterEvent(callback common.IEvent, userData interface{})
	RegisterMsg(callback common.IMessage, userData interface{})
}

func New(info *common.CUserInfo) IWeChat {
	return server.New(info)
}
