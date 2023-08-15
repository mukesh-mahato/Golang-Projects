package main

import (
	"fmt"
	"testing"
)

func TestAddUsers(t *testing.T) {
	server := NewServer()

	for i := 0; i < 10; i++ {
		// go func(i int) {
		server.userch <- fmt.Sprintf("user_%d", i)
		// server.addUser(fmt.Sprintf("user_%d", i))
		// }(i)
	}
	fmt.Println("the loop is done!")
}
