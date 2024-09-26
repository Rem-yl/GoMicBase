package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadYamlConfig(t *testing.T) {
	var host, namespaceId, dataId, group string
	var port int
	config, err := LoadYamlConfig(".", "dev")
	if err != nil {
		t.Error(err.Error())
	}

	host = config.Get("host").(string)
	port = config.Get("port").(int)
	namespaceId = config.Get("namespaceId").(string)
	dataId = config.Get("dataId").(string)
	group = config.Get("group").(string)

	assert.Equal(t, host, "127.0.0.1")
	assert.Equal(t, port, 8848)
	assert.Equal(t, namespaceId, "public")
	assert.Equal(t, dataId, "test.json")
	assert.Equal(t, group, "test")
}
