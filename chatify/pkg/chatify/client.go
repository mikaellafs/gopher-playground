package chatify

import (
	"context"
	"gopher-playground/chatify/internal/async"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Client struct {
	conn      *websocket.Conn
	broadcast chan Message
}

func NewClient(conn *websocket.Conn, broadcast chan Message) *Client {
	log.Println("New client connected...")

	return &Client{
		conn:      conn,
		broadcast: broadcast,
	}
}

func (c *Client) Start() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(2)
	go async.RunEndlessRoutineWithCancel(&wg, ctx, cancel, c.handleIncomingMessage)
	go async.RunEndlessRoutineWithCancel(&wg, ctx, cancel, c.handleBroadcastMessages)

	wg.Wait()
}

// Deal with messages received from client
func (c *Client) handleIncomingMessage() error {
	log.Println("Handling incoming message...")

	_, data, err := c.conn.ReadMessage()
	if err != nil {
		err = errors.Wrap(err, "failed to read message from web socket connection")
		log.Println(err.Error())
		return err
	}

	// Broadcast message to all clients
	msg, err := NewMessageFrom(data)
	if err != nil {
		err = errors.Wrap(err, "failed to parse received message")
		log.Println(err.Error())
		return err
	}

	log.Printf("Message received from CLIENT %s: %s\n", msg.Username, msg.Text)

	c.broadcast <- msg
	return nil
}

// Handle messages that comes from broadcast channel. That is, messages sent by other users
func (c *Client) handleBroadcastMessages() error {
	log.Println("Handling broadcast messages...")

	msg := <-c.broadcast

	log.Printf("Message received from BROADCAST %s: %s\n", msg.Username, msg.Text)

	// Send message to client
	err := c.conn.WriteMessage(websocket.TextMessage, msg.ToBytes())
	if err != nil {
		err = errors.Wrap(err, "failed to write message through web socket connection")
		log.Println(err.Error())
		return err
	}

	return nil
}
