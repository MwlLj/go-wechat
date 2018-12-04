package menu

import (
	"fmt"
	"github.com/MwlLj/go-wechat"
	// "github.com/MwlLj/go-wechat/common"
)

func GetMenuTest(wc wechat.IWeChat) {
	menu := wc.Menu()
	res, err := menu.GetAll(3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	// print
	fmt.Println(res)
}
