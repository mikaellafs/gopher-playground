package connection

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWSUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections by not checking the origin.
			return true
		},
	}
}
