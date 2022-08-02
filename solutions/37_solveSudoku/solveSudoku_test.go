package solvesudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveSudoku(t *testing.T) {
	t.Skip("The method is not implemented yet")
	sodoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(sodoku)
	expectedSolution := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}
	assert.ObjectsAreEqualValues(sodoku, expectedSolution)
}

func Test_providesCorrectBlock(t *testing.T) {
	board := sudoku([][]cell{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '3', '.', '8', '.', '.', '7', '9'},
	})
	assert.ObjectsAreEqualValues(string([]byte{'5', '3', '.', '6', '.', '.', '.', '9', '8'}), string(board.block(0)))
	assert.ObjectsAreEqualValues(string([]byte{'.', '7', '.', '1', '9', '5', '.', '.', '.'}), string(board.block(1)))
	assert.ObjectsAreEqualValues(string([]byte{'.', '.', '.', '.', '.', '1', '.', '6', '.'}), string(board.block(2)))

	assert.ObjectsAreEqualValues(string([]byte{'8', '.', '.', '4', '.', '.', '7', '.', '.'}), string(board.block(3)))
	assert.ObjectsAreEqualValues(string([]byte{'.', '6', '.', '8', '.', '3', '.', '2', '.'}), string(board.block(4)))
	assert.ObjectsAreEqualValues(string([]byte{'.', '.', '3', '.', '.', '1', '.', '.', '6'}), string(board.block(5)))

	assert.ObjectsAreEqualValues(string([]byte{'.', '6', '.', '.', '.', '.', '.', '.', '.'}), string(board.block(6)))
	assert.ObjectsAreEqualValues(string([]byte{'.', '.', '.', '4', '1', '9', '.', '8', '.'}), string(board.block(7)))
	assert.ObjectsAreEqualValues(string([]byte{'.', '.', '3', '.', '8', '.', '.', '7', '9'}), string(board.block(8)))
}
