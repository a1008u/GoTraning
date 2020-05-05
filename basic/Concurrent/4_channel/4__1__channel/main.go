package ___1__channel

import (
	"log"
)

func greet(ch chan string) {
	defer close(ch)
	ch <- "Hello"
}

func main() {
	ch := make(chan string)
	go greet(ch)
	log.Println(<-ch)
}
