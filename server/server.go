package server

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/material"
	"github.com/MwlLj/go-wechat/menu"
	"github.com/MwlLj/go-wechat/pay"
	"github.com/MwlLj/go-wechat/sender"
	"github.com/MwlLj/go-wechat/shop"
	"github.com/MwlLj/go-wechat/store"
	"github.com/MwlLj/go-wechat/template"
	"github.com/MwlLj/go-wechat/token"
	"github.com/MwlLj/go-wechat/user"
	"github.com/MwlLj/go-wechat/utils"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type CServer struct {
	m_exeChannel            chan bool
	m_userInfo              common.CUserInfo
	m_decodeFactory         CDecodeFactory
	m_token                 common.IToken
	m_msgCallback           common.IMessage
	m_msgCallbackUserdata   interface{}
	m_eventCallback         common.IEvent
	m_eventCallbackUserdata interface{}
}

func (this *CServer) init(info *common.CUserInfo) {
	this.m_token = token.New(info)
	this.startListen(info.Port, &info.Url)
}

func (this *CServer) Loop() {
	this.m_exeChannel = make(chan bool, 1)
	<-this.m_exeChannel
	close(this.m_exeChannel)
}

func (this *CServer) makeSignature(timestamp, nonce string) string {
	sl := []string{this.m_userInfo.Token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func (this *CServer) validateUrl(w http.ResponseWriter, r *http.Request) bool {
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signatureGen := this.makeSignature(timestamp, nonce)

	signatureIn := strings.Join(r.Form["signature"], "")
	if signatureGen != signatureIn {
		return false
	}
	echostr := strings.Join(r.Form["echostr"], "")
	fmt.Fprint(w, echostr)
	return true
}

func (this *CServer) parseResContent(body []byte, w http.ResponseWriter) (*common.CWxResXml, error) {
	param := CDecodeParam{}
	param.DecodeType = DecodeTypeResXml
	decoding := this.m_decodeFactory.Decoding(&param)
	if decoding == nil {
		fmt.Fprint(w, "decoding message error")
		return nil, errors.New("decoding message error")
	}
	message := decoding.Parse(body)
	if message == nil {
		fmt.Fprint(w, "parse message request error")
		return nil, errors.New("parse message request error")
	}
	msg := message.(*common.CWxResXml)
	return msg, nil
}

func (this *CServer) handleCheck(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	if !this.validateUrl(w, r) {
		fmt.Fprint(w, "check invalid")
		return errors.New("check invalid")
	}
	return nil
}

func (this *CServer) handlePost(w http.ResponseWriter, r *http.Request) error {
	var err error = nil
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	resXml, err := this.parseResContent(body, w)
	if resXml == nil {
		return err
	}
	if string(resXml.MsgType) != common.MsgTypeEvent {
		// message
		reply := sender.NewReply(&resXml.ToUserName, &resXml.FromUserName, &w)
		if this.m_msgCallback != nil {
			for {
				msg := utils.ResXml2Message(resXml)
				communicate := utils.ResXml2DataCommunicate(resXml)
				err = this.m_msgCallback.OnMessage(reply, msg, *communicate, this.m_msgCallbackUserdata)
				if err != nil {
					break
				}
				if reply.IsSend() == false {
					reply.SendEmptyMessage()
				}
				return nil
			}
		}
	} else {
		// event
		reply := sender.NewReply(&resXml.ToUserName, &resXml.FromUserName, &w)
		if this.m_eventCallback != nil {
			for {
				event := utils.ResXml2Event(resXml)
				communicate := utils.ResXml2DataCommunicate(resXml)
				err = this.m_eventCallback.OnEvent(reply, event, *communicate, this.m_eventCallbackUserdata)
				if err != nil {
					break
				}
				if reply.IsSend() == false {
					reply.SendEmptyMessage()
				}
				return nil
			}
		}
	}
	return err
}

func (this *CServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	this.handleCheck(w, r)
	if r.Method == http.MethodPost {
		this.handlePost(w, r)
	}
}

func (this *CServer) startListen(port int, u *string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		this.handleRequest(w, r)
	}
	go func() {
		url := *u
		if string(url[0]) != "/" {
			url = strings.Join([]string{"/", url}, "")
		}
		host := strings.Join([]string{":", strconv.FormatInt(int64(port), 10)}, "")
		mux := http.NewServeMux()
		mux.HandleFunc(url, handler)
		err := http.ListenAndServe(host, mux)
		if err != nil {
			fmt.Println(err)
			this.m_exeChannel <- false
		}
	}()
}

func (this *CServer) RegisterMsg(callback common.IMessage, userData interface{}) {
	this.m_msgCallback = callback
	this.m_msgCallbackUserdata = userData
}

func (this *CServer) RegisterMsgFunc(callback common.FuncMsgCallback, userData interface{}) {
	this.RegisterMsg(&CMsgCallbackDefault{MsgCallback: callback}, userData)
}

func (this *CServer) RegisterEventFunc(callback common.FuncEventCallback, userData interface{}) {
	this.RegisterEvent(&CEventCallbackDefault{EventCallback: callback}, userData)
}

func (this *CServer) RegisterEvent(callback common.IEvent, userData interface{}) {
	this.m_eventCallback = callback
	this.m_eventCallbackUserdata = userData
}

func (this *CServer) Menu() common.IMenu {
	return menu.New(this.m_token)
}

func (this *CServer) Template() common.ITemplate {
	return template.New(this.m_token)
}

func (this *CServer) Material() common.IMaterial {
	return material.New(this.m_token)
}

func (this *CServer) User() common.IUser {
	return user.New(this.m_token)
}

func (this *CServer) Store() common.IStore {
	return store.New(this.m_token)
}

func (this *CServer) Shop() common.IShop {
	return shop.New(this.m_token)
}

func (this *CServer) Sender() common.ISender {
	return sender.New(this.m_token)
}

func (this *CServer) PayByPaymengCode() common.IPayByPaymentCode {
	return pay.NewPayByPaymentCode(this.m_token)
}

func New(info *common.CUserInfo) *CServer {
	server := CServer{m_userInfo: *info}
	server.init(info)
	return &server
}
