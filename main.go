package main

import (
  "log"
  "fmt"
  "os"
  "time"
  "net/http"
  "net/url"

  "insight-infrastructure/websocket-logger/websocketproxy"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
  return fmt.Print(time.Now().Format(time.RFC3339) + " - " + string(bytes))
}

func main() {
  // set custom log output
  log.SetFlags(0)
  log.SetOutput(new(logWriter))

  log.Println("Starting proxy...")

  // Get env vars
  WS_LOGGER_EXPOSED_PORT := os.Getenv("WS_LOGGER_EXPOSED_PORT")
  if WS_LOGGER_EXPOSED_PORT == "" {
    WS_LOGGER_EXPOSED_PORT = "8080"
  }

  LB_FORWARD_ADDR, err := url.Parse(os.Getenv("LB_FORWARD_ADDR"))
  if err != nil {
    log.Fatalln("Invalid load balancer url")
  }

  // default handler to connect
  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){return})

  err = http.ListenAndServe(
    fmt.Sprintf(":%s", WS_LOGGER_EXPOSED_PORT),
    websocketproxy.NewProxy(LB_FORWARD_ADDR),
  )
  if err != nil {
    log.Fatalln(err)
  }

  return
}
