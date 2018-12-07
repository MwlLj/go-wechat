package user

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	// "github.com/MwlLj/go-wechat/common"
)

func GetFollowUsersTest(wc wechat.IWeChat) {
	user := wc.User()
	res, err := user.GetFollowUsers(3000)
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
