package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var aMutex sync.Mutex

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			aMutex.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			aMutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
