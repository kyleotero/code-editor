package main

import (
  "log"
  "github.com/gorilla/websocket"
  "github.com/kyleotero/code-editor/src/config"
)

type Server struct {
  clients map[]
}

var upgrader = websocket.Upgrader {
  ReadBufferSize: 1024
  WriteBufferSize: 1024,
}

func handler (w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade

  if (err != nil) {
    log.Println(err)
    return
  }
}


