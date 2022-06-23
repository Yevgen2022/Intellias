package main

import "fmt"

func remove(Slice []int32, i int) []int32 { //Функція яка видаляє дублікати
	copy(Slice[i:], Slice[i+1:]) //із збереженням порядку
	return Slice[:len(Slice)-1]
}

func main() {
	var mySlice []int32
	mySlice = make([]int32, 6) //Вхідний слайс.В якому через один елемент
	mySlice[0] = 3             //лежать дублікати для видалення
	mySlice[1] = 4
	mySlice[2] = 4
	mySlice[3] = 3
	mySlice[4] = 6
	mySlice[5] = 3

	for j := 1; j < len(mySlice); j++ {

		mySlice = remove(mySlice, j)
	}
	fmt.Println("Це перший слайс без дублікатів", mySlice)

	myArr := []int32{4, 1, 4, -4, 6, 3, 8, 8} //Це слайс з якого треба видалити
	//та зберегти дублікати
	var myResult []int32
	for j := 0; j < len(myArr); j++ {
		for i := j + 1; i < len(myArr); i++ {
			if myArr[j] == myArr[i] {
				myResult = append(myResult, myArr[i])
				myArr = remove(myArr, j)
			}
		}
	}

	fmt.Println("Видалені дублікати другого слайсу", myResult)
	fmt.Println("Другий слайс без дублікатів", myArr)

}
