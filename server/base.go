package server

import (
	"encoding/xml"
	"time"
)

type CTextRequest struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`
	Content      string        `xml:"Content"`
	MsgId        int64         `xml:"MsgId"`
}

type CTextResponse struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CData         `xml:"ToUserName"`
	FromUserName CData         `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      CData         `xml:"MsgType"`
	Content      CData         `xml:"Content"`
	MsgId        int64         `xml:"MsgId"`
}

type CData string

func (c CData) MarshelXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		str string `xml:",cdata"`
	}{string(c)}, start)
}
