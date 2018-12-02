package server

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CServer struct {
	m_userInfo common.CUserInfo
}

func (this *CServer) init(info *common.CUserInfo) {
	exeChan := make(chan bool, 1)
	this.startListen(info.Port, &info.Url, exeChan)
	<-exeChan
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
	fmt.Fprintf(w, echostr)
	return true
}

func (this *CServer) parseTextRequestBody(r *http.Request) ([]byte, *CTextRequest) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return body, nil
	}
	requestBody := &CTextRequest{}
	xml.Unmarshal(body, requestBody)
	return body, requestBody
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &CTextResponse{}
	textResponseBody.FromUserName = CData(fromUserName)
	textResponseBody.ToUserName = CData(toUserName)
	textResponseBody.MsgType = CData("text")
	textResponseBody.Content = CData(content)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

func (this *CServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request ...")
	r.ParseForm()
	if !this.validateUrl(w, r) {
		fmt.Fprintf(w, "check invalid")
		return
	}

	if r.Method == "POST" {
		body, textRequest := this.parseTextRequestBody(r)
		if textRequest != nil {
			w.Header().Set("Content-Type", "text/xml")
			fmt.Fprintf(w, string(body))
		}
		return
	}
	fmt.Fprintf(w, "not found")
}

func (this *CServer) startListen(port int, u *string, ch chan<- bool) {
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
			ch <- false
		}
	}()
}

func (this *CServer) RegisterEvent(callback common.IEvent, userData interface{}) {
}

func (this *CServer) RegisterMsg(callback common.IMessage, userData interface{}) {
}

func New(info *common.CUserInfo) *CServer {
	server := CServer{m_userInfo: *info}
	server.init(info)
	return &server
}
