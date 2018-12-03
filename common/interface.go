package common

type IEvent interface {
}

type IMessage interface {
	OnMessage(sender ISender, msg *CMessage) error
}

type ISender interface {
	SendMessage(msg *CMessage) error
}
