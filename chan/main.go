package main

import "fmt"

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.loop()
}

func (s *Server) loop() {
	for {
		select {
		case msg := <-s.userch:
			fmt.Printf(msg)
		case <-s.quitch:
			return
		}
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func main() {

}

func sendMessage(msgch chan<- string) {
	msgch <- "hello!"
}

func readMessage(msgch <-chan string) {
	msg := <-msgch
	fmt.Println(msg)
}
