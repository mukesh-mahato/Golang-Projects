package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
	// you can name you for loop
free:
	for {
		select {
		// block here until someone is sending message to the channel
		case msg := <-s.msgch:
			fmt.Printf("recived message from: %s payload %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("server is doing a graceful shut down")
			// gogic for graceful shut down
			break free
		default:
		}
	}
	fmt.Println("the server is shut down")

}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Joy",
		Payload: payload,
	}
	msgch <- msg
}

func gracefulQuitServer(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "hey")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		gracefulQuitServer(s.quitch)
	}()

	select {}
}

// func main() {
// 	now := time.Now()
// 	respch := make(chan string, 128)
// 	wg := &sync.WaitGroup{}

// 	userID := 10
// 	go fetchUserData(userID, respch, wg)
// 	go fetchUserRecomendation(userID, respch, wg)
// 	go fetchUserLikes(userID, respch, wg)
// 	wg.Add(3)
// 	wg.Wait()

// 	close(respch)

// 	for resp := range respch {
// 		fmt.Println(resp)
// 	}

// 	fmt.Println(time.Since(now))
// }

// func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)

// 	respch <- "user data"

// 	wg.Done()
// }

// func fetchUserRecomendation(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(120 * time.Millisecond)

// 	respch <- "user recommendations"

// 	wg.Done()
// }

// func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(50 * time.Millisecond)

// 	respch <- "user likes"

// 	wg.Done()
// }
