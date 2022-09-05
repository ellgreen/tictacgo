package main

import (
	"fmt"
)

func main() {
	var board Board

	board.State = [9]rune{'X', 'O', 'X', 'O', 0, 'O', 'X', 'X', 'X'}
	// board.State = [9]rune{0, 'O', 0, 'O', 'X', 'O', 'X', 'O', 'X'}
	// board.Play(4, 'X')

	// for i := 0; i < 9; i++ {
	// 	turn := 'X'

	// 	if i%2 == 0 {
	// 		turn = 'O'
	// 	}

	// 	board.Play(board.BestMove(turn), turn)
	// }

	board.Render()

	if winner, ok := board.Winner(); ok {
		fmt.Printf("The winner is %s\n", string(winner))
		return
	}

	fmt.Println("There is no winner")
}

type Board struct {
	State [9]rune
}

func (b *Board) Render() {
	fmt.Println("+-------+")

	for row := 0; row < 3; row++ {
		fmt.Printf("| ")

		for col := 0; col < 3; col++ {
			letter := b.State[(3*row)+col]

			if letter == 0 {
				fmt.Print("  ")
				continue
			}

			fmt.Printf("%s ", string(letter))
		}

		fmt.Printf("| \n")
	}

	fmt.Println("+-------+")
}

func (b *Board) AvailablesMoves() []int {
	var moves []int

	for i := 0; i < len(b.State); i++ {
		if b.State[i] == 0 {
			moves = append(moves, i)
		}
	}

	return moves
}

func (b *Board) Play(position int, letter rune) {
	b.State[position] = letter
}

func (b *Board) Winner() (rune, bool) {
	for i := 0; i < 3; i++ {
		// Check rows
		row := b.State[i*3 : (i*3)+3]

		if string(row) == "XXX" {
			return 'X', true
		}

		if string(row) == "OOO" {
			return 'O', true
		}

		// Check columns
		col := []rune{b.State[i], b.State[i+3], b.State[i+6]}

		if string(col) == "XXX" {
			return 'X', true
		}

		if string(col) == "OOO" {
			return 'O', true
		}
	}

	// TODO: Check diag

	return 0, false
}

func (b *Board) BestMove(turn rune) int {
	moves := b.AvailablesMoves()
	var bestMove int

	for move := range moves {
		bestMove = move
	}

	return bestMove
}
