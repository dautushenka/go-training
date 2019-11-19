package practice5

import (
	"math/rand"
	"time"
)

type Player struct {
	Name  string
	Skill int
}

func (s *Player) Play(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(100)
	for s.Skill >= random {
		ch <- 1
		random = rand.Intn(100)
	}
	close(ch)
}
