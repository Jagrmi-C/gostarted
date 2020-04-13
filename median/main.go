package main

import (
	"fmt"
	"sort"
	"log"
)

func median(intsArr []int) float64 {
	var answer float64
	val := len(intsArr) / 2
	log.Println("Half value: ", val)
	if len(intsArr) % 2 == 1 {
		log.Println("%2 == 1")
		answer = float64(intsArr[val])
		log.Println("Answer:", answer)
	} else {
		log.Println("%2 == 0")
		answer = (float64(intsArr[val -1]) + float64(intsArr[val])) / 2
		log.Println("Answer:", answer)
	}
	return answer
}

func main()  {
	intsArr := []int{7, 2, 4, 8, 1, 2, 9}
	// intsArr := []int{7, 2, 4, 8, 1, 9}
	sort.Ints(intsArr)
	log.Println("Sorted list:", intsArr)
	medianVal := median(intsArr)
	fmt.Println(medianVal)
}