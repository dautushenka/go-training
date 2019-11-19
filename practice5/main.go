package practice5

func NewMatch(player1Skill, player2Skill int) *Match {
	return &Match{
		Player1: &Player{
			"Player1",
			player1Skill,
		},
		Player2: &Player{
			"Player2",
			player2Skill,
		},
	}
}
