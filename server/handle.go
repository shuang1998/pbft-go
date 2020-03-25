package server

import (
	"encoding/json"
	"github.com/pipapa/pbft/message"
	"log"
	"net/http"
)

func (s *HttpServer) HttpRequest(w http.ResponseWriter, r *http.Request) {
	var msg message.Request
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		log.Printf("[Http Error] %s", err)
		return
	}
	s.requestRecv <- &msg
}
