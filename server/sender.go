package server

import (
	"encoding/xml"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"net/http"
	"time"
)

type CSender struct {
	m_responseWriter http.ResponseWriter
}

func (this *CSender) SendMessage(msg *common.CMessage) error {
	response := *msg
	response.SetToUserName(&msg.FromUserName)
	response.SetFromUserName(&msg.ToUserName)
	res, err := xml.MarshalIndent(&response, " ", "  ")
	if err != nil {
		return err
	}
	this.m_responseWriter.Header().Set("Content-Type", "text/xml")
	fmt.Fprint(this.m_responseWriter, string(res))
	return nil
}
