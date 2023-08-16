package processor

import "gopher-playground/chatify/pkg/message"

type Handler interface {
	Process(*MContext) error
	Next() Handler
	SetNext(Handler)
}

type MContext struct {
	RawData    []byte
	ParsedData message.Message
}

func NewContext(data []byte) *MContext {
	return &MContext{
		RawData: data,
	}
}
