package main

import (
	"fmt"
	"testing"
)

func TestAddUsers(t *testing.T) {
	server := NewServer()

	for i := 0; i < 10; i++ {
		server.userch <- fmt.Sprintf("user_%d", i)
	}
	fmt.Println("the loop is done!")
}
