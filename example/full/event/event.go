package event

import (
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func RegisterEventTest(wc wechat.IWeChat) {
	wc.RegisterEventFunc(func(reply common.IReply, event *common.CEvent, communicate common.CDataCommunicate, userData interface{}) error {
		msg := common.CMessage{}
		msg.Content = "event: " + event.Event
		msg.MsgType = common.MsgTypeText
		reply.SendMessage(&msg)
		return nil
	}, nil)
}
