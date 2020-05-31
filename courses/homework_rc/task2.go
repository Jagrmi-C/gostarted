package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer
	reset := make(chan bool)

	t = time.AfterFunc(
		randomDuration(),
		func() {
			fmt.Println(time.Now().Sub(start))
			reset <- true
		},
	)
	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
