# websocket-logger

### Quick Start

```
go build
./websocket-logger
```

### Test

```
go test ./...
```
 
### Enviroment Variables
| Variables | Default | Description |
| :--- | :--- | :--- | 
| WS_LOGGER_EXPOSED_PORT | 8080 | port exposed to upgrade and log websocket connections |
| WS_LOGGER_FORWARD_ADDR | ws://localhost:9090/ | address of load balancer to forward all websocket messages to |
| WS_LOGGER_LOG_CLIENT | true | output websocket messages to stdout coming from client |
| WS_LOGGER_LOG_SERVER | true | output websocket messages to stdout coming from server |
