package tui

import (
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

var wsConn *websocket.Conn
var wsUrl string

func connectToWebsocket(url string) error {
	log.Debug("Connecting to websocket at", "url", url)

	// Upgrade HTTP connection to WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}

	wsUrl = url
	wsConn = conn

	return nil
}

// `connectToWebsocket` needs to be called before this function
// to initialize `wsConn` and `wsUrl`
func sendMessage(msg string) error {
	// Check that the URL has been set
	if wsUrl == "" {
		log.Fatal("No websocket URL set!")
	}
	// Check that the connection is still open
	if wsConn == nil {
		log.Debug("Not connected to websocket, reconnecting...")
		err := connectToWebsocket(wsUrl)
		if err != nil {
			log.Error("Failed to reconnect from sendMessage")
			return err
		}
	}

	// Write message to WebSocket
	log.Debug("Sending message", "msg", msg)
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
