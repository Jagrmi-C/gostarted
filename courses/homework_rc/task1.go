package main

import (
	"fmt"
	"sync"
	// "flag"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	// wg.Add(gs)

	for i := 0; i < gs; i++ {
		wg.Add(1)
		go func() {
			// defer wg.Done()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
		wg.Wait()
	}
	// wg.Wait()
	fmt.Printf("%#v\n", wg)
	fmt.Println("end value:", incrementer)
}
