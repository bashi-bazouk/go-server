package services

import (
	"net/http"
	"fmt"
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
)




func Trace(handlerFunc http.HandlerFunc) http.HandlerFunc {
	// wrap a handler function to see some of its input.
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("traced: %s", r.URL)
		handlerFunc(w, r)
	}
}


// Simple Handlers


func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	myItems := []string{"item1", "item2", "item3"}
	a, _ := json.Marshal(myItems)

	w.Write(a)
	return
}

var WebsocketHandler = websocket.Handler(func (ws *websocket.Conn) {
  var s string
  fmt.Fscan(ws, &s)
  fmt.Println("Received: ", s)
})
