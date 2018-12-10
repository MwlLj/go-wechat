package wechat

import (
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/server"
)

type IWeChat interface {
	Menu() common.IMenu
	Template() common.ITemplate
	Material() common.IMaterial
	User() common.IUser
	Store() common.IStore
	Shop() common.IShop
	Sender() common.ISender
	PayByPaymengCode() common.IPayByPaymentCode
	RegisterEvent(callback common.IEvent, userData interface{})
	RegisterEventFunc(callback common.FuncEventCallback, userData interface{})
	RegisterMsg(callback common.IMessage, userData interface{})
	RegisterMsgFunc(callback common.FuncMsgCallback, userData interface{})
	Loop()
}

func New(info *common.CUserInfo) IWeChat {
	return server.New(info)
}
