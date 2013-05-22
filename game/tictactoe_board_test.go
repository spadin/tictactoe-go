package game

import (
	"fmt"
	"github.com/remogatto/prettytest"
	"testing"
)

type testTictactoeBoardSuite struct {
	prettytest.Suite
}

func TestTictactoeBoardRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testTictactoeBoardSuite),
	)
}

func tictactoeBoardSuiteSetup() *TictactoeBoard {
	return NewTictactoeBoard()
}

func (t *testTictactoeBoardSuite) TestEmptyIndexForNewTictactoeBoard() {
	board := tictactoeBoardSuiteSetup()
	t.True(board.IsEmpty(0), "new tictactoeBoard should have empty position")
}

func (t *testTictactoeBoardSuite) TestNonEmptyTictactoeBoardIndex() {
	board := tictactoeBoardSuiteSetup()
	board[0] = X
	t.False(board.IsEmpty(0), "marked position not empty")
}

func (t *testTictactoeBoardSuite) TestMarkingTheTictactoeBoard() {
	board := tictactoeBoardSuiteSetup()
	board.SetMark(X, 0)
	t.False(board.IsEmpty(0), "marks tictactoeBoard a index 0")
}

func (t *testTictactoeBoardSuite) TestTictactoeBoardIsFull() {
	board := tictactoeBoardSuiteSetup()
	playTiedGame(board)
	t.True(board.IsFull(), "determines the tictactoeBoard is full")
}

func (t *testTictactoeBoardSuite) TestTictactoeBoardIsNotFull() {
	board := tictactoeBoardSuiteSetup()
	t.False(board.IsFull(), "determines tictactoeBoard is not full")
}

func (t *testTictactoeBoardSuite) TestTictactoeBoardOpenPositions() {
	board := tictactoeBoardSuiteSetup()
	makeMoves(board, 0, 1, 2)

	expected, actual := []int{3, 4, 5, 6, 7, 8}, board.OpenPositions()

	strExpected := fmt.Sprintf("%v", expected)
	strActual := fmt.Sprintf("%v", actual)

	t.Equal(strExpected, strActual, "returns open positions")
}
