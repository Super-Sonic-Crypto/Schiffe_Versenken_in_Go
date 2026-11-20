// Schiffe versenken Main
// ver 0.1.3

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	const BoardSize int = 10

	var PlayingField = [10][10]string{}
	var RadarField = [10][10]string{}

	Menu := 0

	for {
		RandX := rand.IntN(10)
		RandY := rand.IntN(10)

		fmt.Printf("[1] Spiel starten \n")
		fmt.Printf("[2] Beenden \n\n")
		fmt.Printf("Deine Eingabe: ")
		fmt.Scan(&Menu)
		fmt.Println()

		switch Menu {
		case 1:
			fmt.Printf("--- Spiel Start! ---\n\n")

		case 2:
			fmt.Printf("Spiel beendet. \n")
			fmt.Printf("Einen schönen Tag!")
			return
		default:
			fmt.Printf("Ungültige Eingabe. \n")
			fmt.Printf("Einen schönen Tag.")
			return
		}

		RenderBoard(BoardSize, PlayingField, RadarField)
		fmt.Printf("\n")

		Input := bufio.NewScanner(os.Stdin)

		for i := 0; i < 3; i++ {
			x, y, err := ParseCoordinate(Input, i+1)
			if err != nil {
				fmt.Print(err)
			}

			fmt.Printf("Eingabe %d: X: %d,  Y: %d \n\n", i+1, x, y)
			//if i != 3 {
			PlayingField[y][x] = "#"
			fmt.Printf("\n\n")
			//}

			/*
				if i == 3 {
					RadarField[y][x] = "X"
					RenderBoard(BoardSize, PlayingField, RadarField)
					fmt.Printf("\n\n")
				}
			*/
		}

		RadarField[RandY][RandX] = "X"
		RenderBoard(BoardSize, PlayingField, RadarField)

		fmt.Println()

		for i := range BoardSize {
			for j := range BoardSize {
				if PlayingField[i][j] == "#" && RadarField[i][j] == "X" {
					fmt.Printf("Treffer bei X: %d, Y: %d", RandX, RandY)
					fmt.Printf("Spiel beendet!")
					return
				} else {
					for i := range BoardSize {
						for j := range BoardSize {
							PlayingField[i][j] = "-"
							RadarField[i][j] = "-"
						}
					}
				}

			}
		}

		fmt.Printf("\n\n\n")
	}
}
