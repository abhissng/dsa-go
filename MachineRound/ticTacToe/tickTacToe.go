package tickTacToe

import (
	"fmt"
	"strings"
)

const (
	empty     = " "
	playerOne = "X"
	playerTwo = "O"
)

type Player struct {
	name  string
	piece string
}

type board struct {
	cells [3][3]string
}

func (b *board) isEmpty(row, col int) bool {
	return b.cells[row][col] == empty
}

func (b *board) print() {
	for i := 0; i < 3; i++ {
		fmt.Println(strings.Join(b.cells[i][:], " | "))
	}
}

func (b *board) placeMove(row, col int, piece string) {
	b.cells[row][col] = piece
}

func (b *board) isWinner(piece string) bool {
	// check rows
	for i := 0; i < 3; i++ {
		if b.cells[i][0] == piece && b.cells[i][1] == piece && b.cells[i][2] == piece {
			return true
		}
	}

	//check columns
	for j := 0; j < 3; j++ {
		if b.cells[0][j] == piece && b.cells[1][j] == piece && b.cells[2][j] == piece {
			return true
		}
	}

	// check diagonals
	if b.cells[0][0] == piece && b.cells[1][1] == piece && b.cells[2][2] == piece {
		return true
	}
	if b.cells[0][2] == piece && b.cells[1][1] == piece && b.cells[2][0] == piece {
		return true
	}
	return false

}

func (b *board) isFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.isEmpty(i, j) {
				return false
			}
		}
	}
	return true
}

func Init() {
	var player1, player2 Player
	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Print("Enter Player 1 Name: ")
	fmt.Scanln(&player1.name)
	player1.piece = playerOne

	fmt.Print("Enter Player 2 Name: ")
	fmt.Scanln(&player2.name)
	player2.piece = playerTwo

	currentPlayer := &player1
	gameBoard := board{}

	gameBoard.print()

	for {
		var row, col int
		var move string

		row--
		col--

	}

}
