// See https://cbonte.github.io/haproxy-dconv/1.7/configuration.html#5.2-agent-check

package agentcheck

import "fmt"

type Response interface {
    String() string
}

//  Weight setting
// -----------------------

func NewWeightResponse(p int) WeigthResponse {
    return WeigthResponse{Percentage: p}
}

type WeigthResponse struct {
    Percentage int
}

func (s WeigthResponse) String() string {
    return fmt.Sprintf("%d%%\n", s.Percentage)
}

//  Max connections
// -----------------------

func NewMaxConnResponse(value int) MaxConnResponse {
    return MaxConnResponse{Value: value}
}

type MaxConnResponse struct {
    Value int
}

func (s MaxConnResponse) String() string {
    return fmt.Sprintf("maxconn:%d\n", s.Value)
}

//  Status
// -----------------------

func NewStatusResponse(s StatusResponseType) StatusResponse {
    return StatusResponse{Status: s}
}

type StatusResponseType string
const (
    Up       StatusResponseType = "up"
    Maint    StatusResponseType = "maint"
    Ready    StatusResponseType = "ready"
    Drain    StatusResponseType = "drain"
)

type StatusResponse struct {
    Status StatusResponseType
}

func (s StatusResponse) String() string {
    return string(s.Status) + "\n"
}


//  Status Message
// -----------------------

func NewStatusMessageResponse(s StatusMessageResponseType, msg string) StatusMessageResponse {
    return StatusMessageResponse{Status: s, Message: msg}
}

type StatusMessageResponseType string
const (
    Down     StatusMessageResponseType = "down"
    Failed   StatusMessageResponseType = "fail"
    Stopped  StatusMessageResponseType = "stopped"
)

type StatusMessageResponse struct {
    Status StatusMessageResponseType
    Message string
}

func (s StatusMessageResponse) String() string {
    if len(s.Message) < 1 {
        return string(s.Status) + "\n"
    }
    return fmt.Sprintf("%s#%s\n", s.Status, s.Message)
}
