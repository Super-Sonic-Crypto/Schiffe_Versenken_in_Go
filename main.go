// Schiffe Versenken Main
// ver 0.0.1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const BoardSize int = 10

	var PlayingField = [10][10]string{}
	RenderBoard(BoardSize, PlayingField)

	Input := bufio.NewScanner(os.Stdin)

	for i := 0; i < 3; i++ {
		x, y := ParseCoordinate(Input, i+1)
		fmt.Printf("Eingabe %d: X: %d,  Y: %d \n\n", i+1, x, y)
	}
}

func ParseCoordinate(Input *bufio.Scanner, RunThrough int) (x, y int) {
	for {
		fmt.Printf("%d. Koordinate beginnend mit dem Buchstabe (Spalte) eingeben (z.B. A1): ", RunThrough)
		Input.Scan()
		CoordinateInput := Input.Text()

		if len(CoordinateInput) > 3 || len(CoordinateInput) < 2 {
			fmt.Println("Fehler, bitte mindestens 2 oder höchstens 3 Zeichen eingeben!")
			fmt.Printf("Nochmal Eingeben!\n\n")
			continue
		}

		x = int(CoordinateInput[0] - 'A')
		y, _ = strconv.Atoi(CoordinateInput[1:])

		return x, y
	}
}
