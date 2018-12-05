package store

import (
	"encoding/json"
	"fmt"
	"github.com/MwlLj/go-wechat"
	"github.com/MwlLj/go-wechat/common"
)

func UploadImageTest(wc wechat.IWeChat) {
	st := wc.Store()
	path := "./testresource/test.jpg"
	res, err := st.UploadImage(&path, 3000)
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

func CreateStoreTest(wc wechat.IWeChat) {
	st := wc.Store()
	request := common.CCreateStoreRequest{}
	err := st.CreateStore(&request, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

func GetStoreTest(wc wechat.IWeChat) {
	st := wc.Store()
	request := common.CGetStoreRequest{}
	res, err := st.GetStore(&request, 3000)
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

func GetStoreListTest(wc wechat.IWeChat) {
	st := wc.Store()
	request := common.CGetStoreListRequest{}
	res, err := st.GetStoreList(&request, 3000)
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

func ModifyStoreTest(wc wechat.IWeChat) {
	st := wc.Store()
	request := common.CModifyStoreRequest{}
	err := st.ModifyStore(&request, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

func DeleteStoreTest(wc wechat.IWeChat) {
	st := wc.Store()
	request := common.CDeleteStoreRequest{}
	err := st.DeleteStore(&request, 3000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}

func GetCategoryTest(wc wechat.IWeChat) {
	st := wc.Store()
	res, err := st.GetCategory(3000)
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
