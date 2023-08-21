package main

import (
	"time"
	"net"
	"log"
	cache "github.com/mukeshmahato18/cache/Cache"
	)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(2 * time.Second)
		conn, err := net.Dial("tcp", ":3000")
		if err!= nil {
			log.Fatal(err)
		}
	}

	server := NewServer(opts, cache.NewCacher())
	server.Start()
}
