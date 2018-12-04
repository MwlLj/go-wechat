package common

type FuncMsgCallback func(reply IReply, msg *CMessage, userData interface{}) error
type FuncEventCallback func(reply IReply, event *CEvent, userData interface{}) error

type IEvent interface {
	OnEvent(sender IReply, event *CEvent, userData interface{}) error
}

type IMessage interface {
	OnMessage(sender IReply, msg *CMessage, userData interface{}) error
}

type IReply interface {
	SendMessage(msg *CMessage) error
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
