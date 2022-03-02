// See https://cbonte.github.io/haproxy-dconv/1.7/configuration.html#5.2-agent-check

package haproxy

import "fmt"

type HealthCheck interface {
    String() string
}

//  Weight setting
// -----------------------

func Weight(p int) HealthCheckWeigth {
    return HealthCheckWeigth{Percentage: p}
}

type HealthCheckWeigth struct {
    Percentage int
}

func (s HealthCheckWeigth) String() string {
    return fmt.Sprintf("%d%%\n", s.Percentage)
}

//  Max connections
// -----------------------

func MaxConn(value int) HealthCheckMaxConn {
    return HealthCheckMaxConn{Value: value}
}

type HealthCheckMaxConn struct {
    Value int
}

func (s HealthCheckMaxConn) String() string {
    return fmt.Sprintf("maxconn:%d\n", s.Value)
}

//  Status
// -----------------------

func Status(s HealthCheckStatusType) HealthCheckStatus {
    return HealthCheckStatus{Status: s}
}

type HealthCheckStatusType string
const (
    HealthCheckUp       HealthCheckStatusType = "up"
    HealthCheckMaint    HealthCheckStatusType = "maint"
    HealthCheckReady    HealthCheckStatusType = "ready"
    HealthCheckDrain    HealthCheckStatusType = "drain"
)

type HealthCheckStatus struct {
    Status HealthCheckStatusType
}

func (s HealthCheckStatus) String() string {
    return string(s.Status) + "\n"
}


//  Status Message
// -----------------------

func StatusMessage(s HealthCheckStatusMessageType, msg string) HealthCheckStatusMessage {
    return HealthCheckStatusMessage{Status: s, Message: msg}
}

type HealthCheckStatusMessageType string
const (
    HealthCheckDown     HealthCheckStatusMessageType = "down"
    HealthCheckFailed   HealthCheckStatusMessageType = "failed"
    HealthCheckStopped  HealthCheckStatusMessageType = "stopped"
)

type HealthCheckStatusMessage struct {
    Status HealthCheckStatusMessageType
    Message string
}

func (s HealthCheckStatusMessage) String() string {
    if len(s.Message) < 1 {
        return string(s.Status) + "\n"
    }
    return fmt.Sprintf("%s#%s\n", s.Status, s.Message)
}
