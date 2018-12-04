package template

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func SetIndustryTest(wc wechat.IWeChat) {
	tl := wc.Template()
	request := common.CSetIndustryRequest{}
	request.FirstIndustryId = "1"
	request.SecondIndustryId = "2"
	err := tl.SetIndustry(&request, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetIndustryTest(wc wechat.IWeChat) {
	tl := wc.Template()
	response, err := tl.GetIndustry(3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func GetTemplateListTest(wc wechat.IWeChat) {
	tl := wc.Template()
	response, err := tl.GetTemplateList(3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func SendTemplateMsgTest(wc wechat.IWeChat) {
	tl := wc.Template()
	wc.RegisterEventFunc(func(reply common.IReply, event *common.CEvent, communicate common.CDataCommunicate, userData interface{}) error {
		fmt.Println("[DEBUG] event key:", event.EventKey)
		request := common.CSendTemplateMsgRequest{}
		request.Touser = communicate.FromUserName
		request.TemplateId = "SOni-e9rT401RT1MiZ8gA9gPQwZImxuFduin9XMM8Ko"
		request.Url = "https://www.taobao.com"
		items := make(map[string]common.CTemplateMessageItem)
		items["first"] = common.CTemplateMessageItem{Value: "恭喜你购买成功", Color: "#173177"}
		request.Data = items
		response, err := tl.SendTemplateMsg(&request, 3000)
		if err != nil {
			fmt.Println(err)
			return err
		}
		b, err := json.Marshal(response)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(string(b))
		return nil
	}, nil)
}
