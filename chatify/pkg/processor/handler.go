package processor

import (
	"gopher-playground/chatify/pkg/message"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Process(*MContext) error
	Next() Handler
	SetNext(Handler)
}

type MContext struct {
	ginCtx     *gin.Context
	RawData    []byte
	ParsedData message.Message
}

func NewContext(c *gin.Context, data []byte) *MContext {
	return &MContext{
		ginCtx:  c,
		RawData: data,
	}
}
