package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	cache "github.com/mukeshmahato18/cache/Cache"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOpts

	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error: %s", err)
	}
	log.Printf("server starting on port [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error %s\n", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("conn read error %s\n", err)
			break
		}
		go s.handleCommand(conn, buf[:n])
	}
}

func(s *Server) handleCommand(conn net.Conn, rawCmd []byte) {
	var(
	rawStr = string(rawCmd)
	parts = strings.Split(rawStr, " ")
		)
	if len(parts) == 0 {
		//respond
		fmt.Println("invalid command")
		return
	}

	cmd := Command(parts[0])
	if cmd == CMDSet{
		if len(parts) != 4 {
			//respond
			fmt.Println("invalid SET command")
			return
		}
		ttl, err := strconv.Atoi(parts[3])
		if err != nil {
			
		}
		msg ;= MSGSet{
			Key: []byte(parts[1]),
			Value: []byte(parts[2]),
			TTl: time.Duration(ttl),
		}
		if err := s.handleSetCmd(conn, msg); err != nil {
			//respond
			return
		}
	}
}

func(s *Server) handleSetCmd(conn net.Conn, msg ) error {
	fmt.Println("handling th set command: ", msg)
	return nil
}
