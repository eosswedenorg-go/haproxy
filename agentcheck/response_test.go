package agentcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgentcheckResponse_Weight(t *testing.T) {
	w := NewWeightResponse(62)

	assert.Equal(t, "62%\n", w.String())
}

func TestAgentcheckResponse_MaxConn(t *testing.T) {
	c := NewMaxConnResponse(60)

	assert.Equal(t, "maxconn:60\n", c.String())
}

func TestAgentcheckResponse_Status(t *testing.T) {
	s := NewStatusResponse(Maint)

	assert.Equal(t, "maint\n", s.String())

	s = NewStatusResponse(Drain)

	assert.Equal(t, "drain\n", s.String())
}

func TestAgentcheckResponse_StatusMessage(t *testing.T) {
	s := NewStatusMessageResponse(Down, "")

	assert.Equal(t, "down\n", s.String())

	s = NewStatusMessageResponse(Down, "Some other error")

	assert.Equal(t, "down#Some other error\n", s.String())

	s = NewStatusMessageResponse(Fail, "some error")

	assert.Equal(t, "fail#some error\n", s.String())
}
