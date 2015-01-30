package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var chttp *http.ServeMux

func main() {
	fmt.Println("Hello World !")

	chttp = http.NewServeMux()

	chttp.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, ".") {
		chttp.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not found", 404)
	}
}

type chatMsg struct {
	Level     string `json:"level,omitempty"`
	Msg       string `json:"msg,omitempty"`
	Ttl       int64  `json:"ttl,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Read      bool   `json:"read"`
	Seen      bool   `json:"seen"`
}

type ChatSocket struct {
	upgrader *websocket.Upgrader
}

func NewChatSocket() *ChatSocket {
	var chatSocket = &ChatSocket{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}

	return chatSocket
}

func (cs *ChatSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := cs.upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if _, ok := err.(websocket.HandshakeError); ok {
		wsHandleError(cs, w, conn, fmt.Errorf("Not a websocket handshake"))
		return
	} else if err != nil {
		wsHandleError(cs, w, conn, fmt.Errorf("Connection error"))
		return
	}

	pingBytes := 0

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			wsHandleError(cs, w, conn, fmt.Errorf("Unreadable message"))
			return
		}

		if msgType == websocket.TextMessage {
			d := json.NewDecoder(bytes.NewReader(msg))
			var msg chatMsg
			if decodeErr := d.Decode(&msg); decodeErr != nil {
				wsHandleError(cs, w, conn, fmt.Errorf("Unreadable JSON message %s", decodeErr.Error()))
				return
			}
			log.Println("Socket Received message", string(msg))
			var b bytes.Buffer
			encoder := json.NewEncoder(&b)
			encodeErr := encoder.Encode(msg)
			if err == nil {
				conn.WriteMessage(websocket.TextMessage, b)
			}
		} else {
			pingBytes += len(msg)
		}
	}
}

func wsHandleError(cs *ChatSocket, w http.ResponseWriter, conn *websocket.Conn, err error) {
	http.Error(w, err.Error(), 400)
	log.Println("Websocket connection error : ", err)
}
