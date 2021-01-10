package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

//Handler is declaring the type of our handler.
type Handler func(*Client, interface{})

// Router is used to define the structure of router and its rules.
type Router struct {
	rules map[string]Handler
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//FindHandler is used to find the proper handler to handle the incoming event from websocket.
func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found

}

//NewRouter is used to initalize the router
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

//Handle is used to setup the msg name with its own handler
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

//ServeHTTP is used to start serving http server.
func (r *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	socket, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(response, err.Error())
		return
	}

	client := NewClient(socket, r.FindHandler)
	go client.Write()
	client.Read()

}
