package common

type IEvent interface {
}

type IMessage interface {
	OnMessage(sender ISender, msg *CMessage, userData interface{}) error
}

type ISender interface {
	SendMessage(msg *CMessage) error
}
