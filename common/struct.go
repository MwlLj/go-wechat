package common

import (
	"encoding/xml"
	"time"
)

var (
	MsgTypeText       string = "test"
	MsgTypeImage      string = "image"
	MsgTypeVoice      string = "voice"
	MsgTypeVideo      string = "video"
	MsgTypeShortVideo string = "shortvideo"
	MsgTypeLocation   string = "location"
	MsgTypeLink       string = "link"
)

type CUserInfo struct {
	AppId     string
	AppSecret string
	Port      int
	Url       string
	Token     string
}

type CData string

func (c CData) MarshelXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		str string `xml:",cdata"`
	}{string(c)}, start)
}

type CMessage struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   CData         `xml:"ToUserName"`
	FromUserName CData         `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      CData         `xml:"MsgType"`
	MsgId        int64         `xml:"MsgId"`
	Content      CData         `xml:"Content"`
	PicUrl       CData         `xml:"PicUrl"`
	MediaId      int64         `xml:"MediaId"`
	Format       CData         `xml:"Format"`
	ThumbMediaId CData         `xml:"ThumbMediaId"`
	Location_X   float64       `xml:"Location_X"`
	Location_Y   float64       `xml:"Location_Y"`
	Scale        int           `xml:"Scale"`
	Label        CData         `xml:"Label"`
	Title        CData         `xml:"Title"`
	Description  CData         `xml:"Description"`
	Url          CData         `xml:"Url"`
}

func (this *CMessage) SetToUserName(toUserName *CData) {
	this.ToUserName = *toUserName
}

func (this *CMessage) SetFromUserName(fromUserName *CData) {
	this.FromUserName = *fromUserName
}

func (this *CMessage) SetCreateTime(t *time.Duration) {
	this.CreateTime = *t
}

func (this *CMessage) SetMsgType(msgType string) {
	this.MsgType = CData(msgType)
}

func (this *CMessage) SetMsgId(msgId int64) {
	this.MsgId = msgId
}

func (this *CMessage) SetContent(content string) {
	this.Content = CData(content)
}

func (this *CMessage) SetPicUrl(url string) {
	this.PicUrl = CData(url)
}

func (this *CMessage) SetMediaId(mediaId int64) {
	this.MediaId = mediaId
}

func (this *CMessage) SetFormat(format string) {
	this.Format = CData(format)
}

func (this *CMessage) SetThumbMediaId(thumbMediaId string) {
	this.ThumbMediaId = CData(thumbMediaId)
}

func (this *CMessage) SetLocationX(x float64) {
	this.Location_X = x
}

func (this *CMessage) SetLocationY(y float64) {
	this.Location_Y = y
}

func (this *CMessage) SetScale(scale int) {
	this.Scale = scale
}

func (this *CMessage) SetLabel(label string) {
	this.Label = CData(label)
}

func (this *CMessage) SetTitle(title string) {
	this.Title = CData(title)
}

func (this *CMessage) SetDesription(description string) {
	this.Description = CData(description)
}

func (this *CMessage) SetUrl(url string) {
	this.Url = CData(url)
}
