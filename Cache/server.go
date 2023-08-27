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
	msg err := parseCommand(rawCmd)
	if err != nil {
		return
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

func parseCommand(raw []byte) (*Message, error) {
	var(
	rawStr = string(raw)
	parts = strings.Split(rawStr, " ")
		)
	if len(parts) < 2 {
		return nil, errors.New("invalid protocal format")
	}
	msg:= &Message {
		Cmd: Command(parts[0]),
		Key: []byte(parts[1]),
	}
	if msg.Cmd == CMDSet {
		if len(parts) < 4 {
		return nil, errors.New("invalid SET command")
	}
		msg.Value = []byte(parts[2])
		
		ttl, err := strconv.Atoi(parts[3])
		if err != nil {
			return errors.New("invalid SET TTL")
		}
		msg.TTL = time.Duration(ttl)
	}
	return msg, nil
}
