package processor

import (
	"encoding/json"
	"gopher-playground/chatify/pkg/message"
	"log"

	"github.com/pkg/errors"
)

type Formatter struct {
	format func([]byte) (message.Message, error)
	next   Handler
}

func NewMessageFormatter(format func([]byte) (message.Message, error)) *Formatter {
	return &Formatter{
		format: format,
	}
}

func NewDefaultFormatter() *Formatter {
	return &Formatter{
		format: func(d []byte) (message.Message, error) {
			var msg message.BaseMessage
			err := json.Unmarshal(d, &msg)
			return &msg, err
		},
	}
}

// Manipulate data before broadcasting
func (f *Formatter) Process(c *MContext) error {
	msg, err := f.format(c.RawData)
	if err != nil {
		err = errors.Wrap(err, "failed to manipulate data")
		log.Println(err.Error())
		return err
	}

	c.ParsedData = msg
	return nil
}

func (f *Formatter) Next() Handler {
	return f.next
}

func (f *Formatter) SetNext(h Handler) {
	f.next = h
}
