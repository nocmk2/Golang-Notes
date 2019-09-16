package main

import (
	"fmt"
	"time"
)

func sender(c chan<- string) {
	for {
		c <- "ping"
	}
}

func sender2(c chan<- string) {
	for {
		c <- "haha"
	}
}

func printer(c <-chan string) {
	for {
		a := <-c
		fmt.Println(a)
		time.Sleep(time.Second)
	}
}

func block() {
	var input string
	fmt.Scanln(&input)
}

func main() {
	var c = make(chan string)
	go sender(c)
	go sender2(c)
	go printer(c)
	block()
}
