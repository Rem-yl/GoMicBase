package internal

import (
	"fmt"
	"testing"
)

func TestGetNewUuid(t *testing.T) {
	for i := 0; i < 10; i++ {
		uid := GetNewUuid()
		fmt.Printf("uuid: %s\n", uid)
	}
}

func TestGetRandomPort(t *testing.T) {
	for i := 0; i < 10; i++ {
		port := GetRandomPort("10.7.9.248")
		fmt.Printf("port: %d\n", port)
	}
}
