package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConsulRegistery(t *testing.T) {
	var host string = "127.0.0.1"
	var port int32 = 8080

	c, err := NewConsulRegistery(host, port)
	if err != nil {
		t.Error(err.Error())
	}

	assert.IsType(t, c, &ConsulRegistery{})
}
