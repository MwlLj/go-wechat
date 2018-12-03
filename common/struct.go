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

type CWxResXml struct {
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
	Event        CData         `xml:"Event"`
	EventKey     CData         `xml:"EventKey"`
	Ticket       CData         `xml:"Ticket"`
	Latitude     float64       `xml:"Latitude"`
	Longitude    float64       `xml:"Longitude"`
	Precision    float64       `xml:"Precision"`
}

func (this *CWxResXml) SetToUserName(toUserName *CData) {
	this.ToUserName = *toUserName
}

func (this *CWxResXml) SetFromUserName(fromUserName *CData) {
	this.FromUserName = *fromUserName
}

func (this *CWxResXml) SetCreateTime(t *time.Duration) {
	this.CreateTime = *t
}

func (this *CWxResXml) SetMsgType(msgType string) {
	this.MsgType = CData(msgType)
}

func (this *CWxResXml) SetMsgId(msgId int64) {
	this.MsgId = msgId
}

func (this *CWxResXml) SetContent(content string) {
	this.Content = CData(content)
}

func (this *CWxResXml) SetPicUrl(url string) {
	this.PicUrl = CData(url)
}

func (this *CWxResXml) SetMediaId(mediaId int64) {
	this.MediaId = mediaId
}

func (this *CWxResXml) SetFormat(format string) {
	this.Format = CData(format)
}

func (this *CWxResXml) SetThumbMediaId(thumbMediaId string) {
	this.ThumbMediaId = CData(thumbMediaId)
}

func (this *CWxResXml) SetLocationX(x float64) {
	this.Location_X = x
}

func (this *CWxResXml) SetLocationY(y float64) {
	this.Location_Y = y
}

func (this *CWxResXml) SetScale(scale int) {
	this.Scale = scale
}

func (this *CWxResXml) SetLabel(label string) {
	this.Label = CData(label)
}

func (this *CWxResXml) SetTitle(title string) {
	this.Title = CData(title)
}

func (this *CWxResXml) SetDesription(description string) {
	this.Description = CData(description)
}

func (this *CWxResXml) SetUrl(url string) {
	this.Url = CData(url)
}

type CMessage struct {
	CreateTime   time.Duration
	MsgType      string
	MsgId        int64
	Content      string
	PicUrl       string
	MediaId      int64
	Format       string
	ThumbMediaId string
	LocationX    float64
	LocationY    float64
	Scale        int
	Label        string
	Title        string
	Description  string
	Url          string
}

type CEvent struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Event        string
	EventKey     string
	Ticket       string
	Latitude     float64
	Longitude    float64
	Precision    float64
}
