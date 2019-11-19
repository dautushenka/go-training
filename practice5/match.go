package practice5

type Match struct {
	Player1 *Player
	Player2 *Player
}

func (s *Match) Start() *Player {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go s.Player1.Play(ch1)
	go s.Player2.Play(ch2)
	for true {
		if _, ok := <-ch1; !ok {
			return s.Player2
		}
		if _, ok := <-ch2; !ok {
			return s.Player1
		}
	}
	return &Player{}
}
