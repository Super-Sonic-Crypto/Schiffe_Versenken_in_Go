// Schiffe Versenken Main
// ver 0.0.1

package main

import "fmt"

func main() {
	const BoardSize int = 10

	var PlayingField = [10][10]string{}
	RenderBoard(BoardSize, PlayingField)
}

func RenderBoard(BoardSize int, PlayingField [10][10]string) {

	for i := range BoardSize {
		for j := range 10 {
			PlayingField[i][j] = "-"
		}
	}

	fmt.Print("  ")
	for cor := 'A'; cor < 'A'+10; cor++ {
		fmt.Printf("  %c", cor)
	}

	fmt.Println()

	for i := range 10 {
		if i != 9 { // Abfrage wegen der Formatierung sonst ist die 10 komisch aus der Reihe
			fmt.Printf(" %d  ", i+1)
		} else {
			fmt.Printf("%d  ", i+1)
		}

		for j := range 10 {
			fmt.Print(PlayingField[i][j], "  ")
		}
		fmt.Println()
	}
}
