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
	MaterialTypeImage   string = "image"
	MaterialTypeVoice   string = "voice"
	MaterialTypeVideo   string = "video"
	MaterialTypeThumb   string = "thumb"
	MaterialTypeImgText string = "news"
)

var (
	GroupSendMsgTypeImgText string = "mpnews"
	GroupSendMsgTypeText    string = "text"
	GroupSendMsgTypeVoice   string = "voice"
	GroupSendMsgTypeImage   string = "image"
	GroupSendMsgTypeVideo   string = "mpvideo"
	GroupSendMsgTypeCard    string = "wxcard"
)

var (
	// position event
	EventTypeLocation string = "LOCATION"
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

type CCreateTmpMaterialRequest struct {
	MaterialType string
	Path         string
}

type CCreateTmpMaterialResponse struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

type CGetTmpMaterialRequest struct {
	MediaId string
}

type CGetTmpMaterialResponse struct {
	ResUrl string
}

type CGetTmpHDMaterialRequest struct {
	MediaId string
}

type CGetTmpHDMaterialResponse struct {
	ResUrl string
}

type CImgTextMaterial struct {
	Title              string `json:"title"`
	ThumbMediaId       string `json:"thumb_media_id"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	Content            string `json:"content"`
	ContentSourceUrl   string `json:"content_source_url"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
	Url                string `json:"url"`
}

type CAddForeverImgTextMaterialRequest struct {
	Articles []CImgTextMaterial `json:"articles"`
}

type CAddForeverImgTextMaterialResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MediaId string `json:"media_id"`
}

type CAddForeverOtherMaterialRequest struct {
	MaterialType string
	Path         string
	Title        string
	Introduction string
}

type CAddForeverOtherMaterialResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MediaId string `json:"media_id"`
	Url     string `json:"url"`
}

type CUploadImageRequest struct {
	Media []byte `json:"media"`
}

type CUploadImageResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Url     string `json:"url"`
}

