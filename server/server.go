package server

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
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
	m_msgCallback           common.IMessage
	m_msgCallbackUserdata   interface{}
	m_eventCallback         common.IEvent
	m_eventCallbackUserdata interface{}
}

func (this *CServer) init(info *common.CUserInfo) {
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

func (this *CServer) parseMessage(body []byte, w http.ResponseWriter) *common.CMessage {
	param := CDecodeParam{}
	param.DecodeType = DecodeTypeMessage
	decoding := this.m_decodeFactory.Decoding(&param)
	if decoding == nil {
		fmt.Fprint(w, "decoding message error")
		return nil
	}
	message := decoding.Parse(body)
	if message == nil {
		fmt.Fprint(w, "parse message request error")
		return nil
	}
	msg := message.(*common.CMessage)
	return msg
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
	msg := this.parseMessage(body, w)
	sender := CSender{m_responseWriter: w}
	if this.m_msgCallback != nil {
		for {
			err = this.m_msgCallback.OnMessage(&sender, msg, this.m_msgCallbackUserdata)
			if err != nil {
				break
			}
			return nil
		}
	}
	sender.SendMessage(msg)
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

func (this *CServer) RegisterEvent(callback common.IEvent, userData interface{}) {
	this.m_eventCallback = callback
	this.m_eventCallbackUserdata = userData
}

func New(info *common.CUserInfo) *CServer {
	server := CServer{m_userInfo: *info}
	server.init(info)
	return &server
}
