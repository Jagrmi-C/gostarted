package main

import (
	"fmt"
	"log"
	"sort"
)

func median(intsArr []int) float64 {
	var answer float64
	sort.Ints(intsArr)
	log.Println("Sorted list:", intsArr)
	val := len(intsArr) / 2
	log.Println("Half value: ", val)
	if len(intsArr)%2 == 1 {
		log.Println("%2 == 1")
		answer = float64(intsArr[val])
		log.Println("Answer:", answer)
		return answer
	}

	log.Println("%2 == 0")
	answer = (float64(intsArr[val-1]) + float64(intsArr[val])) / 2
	log.Println("Answer:", answer)

	return answer
}

func main() {
	intsArrOdd := []int{7, 2, 4, 8, 1, 2, 9}
	intsArrEven := []int{7, 2, 4, 8, 1, 9}
	medianVal1 := median(intsArrOdd)
	medianVal2 := median(intsArrEven)
	fmt.Println(medianVal1, medianVal2)
}
