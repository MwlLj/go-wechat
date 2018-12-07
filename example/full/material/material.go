package material

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func AddForeverOtherMaterialTest(wc wechat.IWeChat) {
	media := wc.Material()
	request := common.CAddForeverOtherMaterialRequest{}
	request.MaterialType = common.MaterialTypeImage
	request.Path = "./testresource/test.jpg"
	request.Title = "考拉资讯"
	request.Introduction = "考拉 ... 考拉 ..."
	res, err := media.AddForeverOtherMaterial(&request, 3000)
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

func AddForeverImgTextMaterialTest(wc wechat.IWeChat) {
	media := wc.Material()
	request := common.CAddForeverImgTextMaterialRequest{}
	articles := []common.CImgTextMaterial{}
	article := common.CImgTextMaterial{}
	article.Title = "测试"
	article.ThumbMediaId = "RcnqQNVBusVo6Fx-3qGKRRueLS_k6vJ8gXqbhgfkzmQ"
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
	request.Type = common.MaterialTypeImgText
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
