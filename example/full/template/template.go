package template

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
	"time"
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

func getNowSecondFormat() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
}

func SendTemplateMsgTest(wc wechat.IWeChat) {
	tl := wc.Template()
	wc.RegisterEventFunc(func(reply common.IReply, event *common.CEvent, communicate common.CDataCommunicate, userData interface{}) error {
		if event.EventKey == "order" {
			request := common.CSendTemplateMsgRequest{}
			request.Touser = communicate.FromUserName
			request.TemplateId = "SOni-e9rT401RT1MiZ8gA9gPQwZImxuFduin9XMM8Ko"
			request.Url = "https://www.baidu.com"
			items := make(map[string]common.CTemplateMessageItem)
			items["first"] = common.CTemplateMessageItem{Value: "恭喜你购买成功", Color: "#173177"}
			items["product"] = common.CTemplateMessageItem{Value: "巧克力", Color: "#173177"}
			items["amount"] = common.CTemplateMessageItem{Value: "39.8元", Color: "#173177"}
			items["time"] = common.CTemplateMessageItem{Value: getNowSecondFormat(), Color: "#173177"}
			items["remark"] = common.CTemplateMessageItem{Value: "欢迎再次购买", Color: "#173177"}
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
			var _ = b
		}
		return nil
	}, nil)
}
