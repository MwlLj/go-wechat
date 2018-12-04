# go-wechat
封装微信调用接口

## install
	go get github.com/MwlLj/go-wechat

## import
	import (
	    "github.com/MwlLj/go-wechat"
	    "github.com/MwlLj/go-wechat/common"
	)

## sample example
	package main

	import (
		"fmt"
		"github.com/MwlLj/go-wechat"
		"github.com/MwlLj/go-wechat/common"
	)

	var _ = fmt.Println

	func main() {
		info := common.CUserInfo{
			AppId:     "your appid",
			AppSecret: "your appsecret",
			Port:      80,
			Url:       "/xxx",
			Token:     "your token",
		}
		wc := wechat.New(&info)
		wc.RegisterMsgFunc(func(reply common.IReply, msg *common.CMessage, userData interface{}) error {
			reply.SendMessage(msg)
			return nil
		}, nil)
		wc.Loop()
	}

## detail example
[address](https://github.com/MwlLj/go-wechat/tree/master/example/full)

### 微信公众号测试帐号开发流程
	1. 进入 https://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/index
	2. 填写基本信息
	    appID: 自动生成
	    appsecret: 自动生成
	    URL: 接收微信服务器的url, 如果不配置 nginx, 必须用 80 端口(直接 http://ip/xxx)
	    Token: 任意填写, 但是如果是正式版的公众号, 请写得复杂一些
	3. 可以将上面的例子(记得将 AppId 等基本信息改成你自己的)放在远程服务器上, go run main.go
	4. 点击测试帐号的 提交, 只有步骤3启动之后, 测试帐号的配置信息提交才会成功
