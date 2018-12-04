package sender

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/utils"
	"net/http"
)

type CReply struct {
	ToUserName     common.CData
	FromUserName   common.CData
	ResponseWriter http.ResponseWriter
}

func (this *CReply) SendMessage(msg *common.CMessage) error {
	if msg.MsgType == "" {
		return errors.New("msgType is empty")
	}
	ext := utils.CCommonExt{}
	ext.ToUserName = this.FromUserName
	ext.FromUserName = this.ToUserName
	resXml := utils.Message2ResXml(msg, &ext)
	res, err := xml.MarshalIndent(&resXml, " ", "  ")
	if err != nil {
		return err
	}
	this.ResponseWriter.Header().Set("Content-Type", "text/xml")
	_, err = fmt.Fprint(this.ResponseWriter, string(res))
	if err != nil {
		return err
	}
	return nil
}

func (this *CReply) SendEmptyMessage() error {
	this.ResponseWriter.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(this.ResponseWriter, "")
	if err != nil {
		return err
	}
	return nil
}
