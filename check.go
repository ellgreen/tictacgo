package main

const (
	Incomplete = iota
	Draw
	XWinner
	OWinner
)

var HumanResult = map[Result]string{
	Incomplete: "Incomplete",
	Draw:       "Draw",
	XWinner:    "Winner: X",
	OWinner:    "Winner: O",
}

type Result int

func (b *Board) Check() Result {
	for i := 0; i < 3; i++ {
		// Check rows
		row := b.State[i*3 : (i*3)+3]

		if string(row) == "XXX" {
			return XWinner
		}

		if string(row) == "OOO" {
			return OWinner
		}

		// Check columns
		col := []rune{b.State[i], b.State[i+3], b.State[i+6]}

		if string(col) == "XXX" {
			return XWinner
		}

		if string(col) == "OOO" {
			return OWinner
		}
	}

	// Check diagonal
	diag1 := []rune{b.State[0], b.State[4], b.State[8]}
	diag2 := []rune{b.State[2], b.State[4], b.State[6]}

	if string(diag1) == "XXX" || string(diag2) == "XXX" {
		return XWinner
	}

	if string(diag1) == "OOO" || string(diag2) == "OOO" {
		return OWinner
	}

	// Check if there are any more available moves
	if len(b.AvailablesMoves()) == 0 {
		return Draw
	}

	return Incomplete
}
