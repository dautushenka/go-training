package main

import (
	"fmt"
	"go-training/practice"
)

func main() {
	// Practice 1
	practice.Fillarray(10, "0");
	a := practice.Filldiagonal(5, "1", "0");
	fmt.Println("[Practice 1] Print 2d array with filled dialonals")
	practice.Print2darray(a);

	fmt.Println("")
	// Practice 1.2
	fmt.Println("[Practice 1.2] Print dice statistic")
	practice.Dicestatistic(10000);
}