# Haproxy

Go module for [Haproxy](https://www.haproxy.com)

### Install package

``` bash
go get -u github.com/eosswedenorg-go/haproxy@latest
```

### Agent check module

```go
import "github.com/eosswedenorg-go/haproxy/agentcheck"
```

This module exposes a simple API to deal with HAproxy's auxiliary agent check via TCP Messages.

Read the [official documentation](http://docs.haproxy.org/2.4/configuration.html#5.2-agent-check)
for more information about what effect each type of message have on the HAproxy server.

Include the following line in your go file:

#### Interface

All responses implements this simple interface.
```go
type Response interface {
    // Returns the TCP Message to send back to haproxy.
    String() string
}
```

#### Weight Response

```go
func NewWeightResponse(percent int) WeigthResponse
```

Construct a weight response message, `percent` should be a value between 0-100 and indicates to HAproxy to
set the server's weight proportional to the initial weight when HAproxy first started.

For example: if a server is set to `200` on startup and a message with `75` percent is received from the agent
Haproxy will update the server's weight to `150`


#### Max Connections Response

```go
func NewMaxConnResponse(value int) MaxConnResponse
```

Construct a Max Connections Response. This response tells HAproxy to set `max connections`
of the server to the value specified.


#### Status Response

```go
func NewStatusResponse(status StatusResponseType) StatusResponse
```
Constructs a status response. `status` can be one of the following constants `Up`, `Maint`, `Ready` or `Drain`

#### Status Message Response

```go
func NewStatusMessageResponse(status StatusMessageResponseType, msg string) StatusMessageResponse
```
Constructs a status message response. Same as `StatusResponse` but with an optinal message `msg`.
To omit the message just pass a empty string.

`status` can be one of the following constants `Down`, `Failed` or `Stopped`.

#### Example

```go
package main

import (
    "fmt"
	"log"
	"net"
    "github.com/eosswedenorg-go/haproxy/agentcheck"
)

func onConnection(c net.Conn) {

    // Perform whatever logic is needed to check the status of the server.

    // in this example, we just tell HAproxy that the server is down with some helpful message.
    r := agentcheck.NewStatusMessageResponse(agentcheck.Down, "process is not running")

    // Send the message
    fmt.Fprint(c, r)

    // Shut down the connection.
    c.Close()
}

// TCP Code taken from: https://pkg.go.dev/net#example-Listener
func main() {

	// Listen on TCP port 2000 on localhost
	l, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go onConnection(conn)
	}
}
```

### Author

Henrik Hautakoski - [Sw/eden](https://eossweden.org/) - [henrik@eossweden.org](mailto:henrik@eossweden.org)
