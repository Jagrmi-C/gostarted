package main

import (
	"fmt"
	"sort"
)

// averageArray that calculate average value
func averageArray(numbers [6]int) float64 {
	var sum int = 0
	for _, value := range numbers {
		sum += value
	}
	return float64(sum) / float64(len(numbers))
}

// maxStringValue that returns the longest word
func maxStringValue(s []string) (maxVal string) {
	for _, value := range s {
		if len(value) > len(maxVal) {
			maxVal = value
		}
	}
	return
}

// reverseSlice that returns the reversed copy
func reverseSlice(s []int64) []int64 {
	reversed := make([]int64, 0, len(s))
	for i := 0; i < len(s); i++ {
		reversed = append(reversed, s[len(s)-1-i])
	} // another way to write a loop
	return reversed
}

// printSorted that is sorted values of map by key
func printSorted(m map[int]string) {
	keys := make([]int, 0, len(m))
	s := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		s = append(s, m[k])
	}
	fmt.Println("From map:", m, "->", s)
}

func main() {
	// Task 1
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(
		"Average value for array", numbers, "is", averageArray(numbers),
	)
	// Task 2
	words := []string{"one", "two"}
	ans1 := maxStringValue(words)
	words = append(words, "three")
	ans2 := maxStringValue(words)
	words = []string{"four", "ten", "eleven", "five", "seven", "twenty"}
	ans3 := maxStringValue(words)
	ans4 := maxStringValue(make([]string, 0))

	fmt.Println(ans1, ans2, ans3, ans4)
	// Task 3
	s1 := []int64{1001, 232, 354, 8784, 45, 6}
	ans5 := reverseSlice(s1)
	fmt.Println("Origin:", s1, "-> Reversed:", ans5)
	// Task 4
	myMap1 := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(myMap1)
	myMap2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(myMap2)
	myMap3 := map[int]string{11: "ar", 100: "vr", 50: "vr", 5: "kn", 78: "ar"}
	printSorted(myMap3)
}
