package main

import (
	"fmt"
	"go-training/practice1"
	"go-training/practice2"
	"go-training/practice2_1"
	"go-training/practice3"
	"go-training/practice4"
)

func runPractice1() {
	// Practice 1
	practice1.Fillarray(10, "0")
	a := practice1.Filldiagonal(5, "1", "0")
	fmt.Println("[Practice 1] Print 2d array with filled dialonals")
	practice1.Print2darray(a)

	fmt.Println("")
	// Practice 1.2
	fmt.Println("[Practice 1.2] Print dice statistic")
	practice1.Dicestatistic(10000)
}

func runPractice2() {
	collection := practice2.Collection{}
	for i := 0; i < 10; i++ {
		collection.Add(practice2.Element(i))
	}
	collection.Get(5)
	if _, error := collection.Get(5); error != nil {
		panic("There is no element with required index")
	}
	collection.Remove(5)
	fmt.Printf("First: %v\n", collection.First())
	fmt.Printf("Last: %v\n", collection.Last())
	fmt.Printf("Curent: %v\n", collection.Value())
	fmt.Printf("Next: %v\n", collection.Next())
	fmt.Printf("Next: %v\n", collection.Next())
	fmt.Printf("Prev: %v\n", collection.Prev())
	fmt.Printf("Length: %v\n", collection.Length())
	collection.Print()
}

func runPractice2_1() {
	collection := practice2_1.Collection{}
	for i := 0; i < 10; i++ {
		collection.Add(i)
	}
	collection.Print()
	collection.Get(5)
	if _, error := collection.Get(5); error != nil {
		panic("There is no element with required index")
	}
	collection.Remove(5)
	fmt.Println(collection.First())
	fmt.Println(collection.Last())
	fmt.Println(collection.Value())
	fmt.Println(collection.Next())
	fmt.Println(collection.Next())
	fmt.Println(collection.Prev())
	fmt.Println(collection.Length())
	collection.Print()
}

func runPractice3() {
	//practice3.RunBot()
	//practice3.RunBot2()
	practice3.RunBot3()
}

func runPractice4() {
	service := practice4.Meteorologist{}
	//weather := service.DailyForecast("London", 7)
	weather := service.WeatherForecast("Mahilyow")
	fmt.Println(weather)
}

func main() {
	runPractice4()
}
