package main

import "math"

func Minimax(board Board, depth int, maximiser rune, turn rune) float32 {
	result := board.Check()

	if depth == 0 || result != Incomplete {
		return Evaluate(&board, maximiser)
	}

	if maximiser == turn {
		// Maximising

		var maxScore float32 = -math.MaxFloat32

		for _, move := range board.AvailablesMoves() {
			newBoard := board
			newBoard.Play(move, turn)

			score := Minimax(newBoard, depth-1, maximiser, NextTurn(turn))

			if score > maxScore {
				maxScore = score
			}
		}

		return maxScore
	}

	// Minimising
	var minScore float32 = math.MaxFloat32

	for _, move := range board.AvailablesMoves() {
		newBoard := board
		newBoard.Play(move, turn)

		score := Minimax(newBoard, depth-1, maximiser, NextTurn(turn))

		if score < minScore {
			minScore = score
		}
	}

	return minScore
}

func Evaluate(board *Board, maximiser rune) float32 {
	result := board.Check()

	if result == Incomplete || result == Draw {
		return 0
	}

	if result == XWinner && maximiser == 'X' {
		return 1
	}

	if result == OWinner && maximiser == 'O' {
		return 1
	}

	return -1
}
