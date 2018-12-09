package main

import (
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/example/full/event"
	"github.com/MwlLj/go-wechat/example/full/material"
	"github.com/MwlLj/go-wechat/example/full/menu"
	"github.com/MwlLj/go-wechat/example/full/message"
	"github.com/MwlLj/go-wechat/example/full/sender"
	"github.com/MwlLj/go-wechat/example/full/store"
	"github.com/MwlLj/go-wechat/example/full/template"
	"github.com/MwlLj/go-wechat/example/full/user"
)

var _ = fmt.Println

func messageTest(wc wechat.IWeChat) {
	message.RegisterMsgTest(wc)
	message.RegisterMsgFuncTest(wc)
}

func eventTest(wc wechat.IWeChat) {
	event.RegisterEventTest(wc)
}

func menuTest(wc wechat.IWeChat) {
	menu.DeleteMenuTest(wc)
	menu.CreateMenuTest(wc)
	menu.GetMenuTest(wc)
}

func templateTest(wc wechat.IWeChat) {
	// template.SetIndustryTest(wc)
	// template.GetIndustryTest(wc)
	// template.GetTemplateListTest(wc)
	template.SendTemplateMsgTest(wc)
}

func materialTest(wc wechat.IWeChat) {
	// material.UploadImageTest(wc)
	// material.AddForeverOtherMaterialTest(wc)
	material.AddForeverImgTextMaterialTest(wc)
	material.GetMaterialTotalTest(wc)
	material.GetMaterialListTest(wc)
}

func userTest(wc wechat.IWeChat) {
	user.GetFollowUsersTest(wc)
}

func storeTest(wc wechat.IWeChat) {
	// store.UploadImageTest(wc)
	store.CreateStoreTest(wc)
}

func senderTest(wc wechat.IWeChat) {
	// sender.GroupSendByTagTest(wc)
	sender.PreviewMessageTest(wc)
}

func main() {
	info := common.CUserInfo{
		// AppId:     "wxfedcab8946a21ccc",
		AppSecret: "7ae3cc7b23b34a7b8d8cea44f7e9177f",
		Port:      80,
		Url:       "/test/wechat/hello",
		Token:     "c0082b4d-b46f-4d67-b0eb-7a0d15dd5ef2",
	}
	wc := wechat.New(&info)
	// message test
	messageTest(wc)
	// event test
	eventTest(wc)
	// menu test
	// menuTest(wc)
	// template test
	templateTest(wc)
	// material test
	// materialTest(wc)
	// user test
	// userTest(wc)
	// store test
	// storeTest(wc)
	// sender test
	senderTest(wc)
	wc.Loop()
}
