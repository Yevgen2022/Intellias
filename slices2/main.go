package main

import (
	"fmt"
)

func main() {
	task1()
	task2()
}

func task1() {
	mySlice := []int{3, 4, 4, 3, 6, 3}
	var result []int
	duplicates := make(map[int]bool)

	for _, v := range mySlice {
		if _, ok := duplicates[v]; !ok {
			result = append(result, v)
		} else {
			continue
		}

		duplicates[v] = true
	}

	fmt.Println("Це перший слайс без дублікатів: ", result)
}

func task2() {

}
