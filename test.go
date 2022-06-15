package main

import "fmt"

func main() {
	var oneApple float64 = 5.99
	var onePear int = 7
	var allMoney int = 23

	fmt.Println("Щоб купити 9 яблук та 8 груш, треба витратити", oneApple*float64(onePear), "грн.") //  41.93 грн.
	fmt.Println("На наші гроші ми можемо купити", allMoney/onePear, "груші")                        //  3 груші
	fmt.Println("На наші гроші ми можемо купити", int(float64(allMoney)/oneApple), "яблука")        //  3 яблука
	fmt.Println("Чи ми можемо купити 2 яблука та 2 груші", allMoney >= int(2*oneApple)+2*onePear)   //  false
}
