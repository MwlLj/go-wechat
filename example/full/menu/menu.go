package menu

import (
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func DeleteMenuTest(wc wechat.IWeChat) {
	menu := wc.Menu()
	err := menu.DeleteAll(3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("delete menu success")
}

func CreateMenuTest(wc wechat.IWeChat) {
	// data
	buttons := []common.CButton{}
	button1 := common.CButton{}
	button1.Name = "下单"
	button1.Type = common.ButtonTypeClick
	button1.Key = "order"
	buttons = append(buttons, button1)
	buttonSearch := common.CButton{}
	searchButtons := []common.CButton{}
	buttonSearch.Name = "搜索"
	baiduSearch := common.CButton{}
	baiduSearch.Name = "百度一下"
	baiduSearch.Type = common.ButtonTypeView
	baiduSearch.Url = "https://www.baidu.com"
	searchButtons = append(searchButtons, baiduSearch)
	sougouSearch := common.CButton{}
	sougouSearch.Name = "搜狗搜索"
	sougouSearch.Type = common.ButtonTypeView
	sougouSearch.Url = "http://www.sogou.com"
	searchButtons = append(searchButtons, sougouSearch)
	buttonSearch.SubButton = searchButtons
	buttons = append(buttons, buttonSearch)
	button3 := common.CButton{}
	button3.Name = "扫一扫"
	button3.Type = common.ButtonTypeScancodeWaitmsg
	button3.Key = "take_photo"
	buttons = append(buttons, button3)
	// create
	menu := wc.Menu()
	err := menu.Create(&buttons, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("create menu success")
}

func GetMenuTest(wc wechat.IWeChat) {
	menu := wc.Menu()
	res, err := menu.GetAll(3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print
	var _ = res
	fmt.Println("get menu success")
}
