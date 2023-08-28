package netsrv

import (
	"bufio"
	"fmt"
	"go-course/homework-11/pkg/search"
	"log"
	"net"
)

type Netsrv struct {
	search *search.Search
}

func New(search *search.Search) *Netsrv {
	return &Netsrv{search: search}
}

func handler(conn net.Conn, search *search.Search) {
	defer conn.Close()
	defer fmt.Println("Connection Closed")

	r := bufio.NewReader(conn)

	for {
		query, _, err := r.ReadLine()
		if err != nil {
			return
		}

		result := search.Results(string(query))

		for _, s := range result {
			_, err = conn.Write([]byte(s))
			if err != nil {
				return
			}
		}
	}
}

func (n *Netsrv) Start() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Сервер слушает на порту: 8000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn, n.search)
	}
}
