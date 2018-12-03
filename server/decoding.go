package server

import (
	"encoding/xml"
	"github.com/MwlLj/go-wechat/common"
)

var (
	DecodeTypeMessage string = "message"
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
	}
	return nil
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
