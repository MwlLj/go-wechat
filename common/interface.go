package common

type IMessageCallback interface {
	OnMessage(msg *string)
}
