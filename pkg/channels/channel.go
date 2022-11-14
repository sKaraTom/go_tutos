package main

import (
	"fmt"
	"time"
)

func run(ch chan string, name string) {
	time.Sleep(time.Second * 2) // attente 2 secondes
	fmt.Println("fonction run() :", name)
	ch <- name
}

func main() {

	now := time.Now()

	ch := make(chan string)

	go run(ch, "channel 1")
	go run(ch, "channel 2")

	// La première goroutine qui aura écrit sur un channel sera la première à être lu dans notre programme.
	fmt.Println("fonction main() :", <-ch)
	fmt.Println("fonction main() :", <-ch)

	fmt.Println(time.Now().Sub(now))
}
