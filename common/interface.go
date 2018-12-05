package common

type FuncMsgCallback func(reply IReply, msg *CMessage, communicate CDataCommunicate, userData interface{}) error
type FuncEventCallback func(reply IReply, event *CEvent, communicate CDataCommunicate, userData interface{}) error

type IEvent interface {
	OnEvent(sender IReply, event *CEvent, communicate CDataCommunicate, userData interface{}) error
}

type IMessage interface {
	OnMessage(sender IReply, msg *CMessage, communicate CDataCommunicate, userData interface{}) error
}

type IReply interface {
	SendMessage(msg *CMessage) error
	SendEmptyMessage() error
	IsSend() bool
}

type IToken interface {
	GetToken(timeoutMS int64) (token []byte, e error)
	UpdateToken(timeoutMS int64) (token []byte, e error)
}

type IMenu interface {
	Create(data *[]CButton, timeoutMS int64) error
	GetAll(timeoutMS int64) (*CGetMenuJson, error)
	DeleteAll(timeoutMS int64) error
}

type ITemplate interface {
	SetIndustry(request *CSetIndustryRequest, timeoutMS int64) error
	GetIndustry(timeoutMS int64) (*CGetIndustryResponse, error)
	GetTemplateId(request *CGetTemplateIdRequest, timeoutMS int64) (*CGetTemplateIdResponse, error)
	GetTemplateList(timeoutMS int64) (*CGetTemplateListResponse, error)
	DeleteTemplate(request *CDeleteTemplateRequest, timeoutMS int64) error
	SendTemplateMsg(request *CSendTemplateMsgRequest, timeoutMS int64) (*CSendTemplateMsgResponse, error)
}

type IMaterial interface {
	UploadImage(path *string, timeoutMS int64) (*CUploadImageResponse, error)
}

type IStore interface {
	CreateStore(request *CCreateStoreRequest, timeoutMS int64) error
	GetStore(request *CGetStoreRequest, timeoutMS int64) (*CGetStoreResponse, error)
	GetStoreList(request *CGetStoreListRequest, timeoutMS int64) (*CGetStoreListResponse, error)
	ModifyStore(request *CModifyStoreRequest, timeoutMS int64) error
	DeleteStore(request *CDeleteStoreRequest, timeoutMS int64) error
	GetCategory(timeoutMS int64) (*CGetCategoryResponse, error)
}
