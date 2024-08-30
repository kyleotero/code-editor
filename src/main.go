package main

import (
  "net/http"
)

func main() {
  http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
    handler(w, r)
  })
}
