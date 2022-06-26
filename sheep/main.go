package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := countSheepVolodia(0)
	fmt.Println(res)
}

func countSheep(num int) string {
	var listSheep = ""

	if num == 0 {
		return (listSheep)

	} else if num >= 1 {
		for i := 1; i <= num; i++ {
			listSheep += strconv.Itoa(i) + " sheep..."
		}
	}
	return (listSheep)
}

func countSheepVolodia(num int) string {
	var listSheep = ""

	for i := 1; i <= num; i++ {
		listSheep += strconv.Itoa(i) + " sheep..."
	}

	return listSheep
}
