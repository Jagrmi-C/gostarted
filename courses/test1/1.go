package main

import (
    "fmt"
    "math/rand"
    "time"
)

func Merge2Channels(
    f func(int) int,
    in1 <-chan int,
    in2 <-chan int,
    out chan<- int,
    n int,
) {
    fmt.Println("Start")
    ok := true
	for ok {
		select {
        case n := <-in1:
            res := updattter(n) + updattter(<-in2)
            out <- res
        case n := <-in2:
            res := updattter(n) + updattter(<-in1)
            out <- res
        default:
			time.Sleep(time.Millisecond)
        }
    }
    
}

func sender(c chan int, n int) {
	for i := 0; i < n; i++ {
        num := rand.Intn(10) + i
        fmt.Println("Add num", num)
        c <- num
    }
    close(c)
}

func updattter(in int) (v int) {
    for {
        return in + 1
    }
}

func printer(out chan int, n int)  {
    // fmt.Println("start printer")
    for i := 0; i < n; i++ {
        num := <-out
        fmt.Println("RES", num)
    }
}

func main() {
    const n = 3
    var in1 chan int = make(chan int, n)
    var in2 chan int = make(chan int, n)

    var out chan int = make(chan int, n)

    go sender(in1, n)
	go sender(in2, n)

    go Merge2Channels(updattter, in1, in2, out, n)
    
    go printer(out, n)

    // res := <-out
    // return res
    var input string
    fmt.Scanln(&input)
}