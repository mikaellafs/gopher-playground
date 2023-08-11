package chatify

import (
	"context"
	"encoding/json"
	"gopher-playground/chatify/internal/async"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Client struct {
	conn      *websocket.Conn
	broadcast chan []byte
	format    func([]byte) (message, error)
}

func NewClient(conn *websocket.Conn, broadcast chan []byte, format func([]byte) (message, error)) *Client {
	log.Println("New client connected...")

	return &Client{
		conn:      conn,
		broadcast: broadcast,
		format:    format,
	}
}

func (c *Client) Start() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(2)
	go async.RunEndlessRoutineWithCancel(&wg, ctx, cancel, c.handleIncomingMessage, time.Second)
	go async.RunEndlessRoutineWithCancel(&wg, ctx, cancel, c.handleBroadcastMessages, time.Second)

	wg.Wait()
	log.Println("Ending client connection...")
}

// Deal with messages received from client
func (c *Client) handleIncomingMessage() error {
	_, data, err := c.conn.ReadMessage()
	if err != nil {
		err = errors.Wrap(err, "failed to read message from web socket connection")
		log.Println(err.Error())
		return err
	}

	// Manipulate data before broadcasting
	msg, err := c.format(data)
	if err != nil {
		err = errors.Wrap(err, "failed to manipulate data")
		log.Println(err.Error())
		return err
	}

	// Marshal back to bytes
	data, _ = json.Marshal(msg)

	// Broadcast message to all clients
	c.broadcast <- data
	return nil
}

// Handle messages that comes from broadcast channel. That is, messages sent by other users
func (c *Client) handleBroadcastMessages() error {
	select {
	case msg := <-c.broadcast:
		// Send message to client
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			err = errors.Wrap(err, "failed to write message through web socket connection")
			log.Println(err.Error())
			return err
		}
	default:
	}

	return nil
}
