package sender

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func GroupSendByTagTest(wc wechat.IWeChat) {
	send := wc.Sender()
	request := common.CGroupSendByTagRequest{}
	request.Filter = common.CGroupSendFilterInfo{IsToAll: true, TagId: 1}
	request.ImgText = common.CGroupSendMsgContent{MediaId: "RcnqQNVBusVo6Fx-3qGKRej6CNJlQZslLdjrdDOWd48"}
	request.MsgType = common.GroupSendMsgTypeImgText
	res, err := send.GroupSendByTag(&request, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func PreviewMessageTest(wc wechat.IWeChat) {
	send := wc.Sender()
	request := common.CPreviewMessageRequest{}
	// request.ToWxName = "wxid_4j5jlu27drb922"
	request.ToWxName = "lucky-xiqiao"
	request.ImgText = common.CGroupSendMsgContent{MediaId: "RcnqQNVBusVo6Fx-3qGKRej6CNJlQZslLdjrdDOWd48"}
	request.MsgType = common.GroupSendMsgTypeImgText
	res, err := send.PreviewMessage(&request, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
