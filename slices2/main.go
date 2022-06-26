package main

import "fmt"

func main() {
	result := task1([]int{3, 4, 4, 3, 6, 3})
	fmt.Println("Це перший слайс без дублікатів: ", result)

	task2()
}

func task1(mySlice []int) []int {
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

	return result
}

func task2() {

}
