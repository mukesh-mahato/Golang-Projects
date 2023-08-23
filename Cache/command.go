package main

import "time"

type Commnd string

const (
  CMDSet Command = "SET"
  CMDGet Command = "SET"
)

type MSGSet struct {
  Key []byte
  Set []byte
  TTL time.Duration
}

type Message struct {
  cmd Command
  key []byte
  value []byte
  TTL time.Duration
}

func parseCommand(raw []byte) (*Message, error) {
  
}
