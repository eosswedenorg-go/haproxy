
package agentcheck

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestWeight(t *testing.T) {

    w := NewWeightResponse(62)

    assert.Equal(t, "62%\n", w.String())
}

func TestMaxConn(t *testing.T) {

    c := NewMaxConnResponse(60)

    assert.Equal(t, "maxconn:60\n", c.String())
}

func TestStatus(t *testing.T) {

    s := NewStatusResponse(Maint)

    assert.Equal(t, "maint\n", s.String())

    s = NewStatusResponse(Drain)

    assert.Equal(t, "drain\n", s.String())
}


func TestStatusMessage(t *testing.T) {

    s := NewStatusMessageResponse(Down, "")

    assert.Equal(t, "down\n", s.String())

    s = NewStatusMessageResponse(Failed, "some error")

    assert.Equal(t, "failed#some error\n", s.String())
}
