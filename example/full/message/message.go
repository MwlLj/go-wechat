package message

import (
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

type CMessageCallback struct {
}

func (this *CMessageCallback) OnMessage(reply common.IReply, msg *common.CMessage, communicate common.CDataCommunicate, userData interface{}) error {
	msg.Content = "hello struct"
	reply.SendMessage(msg)
	return nil
}

func RegisterMsgTest(wc wechat.IWeChat) {
	wc.RegisterMsg(&CMessageCallback{}, nil)
}

func RegisterMsgFuncTest(wc wechat.IWeChat) {
	wc.RegisterMsgFunc(func(reply common.IReply, msg *common.CMessage, communicate common.CDataCommunicate, userData interface{}) error {
		msg.Content = "hello function"
		reply.SendMessage(msg)
		return nil
	}, nil)
}
