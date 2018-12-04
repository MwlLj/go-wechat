package server

import (
	"github.com/MwlLj/go-wechat/common"
)

type CMsgCallbackDefault struct {
	MsgCallback common.FuncMsgCallback
}

func (this *CMsgCallbackDefault) OnMessage(reply common.IReply, msg *common.CMessage, communicate common.CDataCommunicate, userData interface{}) error {
	return this.MsgCallback(reply, msg, communicate, userData)
}

type CEventCallbackDefault struct {
	EventCallback common.FuncEventCallback
}

func (this *CEventCallbackDefault) OnEvent(reply common.IReply, event *common.CEvent, communicate common.CDataCommunicate, userData interface{}) error {
	return this.EventCallback(reply, event, communicate, userData)
}
