package internal

import (
    "github.com/gorilla/websocket"
)

// FindHandler is used to define the type of function with its returning vlaues.
type FindHandler func(string) (Handler, bool)

// Client is used to define the message sending using the socket connection and its appropriate handler.
type Client struct {
    send        chan Message
    socket      *websocket.Conn
    findHandler FindHandler
}

//Message payload
type Message struct {
    Name string      `json:"name"`
    Data interface{} `json:"data"`
}

// Read is used to read from the websocket conneciton.
func (client *Client) Read() {
    var message Message
    for {
        if err := client.socket.ReadJSON(&message); err != nil {
            break
        }
        if handler, found := client.findHandler(message.Name); found {
            handler(client, message.Data)
        }
    }
    client.socket.Close()
}

// Write is used to write to the websocket connection
func (client *Client) Write() {
    for msg := range client.send {
        if err := client.socket.WriteJSON(msg); err != nil {
            break
        }

    }
    client.socket.Close()
}

//NewClient is used to initalize the client websocket connection.
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
    return &Client{
        send:        make(chan Message),
        socket:      socket,
        findHandler: findHandler,
    }
}
