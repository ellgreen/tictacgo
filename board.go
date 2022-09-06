package main

import (
	"math"
)

type Board struct {
	State [9]rune
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

func (b *Board) BestMove(turn rune) int {
	var bestMove int
	var bestScore float32 = -math.MaxFloat32

	for _, move := range b.AvailablesMoves() {
		board := *b
		board.Play(move, turn)

		score := Minimax(board, 7, turn, NextTurn(turn))

		if score > bestScore {
			bestMove, bestScore = move, score
		}
	}

	return bestMove
}

func NextTurn(turn rune) rune {
	if turn == 'X' {
		return 'O'
	}

	return 'X'
}
