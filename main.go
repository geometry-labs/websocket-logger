package main

import (
  "log"
  "net/http"

  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// endpoint: /ws
func websocketLogger(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println(err)
    return
  }

  for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
      log.Println(err)
      return
    }
    if err := conn.WriteMessage(messageType, p); err != nil {
      log.Println(err)
      return
    }
  }

  return
}

func main() {

  // allow all origins
  upgrader.CheckOrigin = func(r *http.Request) bool { return true }

  http.HandleFunc("/ws", websocketLogger)

  http.ListenAndServe(":8080", nil)

  return
}
