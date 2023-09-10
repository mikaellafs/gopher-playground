package async

import (
	"log"
	"sync"
)

type Broadcaster struct {
	write    *sync.Mutex
	clients  map[string]chan []byte
	Receiver chan []byte
}

func NewBroadcaster(data chan []byte) *Broadcaster {
	return &Broadcaster{
		write:    &sync.Mutex{},
		clients:  map[string]chan []byte{},
		Receiver: data,
	}
}

func (b *Broadcaster) Register(id string) chan []byte {
	b.write.Lock()
	defer b.write.Unlock()

	b.clients[id] = make(chan []byte)
	return b.clients[id]
}

func (b *Broadcaster) Unregister(id string) {
	b.write.Lock()
	defer b.write.Unlock()

	b.clients[id] = nil
}

func (b *Broadcaster) Send(data []byte) {
	log.Println("Sending data to broadcaster routine")
	b.Receiver <- data
}

func (b *Broadcaster) Get(id string) chan []byte {
	b.write.Lock()
	defer b.write.Unlock()

	return b.clients[id]
}

func (b *Broadcaster) Start(data []byte) error {
	log.Println("Broadcasting data:", string(data))

	for _, client := range b.clients {
		client <- data
	}

	return nil
}
