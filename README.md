# websocket-logger

### Quick Start

```
go build
./websocket-logger
```

### Enviroment Variables
| Variables | Default | Description |
| :--- | :--- | :--- | 
| WS_LOGGER_EXPOSED_PORT | 8080 | port exposed to upgrade and log websocket connections |
| LB_FORWARD_ADDR | ??? | address of load balancer to forward all websocket messages to |
