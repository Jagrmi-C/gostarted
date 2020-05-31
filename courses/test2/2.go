package main

import (
	"fmt"
	"time"
)

const n = 3

func main() {
	a := make(chan int, n)
	b := make(chan int, n)
	c := make(chan int)
	go sender(a, 0, 3)
	go sender(b, 3, 6)
	go receiver(a, b, c)
	time.Sleep(time.Millisecond)
	c <- 0
	time.Sleep(time.Second)
}

func sender(c chan int, x, y int) {
	for i := x; i < y; i++ {
		c <- i
	}
}

func receiver(a, b, c chan int) {
	ok := true
	for ok {
		select {
			case n := <-a:
			fmt.Println(n)
		case n := <-b:
			fmt.Println(n)
		case <-c:
			ok = false
		default:
			time.Sleep(time.Millisecond)
		}
	}
	fmt.Println("RECEIVER DONE")
}