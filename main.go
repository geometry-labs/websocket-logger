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

    _, data, err := conn.ReadMessage()
    if err != nil {
      log.Println(err)
      return
    }

    // log
    log.Printf("%s - %s", r.RemoteAddr, string(data))
  }

  return
}

func main() {

  // allow all origins
  upgrader.CheckOrigin = func(r *http.Request) bool {
    // if req.Header.Get("Origin") != "http://"+req.Host {
	  //   http.Error(w, "Origin not allowed", http.StatusForbidden)
	  //   return
    // }
    return true
  }

  http.HandleFunc("/ws", websocketLogger)

  http.ListenAndServe(":8080", nil)

  return
}
