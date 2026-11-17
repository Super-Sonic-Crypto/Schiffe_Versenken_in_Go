// Schiffe versenken Main
// ver 0.1.1

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const BoardSize int = 10

	var PlayingField = [10][10]string{}
	RenderBoard(BoardSize, PlayingField)
	fmt.Printf("\n")

	Input := bufio.NewScanner(os.Stdin)

	for i := 0; i < 3; i++ {
		x, y, err := ParseCoordinate(Input, i+1)

		if err != nil {
			fmt.Print(err)
		}

		fmt.Printf("Eingabe %d: X: %d,  Y: %d \n\n", i+1, x, y)
	}
}
