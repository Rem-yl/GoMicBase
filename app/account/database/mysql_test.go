package database

import (
	"GoMicBase/pkg/cfg"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMysqlDB(t *testing.T) {
	config := &cfg.MysqlConfig{
		TableName: "test",
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Password:  "123456",
	}

	_, err := NewMysqlDB(config)
	assert.Equal(t, nil, err)

	config = &cfg.MysqlConfig{
		TableName: "notest",
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Password:  "123456",
	}

	_, err = NewMysqlDB(config)
	assert.NotEqual(t, nil, err)
}
