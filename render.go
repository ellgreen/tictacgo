package main

import "fmt"

func (b *Board) Render() {
	fmt.Println("+-------+")
	defer fmt.Println("+-------+")

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
}
