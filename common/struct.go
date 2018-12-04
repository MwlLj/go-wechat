package common

import (
	"encoding/xml"
	"time"
)

var (
	ErrorCodeSuccess int = 0
)

var (
	MsgTypeText       string = "text"
	MsgTypeImage      string = "image"
	MsgTypeVoice      string = "voice"
	MsgTypeVideo      string = "video"
	MsgTypeShortVideo string = "shortvideo"
	MsgTypeLocation   string = "location"
	MsgTypeLink       string = "link"
	MsgTypeEvent      string = "event"
)

var (
	ButtonTypeClick           string = "click"
	ButtonTypeView            string = "view"
	ButtonTypeScancodeWaitmsg string = "scancode_waitmsg"
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
	Status       CData         `xml:"Status"`
}

type CDataCommunicate struct {
	ToUserName   string
	FromUserName string
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
	CreateTime time.Duration
	MsgType    string
	Event      string
	EventKey   string
	Ticket     string
	Latitude   float64
	Longitude  float64
	Precision  float64
	Status     string
}

type CButton struct {
	SubButton           []CButton `json:"sub_button"`
	Type                string    `json:"type"`
	Name                string    `json:"name"`
	Key                 string    `json:"key"`
	Url                 string    `json:"url"`
	MediaId             string    `json:"media_id"`
	MiniProgramAppId    string    `json:"appid"`
	MiniProgramPagePath string    `json:"pagepath"`
}

type CGetMenuButton struct {
	Button []CButton `json:"button"`
}

type CGetMenuJson struct {
	Menu            CGetMenuButton `json:"menu"`
	ConditionalMenu CGetMenuButton `json:"conditionalmenu"`
}

type CSetIndustryRequest struct {
	FirstIndustryId  string `json:"industry_id1"`
	SecondIndustryId string `json:"industry_id2"`
}

type CGetIndustryInfo struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

type CGetIndustryResponse struct {
	ErrCode           int              `json:"errcode"`
	ErrMsg            string           `json:"errmsg"`
	PrimaryIndustry   CGetIndustryInfo `json:"primary_industry"`
	SecondaryIndustry CGetIndustryInfo `json:"secondary_industry"`
}

type CGetTemplateIdRequest struct {
	TemplateLibId string `json:"template_id_short"`
}

type CGetTemplateIdResponse struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	TemplateId string `json:"template_id"`
}

type CGetTemplateListInfo struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type CGetTemplateListResponse struct {
	ErrCode      int                    `json:"errcode"`
	ErrMsg       string                 `json:"errmsg"`
	TemplateList []CGetTemplateListInfo `json:"template_list"`
}

type CDeleteTemplateRequest struct {
	TemplateId string `json:"template_id"`
}

type CTemplateMessageItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type CSendTemplateMsgRequest struct {
	Touser     string                          `json:"touser"`
	TemplateId string                          `json:"template_id"`
	Url        string                          `json:"url"`
	Data       map[string]CTemplateMessageItem `json:"data"`
}

type CSendTemplateMsgResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgId   int64  `json:"msgid"`
}
