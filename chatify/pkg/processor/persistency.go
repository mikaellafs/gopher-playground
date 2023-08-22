package processor

import (
	"log"

	"gopher-playground/chatify/pkg/message"

	"github.com/pkg/errors"
)

type Persistency struct {
	next  Handler
	store message.Store
}

func NewMessagePersister(store message.Store) *Persistency {
	return &Persistency{
		store: store,
	}
}

// Save message in some storage
func (p *Persistency) Process(c *MContext) error {
	err := p.store.SaveMessage(c.ParsedData)
	if err != nil {
		err = errors.Wrap(err, "failed to persist message")
		log.Println(err.Error())
		return err
	}

	return nil
}

func (p *Persistency) Next() Handler {
	return p.next
}

func (p *Persistency) SetNext(h Handler) {
	p.next = h
}
