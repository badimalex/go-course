package main

import (
	"go-course/homework-14/pkg/messages"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type Server struct {
	messages []messages.Message
	nextID   int
	mu       sync.Mutex
}

func (s *Server) Send(content string, resp *bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg := messages.Message{
		ID:      s.nextID,
		Time:    time.Now(),
		Content: content,
	}

	s.messages = append(s.messages, msg)
	s.nextID++
	*resp = true

	return nil
}

func (s *Server) Messages(content string, resp *[]messages.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	*resp = s.messages
	return nil
}

func main() {
	srv := new(Server)
	err := rpc.Register(srv)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeConn(conn)
	}
}
