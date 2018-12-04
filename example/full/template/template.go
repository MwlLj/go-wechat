package template

import (
	"fmt"
	"github.com/MwlLj/go-wechat/common"
)

func SetIndustryTest(wc common.IWeChat) {
	tl := wc.Template()
	request := common.CSetIndustryRequest{}
	err := tl.SetIndustry(&request, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
}
