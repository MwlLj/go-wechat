package common

type FuncMsgCallback func(reply IReply, msg *CMessage, userData interface{}) error

type IEvent interface {
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
