package main

import (
	"fmt"
	"log"
	"sort"
)

func median(intsArr []int) float64 {
	sort.Ints(intsArr)
	log.Println("Sorted list:", intsArr)
	val := len(intsArr) / 2
	if len(intsArr)%2 == 1 {
		log.Println("%2 == 1")
		return float64(intsArr[val])
	}

	log.Println("%2 == 0")
	return (float64(intsArr[val-1]) + float64(intsArr[val])) / 2
}

func main() {
	intsArrOdd := []int{7, 2, 4, 8, 1, 2, 9}
	intsArrEven := []int{7, 2, 4, 8, 1, 9}
	medianVal1 := median(intsArrOdd)
	medianVal2 := median(intsArrEven)
	fmt.Println(medianVal1, medianVal2)
}
