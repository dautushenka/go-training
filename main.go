package main

import (
	"fmt"
	"go-training/practice1"
	"go-training/practice2"
)

func runPractice1() {
	// Practice 1
	practice1.Fillarray(10, "0");
	a := practice1.Filldiagonal(5, "1", "0");
	fmt.Println("[Practice 1] Print 2d array with filled dialonals")
	practice1.Print2darray(a);

	fmt.Println("")
	// Practice 1.2
	fmt.Println("[Practice 1.2] Print dice statistic")
	practice1.Dicestatistic(10000);
}

func runPractice2() {
	collection := practice2.Collection{};
	for i := 0; i < 10; i++ {
		collection.Add(practice2.Element(i));
	}
	collection.Get(5);
	if _, error := collection.Get(5); error != nil {
		panic("There is no element with required index")
	}
	collection.Remove(5);
	fmt.Printf("First: %v\n", collection.First());
	fmt.Printf("Last: %v\n", collection.Last());
	fmt.Printf("Curent: %v\n", collection.Value());
	fmt.Printf("Next: %v\n", collection.Next());
	fmt.Printf("Next: %v\n", collection.Next());
	fmt.Printf("Prev: %v\n", collection.Prev());
	fmt.Printf("Length: %v\n", collection.Length());
	collection.Print();
}

func main() {
	runPractice2();
}