package main

import (
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

var _ = fmt.Println

type CMessageCallback struct {
}

func (this *CMessageCallback) OnMessage(reply common.IReply, msg *common.CMessage, userData interface{}) error {
	msg.Content = "hello struct"
	reply.SendMessage(msg)
	return nil
}

func main() {
	info := common.CUserInfo{
		AppId:     "wxfedcab8946a21ccc",
		AppSecret: "7ae3cc7b23b34a7b8d8cea44f7e9177f",
		Port:      80,
		Url:       "/test/wechat/hello",
		Token:     "c0082b4d-b46f-4d67-b0eb-7a0d15dd5ef2",
	}
	wc := wechat.New(&info)
	// wc.RegisterMsg(&CMessageCallback{}, nil)
	wc.RegisterMsgFunc(func(reply common.IReply, msg *common.CMessage, userData interface{}) error {
		msg.Content = "hello function"
		reply.SendMessage(msg)
		return nil
	}, nil)
	wc.RegisterEventFunc(func(reply common.IReply, event *common.CEvent, userData interface{}) error {
		msg := common.CMessage{}
		msg.Content = "event: " + event.Event
		msg.MsgType = common.MsgTypeText
		reply.SendMessage(&msg)
		return nil
	}, nil)
	wc.Loop()
}
