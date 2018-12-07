package common

import (
	"encoding/xml"
	"time"
)

var (
	ErrorCodeSuccess      int = 0
	ErrorCodeTokenInvaild int = 40001
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

type CUploadImageRequest struct {
	Media []byte `json:"media"`
}

type CUploadImageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Url     string `json:"url"`
}

type CStorePhotoUrl struct {
	PhotoUrl string `json:"photo_url"`
}

type CStoreBaseInfo struct {
	Sid            string           `json:"sid"`
	BusinessName   string           `json:"business_name"`
	BranchName     string           `json:"branch_name"`
	Province       string           `json:"province"`
	City           string           `json:"city"`
	District       string           `json:"district"`
	Address        string           `json:"address"`
	Telephone      string           `json:"telephone"`
	Categories     []string         `json:"categories"`
	OffsetType     int              `json:"offset_type"`
	Longitude      float64          `json:"longitude"`
	Latitude       float64          `json:"latitude"`
	PhotoList      []CStorePhotoUrl `json:"photo_list"`
	Recommend      string           `json:"recommend"`
	Special        string           `json:"special"`
	Introduction   string           `json:"introduction"`
	OpenTime       string           `json:"open_time"`
	AvgPrice       int              `json:"avg_price"`
	AvailableState int              `json:"available_state'`
	UpdateStatus   int              `json:"update_status"`
	PoiId          string           `json:"poi_id"`
}

type CStoreBusiness struct {
	BaseInfo CStoreBaseInfo `json:"base_info"`
}

type CCreateStoreRequest struct {
	Business CStoreBusiness `json:"business"`
}

type CGetStoreRequest struct {
	PoiId string `json:"poi_id"`
}

type CGetStoreResponse struct {
	ErrCode  int            `json:"errcode"`
	ErrMsg   string         `json:"errmsg"`
	Business CStoreBusiness `json:"business"`
}

type CGetStoreListRequest struct {
	Begin int `json:"begin"`
	Limit int `json:"limit"`
}

type CGetStoreListResponse struct {
	ErrCode      int              `json:"errcode"`
	ErrMsg       string           `json:"errmsg"`
	BusinessList []CStoreBusiness `json:"business_list"`
	TotalCount   string           `json:"total_count"`
}

type CModifyStoreRequest struct {
	Business CStoreBusiness `json:"business"`
}

type CDeleteStoreRequest struct {
	PoiId string `json:"poi_id"`
}

type CGetCategoryResponse struct {
	ErrCode      int      `json:"errcode"`
	ErrMsg       string   `json:"errmsg"`
	CategoryList []string `json:"category_list"`
}

// user mgr
type CTag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CTagName struct {
	Name string `json:"name"`
}

type CUserTag struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type CCreateTagRequest struct {
	Tag CTagName `json:"tag"`
}

type CCreateTagResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Tag     CTag   `json:"tag"`
}

type CGetTagListResponse struct {
	ErrCode int        `json:"errcode"`
	ErrMsg  string     `json:"errmsg"`
	Tags    []CUserTag `json:"tags"`
}

type CUpdateTagRequest struct {
	Tag CTag `json:"tag"`
}

type CDeleteTagRequest struct {
	Tag CTag `json:"tag"`
}

type CGetTagUserListRequest struct {
	TagId      int64  `json:tagid"`
	NextOpenid string `json:"next_openid"`
}

type CGetTagUserListResponseData struct {
	OpenId []string `json:"openid"`
}

type CGetTagUserListResponse struct {
	ErrCode    int                         `json:"errcode"`
	ErrMsg     string                      `json:"errmsg"`
	Count      int64                       `json:"count"`
	Data       CGetTagUserListResponseData `json:"data"`
	NextOpenid string                      `json:"next_openid"`
}

type CAddTagToUsersRequest struct {
	OpenIdList []string `json:"openid_list"`
	TagId      int64    `json:"tagid"`
}

type CDeleteTagToUsersRequest struct {
	OpenIdList []string `json:"openid_list"`
	TagId      int64    `json:"tagid"`
}

type CGetTagsByUserRequest struct {
	OpenId string `json:"openid"`
}

type CGetTagsByUserResponse struct {
	ErrCode   int      `json:"errcode"`
	ErrMsg    string   `json:"errmsg"`
	TagidList []string `json:"tagid_list"`
}

type CUpdateUserRemarkRequest struct {
	OpenId string `json:"openid"`
	Remark string `json:"remark"`
}

type CGetUserBaseInfoRequest struct {
	OpenId string
	Lang   string
}

type CGetUserBaseInfoResponse struct {
	ErrCode        int    `json:"errcode"`
	ErrMsg         string `json:"errmsg"`
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Nickname       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	HeadImgUrl     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	UnionId        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupId        int    `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        uint64 `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

type CUserInfoMultiQuery struct {
	OpenId string `json:"openid"`
	Lang   string `json:"lang"`
}

type CGetUserBaseInfoMultiRequest struct {
	UserList []CUserInfoMultiQuery `json:"user_list"`
}

type CUserBaseInfo struct {
	Subscribe      int    `json:"subscribe"`
	Openid         string `json:"openid"`
	Nickname       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	HeadImgUrl     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	UnionId        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupId        int    `json:"groupid"`
	TagidList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        uint64 `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

type CGetUserBaseInfoMultiResponse struct {
	ErrCode      int             `json:"errcode"`
	ErrMsg       string          `json:"errmsg"`
	UserInfoList []CUserBaseInfo `json:"user_info_list"`
}

type CGetFollowUsersResponse struct {
	OpenIds []string
}
