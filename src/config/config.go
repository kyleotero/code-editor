package config

import (
  "time"
)

type Action int

const (
  Write = iota
  Delete
)

type Update struct {
  character rune
  index int
}

type UpdateRequest struct {
  action Action
  updates []Update 
}

type ClientObject struct {
  TimeJoined time.Time `json:"timeJoined,omitempty"`
  IPAdress string `json:"ipAddress,omitempty"`
  Token string `json:"token,omitempty"`
  ClientWebSocket *websocket.Conn
  mu sync.Mutex
}
