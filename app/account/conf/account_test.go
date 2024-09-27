package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccountConfig(t *testing.T) {
	path := "./"
	name := "test"

	_, err := NewAccountConfig(path, name)
	assert.Equal(t, nil, err)
}
