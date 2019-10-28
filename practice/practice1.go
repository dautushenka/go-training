package practice

import (
	"fmt"
)

func Fillarray(size int, char string) (array [][]string) {
	row := make([]string, size);
	for index := range(row) {
		row[index] = char;
	}
	array = make([][]string, size);
	for x := range(array) {
		array[x] = row;
	}
	
	return;
}

func Filldiagonal(size int, char string, defaultChar string) (array [][]string) {
	array = make([][]string, size);
	for x := range(array) {
		row := make([]string, size);
		for y := range(row) {
			if (y == x || (y + x) == (size - 1)) {
				row[y] = char;
			} else {
				row[y] = defaultChar;
			}
		}
		array[x] = row;
	}
	return;
}

func Print2darray(array [][]string) {
	for _, v := range(array) {
		fmt.Println(v);
		// for y := range(array[x]) {
			
		// }
	}
}