type CUploadVideoRequest struct {
	// get from addOtherMaterial interface
	MediaId     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CUploadVideoResponse struct {
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type CGetForeverMaterialRequest struct {
	MediaId string `json:"media_id"`
}

type CGetForeverMaterialResponse struct {
	NewsItem    []CImgTextMaterial `json:"news_item"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DownUrl     string             `json:"down_url"`
	ErrCode     int                `json:"errcode"`
	ErrMsg      string             `json:"errmsg"`
}

type CDeleteForeverMaterialRequest struct {
	MediaId string `json:"media_id"`
}

type CUpdateForeverImgTextMaterialRequest struct {
	MediaId  string           `json:"media_id"`
	Index    int              `json:"index"`
	Articles CImgTextMaterial `json:"articles"`
}

type CGetMaterialTotalResponse struct {
	VoiceCount int    `json:"voice_count"`
	VideoCount int    `json:"video_count"`
	ImageCount int    `json:"image_count"`
	NewsCount  int    `json:"news_count"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type CImgTextMaterialListContent struct {
	NewsItem []CImgTextMaterial `json:"news_item"`
}

type CMaterialListItem struct {
	MediaId    string                      `json:"media_id"`
	Content    CImgTextMaterialListContent `json:"content"`
	UpdateTime int64                       `json:"update_time"`
	Name       string                      `json:"name"`
	Url        string                      `json:"url"`
}

type CGetMaterialListRequest struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}

type CGetMaterialListResponse struct {
	TotalCount int                 `json:"total_count"`
	ItemCount  int                 `json:"item_count"`
	Item       []CMaterialListItem `json:"item"`
	ErrCode    int                 `json:"errcode"`
	ErrMsg     string              `json:"errmsg"`
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

type CGetBlackListUsersResponse struct {
	OpenIds []string
}

type CTakeUsersToBlackListRequest struct {
	OpenIdList []string `json:"openid_list"`
}

type CUnTakeUsersToBlackListRequest struct {
	OpenIdList []string `json:"openid_list"`
}

type CGroupSendFilterInfo struct {
	IsToAll bool `json:"is_to_all"`
	TagId   int  `json:"tag_id"`
}

type CGroupSendCardExt struct {
	Code      string `json:"code"`
	TimeStamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type CGroupSendMsgContent struct {
	MediaId     string            `json:"media_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Content     string            `json:"content"`
	CardId      string            `json:"card_id"`
	CardExt     CGroupSendCardExt `json:"card_ext"`
}

type CGroupSendByTagRequest struct {
	Filter            CGroupSendFilterInfo `json:"filter"`
	ImgText           CGroupSendMsgContent `json:"mpnews"`
	Text              CGroupSendMsgContent `json:"text"`
	Voice             CGroupSendMsgContent `json:"voice"`
	Image             CGroupSendMsgContent `json:"image"`
	Video             CGroupSendMsgContent `json:"mpvideo"`
	Card              CGroupSendMsgContent `json:"wxcard"`
	MsgType           string               `json:"msgtype"`
	SendIgnoreReprint int                  `json:"send_ignore_reprint"`
}

type CGroupSendByTagResponse struct {
	MsgId     int64  `json:"msg_id"`
	MsgDataId int64  `json:"msg_data_id"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type CGroupSendByOpenIdsRequest struct {
	Touser            []string             `json:"touser"`
	ImgText           CGroupSendMsgContent `json:"mpnews"`
	Text              CGroupSendMsgContent `json:"text"`
	Voice             CGroupSendMsgContent `json:"voice"`
	Image             CGroupSendMsgContent `json:"image"`
	Video             CGroupSendMsgContent `json:"mpvideo"`
	Card              CGroupSendMsgContent `json:"wxcard"`
	MsgType           string               `json:"msgtype"`
	SendIgnoreReprint int                  `json:"send_ignore_reprint"`
}

type CGroupSendByOpenIdsResponse struct {
	MsgId     int64  `json:"msg_id"`
	MsgDataId int64  `json:"msg_data_id"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type CDeleteGroupSendRequest struct {
	MsgId      int64 `json:"msg_id"`
	ArticleIdx int   `json:"article_idx"`
}

type CPreviewMessageRequest struct {
	ToWxName string               `json:"towxname"`
	Touser   string               `json:"touser"`
	ImgText  CGroupSendMsgContent `json:"mpnews"`
	Text     CGroupSendMsgContent `json:"text"`
	Voice    CGroupSendMsgContent `json:"voice"`
	Image    CGroupSendMsgContent `json:"image"`
	Video    CGroupSendMsgContent `json:"mpvideo"`
	Card     CGroupSendMsgContent `json:"wxcard"`
	MsgType  string               `json:"msgtype"`
}

type CPreviewMessageResponse struct {
	MsgId   int64  `json:"msg_id"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type CCommodityProperty struct {
	Id  string `json:"id"`
	Vid string `json:"vid"`
}

type CCommoditySkuInfo struct {
	Id  string   `json:"id"`
	Vid []string `json:"vid"`
}

type CCommodityProductBaseDetail struct {
	Text string `json:"text"`
	Img  string `json:"img`
}

type CCommodityProductBase struct {
	CategoryId []string                      `json:"category_id"`
	Property   []CCommodityProperty          `json:"property"`
	Name       string                        `json:"name"`
	SkuInfo    CCommoditySkuInfo             `json:"sku_info"`
	MainImg    string                        `json:"main_img"`
	DetailHtml string                        `json:"detail_html"`
	Img        []string                      `json:"img"`
	Detail     []CCommodityProductBaseDetail `json:"detail"`
	BuyLimit   int                           `json:"buy_limit"`
}

type CCommoditySkuListItem struct {
	SkuId       string `json:"sku_id"`
	Price       string `json:"price"`
	IconUrl     string `json:"icon_url"`
	ProductCode string `json:"product_code"`
	OriPrice    int64  `json:"ori_price"`
	Quantity    int    `json:"quantity"`
}

type CCommodityAttrExtLocation struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	Address  string `json:"address"`
}

type CCommodityAttrExt struct {
	Location         CCommodityAttrExtLocation `json:"location"`
	IsPostFree       int                       `json:"isPostFree"`
	IsHasReceipt     int                       `json:"isHasReceipt"`
	IsUnderGuaranty  int                       `json:"isUnderGuaranty"`
	IsSupportReplace int                       `json:"isSupportReplace"`
}

type CCommodityDeliveryInfoExpress struct {
	Id    int64 `json:"id"`
	Price int64 `json:"price"`
}

type CCommodityDeliveryInfo struct {
	DeliveryType int                             `json:"delivery_type"`
	TemplateId   int                             `json:"template_id"`
	Express      []CCommodityDeliveryInfoExpress `json:"express"`
}

type CAddCommodityRequest struct {
	ProductBase  CCommodityProductBase   `json:"product_base"`
	SkuList      []CCommoditySkuListItem `json:"sku_list"`
	AttrExt      CCommodityAttrExt       `json:"attrext"`
	DeliveryInfo CCommodityDeliveryInfo  `json:"delivery_info"`
}

type CAddCommodityResponse struct {
	ProductId string `json:"product_id"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

type CDeleteCommodityRequest struct {
	ProductId string `json:"product_id"`
}

type CUpdateCommodityRequest struct {
	ProductId    string                  `json:"product_id"`
	ProductBase  CCommodityProductBase   `json:"product_base"`
	SkuList      []CCommoditySkuListItem `json:"sku_list"`
	AttrExt      CCommodityAttrExt       `json:"attrext"`
	DeliveryInfo CCommodityDeliveryInfo  `json:"delivery_info"`
}

type CCommodityDetail struct {
	ProductBase  CCommodityProductBase   `json:"product_base"`
	SkuList      []CCommoditySkuListItem `json:"sku_list"`
	AttrExt      CCommodityAttrExt       `json:"attrext"`
	DeliveryInfo CCommodityDeliveryInfo  `json:"delivery_info"`
}

type CGetCommodityRequest struct {
	ProductId string `json:"product_id"`
}

type CGetCommodityResponse struct {
	ProductInfo CCommodityDetail `json:"product_info"`
	ErrCode     int              `json:"errcode"`
	ErrMsg      string           `json:"errmsg"`
}

type CGetCommodityByStatusRequest struct {
	CommodityStatus int `json:"status"`
}

type CGetCommodityByStatusResponse struct {
	ProductInfo []CCommodityDetail `json:"product_info"`
	ErrCode     int                `json:"errcode"`
	ErrMsg      string             `json:"errmsg"`
}

type CUpdateCommodityStatusRequest struct {
	ProductId       string `json:"product_id"`
	CommodityStatus int    `json:"status"`
}

type CGetSubClassesByClassifyRequest struct {
	CateId int64 `json:"cate_id"`
}

type CCommodityCateInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CGetSubClassesByClassifyResponse struct {
	CateList []CCommodityCateInfo `json:"cate_list"`
	ErrCode  int                  `json:"errcode"`
	ErrMsg   string               `json:"errmsg"`
}

type CGetAllSkuByClassifyRequest struct {
	CateId string `json:"cate_id"`
}

type CCommoditySkuValueList struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CCommoditySkuTable struct {
	Id        string                   `json:"id"`
	Name      string                   `json:"name"`
	ValueList []CCommoditySkuValueList `json:"value_list"`
}

type CGetAllSkuByClassifyResponse struct {
	SkuTable CCommoditySkuTable `json:"sku_table"`
	ErrCode  int                `json:"errcode"`
	ErrMsg   string             `json:"errmsg"`
}

type CGetAllPropertyByClassifyRequest struct {
	CateId string `json:"cate_id"`
}

type CCommodityClassifyPropertyValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CCommodityClassifyProperty struct {
	Id            string                            `json:"id"`
	Name          string                            `json:"name"`
	PropertyValue []CCommodityClassifyPropertyValue `json:"property_value"`
}

type CGetAllPropertyByClassifyResponse struct {
	Properties CCommodityClassifyProperty `json:"properties"`
	ErrCode    int                        `json:"errcode"`
	ErrMsg     string                     `json:"errmsg"`
}

type CStockSkuInfo struct {
	Id  string
	Vid string
}

type CAddStockRequest struct {
	ProductId string
	SkuInfo   []CStockSkuInfo
	Quantity  int
}

type CReduceStockRequest struct {
	ProductId string
	SkuInfo   []CStockSkuInfo
	Quantity  int
}

type CPayByPaymentCodeRequest struct {
	Key            string
	SignType       string
	MchId          string
	DeviceInfo     string
	Body           string
	Detail         string
	Attach         string
	OutTradeNo     string
	TotalFee       int
	FeeType        string
	SpbillCreateIp string
	GoodsTag       string
	LimitPay       string
	TimeStart      string
	TimeExpire     string
	Receipt        string
	AuthCode       string
	SceneInfo      string
}

type CPayByPaymentCodeResponse struct {
	XMLName            xml.Name `xml:"xml"`
	ReturnCode         string   `xml:"return_code"`
	ReturnMsg          string   `xml:"return_msg"`
	Appid              string   `xml:"appid"`
	MchId              string   `xml:"mch_id"`
	DeviceInfo         string   `xml:"device_info"`
	NonceStr           string   `xml:"nonce_str"`
	Sign               string   `xml:"sign"`
	ResultCode         string   `xml:"result_code"`
	ErrCode            string   `xml:"err_code"`
	ErrCodeDes         string   `xml:"err_code_des"`
	OpenId             string   `xml:"openid"`
	IsSubscribe        string   `xml:"is_subscribe"`
	TradeType          string   `xml:"trade_type"`
	TradeState         string   `xml:"trade_state"`
	BankType           string   `xml:"bank_type"`
	FeeType            string   `xml:"fee_type"`
	TotalFee           int      `xml:"total_fee"`
	SettlementTotalFee int      `xml:"settlement_total_fee"`
	CouponFee          int      `xml:"coupon_fee"`
	CouponCount        int      `xml:"coupon_count"`
	CouponType         int      `xml:"coupon_type_$n"`
	CashFeeType        string   `xml:"cash_fee_type"`
	CashFee            int      `xml:"cash_fee"`
	TransactionId      string   `xml:"transaction_id"`
	OutTradeNo         string   `xml:"out_trade_no"`
	Attach             string   `xml:"attach"`
	TimeEnd            string   `xml:"time_end"`
	PromotionDetail    string   `xml:"promotion_detail"`
}

type CQueryOrderRequest struct {
	MchId         string
	TransactionId string
	OutTradeNo    string
	SignType      string
}
