package practice1

import (
	"fmt"
	"math/rand"
	"time"
)

func Shootdice() int {
	return rand.Intn(6) + rand.Intn(6) + 2;
}

func Dicestatistic(iterations int) {
	var a [13]int;
	var v int;
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < iterations; i++ {
		v = Shootdice();
		a[v]++;
	}
	for number, count := range(a[2:]) {
		fmt.Printf("%d - %.1f%%\n", number + 2, (float32(count) / float32(iterations)) * 100);
	}
}