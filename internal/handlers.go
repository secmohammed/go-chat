package internal

import (
    "fmt"

    "github.com/mitchellh/mapstructure"
)

//Channel struct.
type Channel struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

//AddChannel is used to add a channel for the chat rooms.
func AddChannel(client *Client, data interface{}) {
    var channel Channel
    var message Message
    mapstructure.Decode(data, &channel)
    fmt.Printf("%#v\n", channel)
    channel.ID = "ABC"
    message.Name = "channel add"
    message.Data = channel
    client.send <- message
}
