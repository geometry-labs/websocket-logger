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

  WS_LOGGER_FORWARD_ADDR := os.Getenv("WS_LOGGER_FORWARD_ADDR")
  if WS_LOGGER_FORWARD_ADDR == "" {
    WS_LOGGER_FORWARD_ADDR = "ws://localhost:9090/"
  }

  WS_LOGGER_LOG_CLIENT := os.Getenv("WS_LOGGER_LOG_CLIENT")
  if WS_LOGGER_LOG_CLIENT == "" {
    WS_LOGGER_LOG_CLIENT = "true"
  }

  WS_LOGGER_LOG_SERVER := os.Getenv("WS_LOGGER_LOG_SERVER")
  if WS_LOGGER_LOG_SERVER == "" {
    WS_LOGGER_LOG_SERVER = "true"
  }

  // parse forward address
  fwd, err := url.Parse(WS_LOGGER_FORWARD_ADDR)
  if err != nil {
    log.Fatalf("ERROR: forward url cannot be parsed: %s\n", WS_LOGGER_FORWARD_ADDR)
  }

  // default handler to connect
  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){return})

  err = http.ListenAndServe(
    fmt.Sprintf(":%s", WS_LOGGER_EXPOSED_PORT),
    websocketproxy.NewProxy(fwd, WS_LOGGER_LOG_CLIENT == "true", WS_LOGGER_LOG_SERVER == "true"),
  )
  if err != nil {
    log.Fatalln(err)
  }

  return
}
