package game

type tictactoeRules struct {
	board Board
}

func NewTictactoeRules(board Board) (rules Rules) {
	rules = &tictactoeRules{board: board}
	return
}

func winningCombinations() [][3]int {
	return [][3]int{
		[3]int{0, 1, 2}, [3]int{4, 5, 6}, [3]int{7, 8, 9},
		[3]int{0, 3, 6}, [3]int{1, 4, 7}, [3]int{2, 5, 8},
		[3]int{0, 4, 8}, [3]int{2, 4, 6},
	}
}

func (rules *tictactoeRules) IsGameover() bool {
	return rules.HasWinner() || rules.board.IsFull()
}

func (rules *tictactoeRules) HasWinner() bool {
	return !(rules.Winner() == nil)
}

func (rules *tictactoeRules) Winner() (winner interface{}) {
	winner = nil
	board := rules.board.(*TictactoeBoard)
	for _, combination := range winningCombinations() {
		if board[combination[0]] == E {
			continue
		}
		if board[combination[0]] == board[combination[1]] &&
			board[combination[0]] == board[combination[2]] {
			winner = board[combination[0]]
			break
		}
	}
	return winner

}