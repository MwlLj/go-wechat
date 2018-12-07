package material

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func AddForeverImgTextMaterialTest(wc wechat.IWeChat) {
	media := wc.Material()
	request := common.CAddForeverImgTextMaterialRequest{}
	articles := []common.CImgTextMaterial{}
	article := common.CImgTextMaterial{}
	article.Title = "测试"
	article.ThumbMediaId = "http://mmbiz.qpic.cn/mmbiz/gLO17UPS6FS2xsypf378iaNhWacZ1G1UplZYWEYfwvuU6Ont96b1roYsCNFwaRrSaKTPCUdBK9DgEHicsKwWCBRQ/0"
	article.ShowCoverPic = 1
	article.Content = "Hello World"
	article.ContentSourceUrl = "https://www.baidu.com"
	articles = append(articles, article)
	request.Articles = articles
	res, err := media.AddForeverImgTextMaterial(&request, 3000)
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

func GetMaterialTotalTest(wc wechat.IWeChat) {
	media := wc.Material()
	res, err := media.GetMaterialTotal(3000)
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

func GetMaterialListTest(wc wechat.IWeChat) {
	media := wc.Material()
	request := common.CGetMaterialListRequest{}
	request.Type = common.MaterialTypeNews
	request.Offset = 0
	request.Count = 20
	res, err := media.GetMaterialList(&request, 3000)
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
