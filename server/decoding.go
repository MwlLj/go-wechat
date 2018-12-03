package server

import (
	"encoding/xml"
	"github.com/MwlLj/go-wechat/common"
)

var (
	DecodeTypeResXml  string = "resxml"
	DecodeTypeMessage string = "message"
	DecodeTypeEvent   string = "event"
)

type IDecoding interface {
	Parse(body []byte) interface{}
}

type CDecodeParam struct {
	DecodeType string
}

type CDecodeFactory struct {
}

func (this *CDecodeFactory) Decoding(param *CDecodeParam) IDecoding {
	if param.DecodeType == DecodeTypeMessage {
		return &CMessageDecoding{}
	} else if param.DecodeType == DecodeTypeEvent {
		return &CEventDecoding{}
	} else if param.DecodeType == DecodeTypeResXml {
		return &CResXmlDecoding{}
	}
	return nil
}

type CResXmlDecoding struct {
}

func (this *CResXmlDecoding) Parse(body []byte) interface{} {
	res := common.CWxResXml{}
	err := xml.Unmarshal(body, &res)
	if err != nil {
		return nil
	}
	return &res
}

type CMessageDecoding struct {
}

func (this *CMessageDecoding) Parse(body []byte) interface{} {
	msg := common.CMessage{}
	err := xml.Unmarshal(body, &msg)
	if err != nil {
		return nil
	}
	return &msg
}

type CEventDecoding struct {
}

func (this *CEventDecoding) Parse(body []byte) interface{} {
	event := common.CEvent{}
	err := xml.Unmarshal(body, &event)
	if err != nil {
		return nil
	}
	return &event
}
