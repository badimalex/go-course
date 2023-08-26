package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	table := make(chan string)

	player1 := Player{name: "Player A"}
	player2 := Player{name: "Player B"}

	wg.Add(2)

	go player(table, &player1)
	go player(table, &player2)

	table <- "begin"

	wg.Wait()

	fmt.Println("Game Over!")
	fmt.Printf("%d:%d\n", player1.score, player2.score)
}

type Player struct {
	name  string
	score int
}

func player(ch chan string, player *Player) {
	defer wg.Done()

	for val := range ch {
		if val == "end" {
			break
		}

		time.Sleep(time.Second)

		fmt.Printf("%s: %s\n", player.name, val)

		if val == "stop" {
			player.score++
			if player.score == 3 {
				ch <- "end"
				break
			} else {
				ch <- "begin"
			}
		}

		if val == "begin" {
			ch <- "ping"
		}

		if val == "ping" {
			if rand.Intn(100) < 20 {
				ch <- "stop"
			} else {
				ch <- "pong"
			}
		}

		if val == "pong" {
			if rand.Intn(100) < 20 {
				ch <- "stop"
			} else {
				ch <- "ping"
			}
		}
	}
}
