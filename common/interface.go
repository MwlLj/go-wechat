package common

type IEvent interface {
}

type IMessage interface {
	OnMessage(sender IReply, msg *CMessage, userData interface{}) error
}

type IReply interface {
	SendMessage(msg *CMessage) error
}
