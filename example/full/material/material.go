package material

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	// "github.com/MwlLj/go-wechat/common"
)

func UploadImageTest(wc wechat.IWeChat) {
	media := wc.Material()
	path := "./testresource/test.jpg"
	res, err := media.UploadImage(&path, 3000)
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
