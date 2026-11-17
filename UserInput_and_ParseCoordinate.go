// Koordinaten Eingabe, abfragen nach Fehlerzuständen
// Achtung keine Abfrage nach den Schiffen, ob das Setzen logisch ist

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ParseCoordinate(Input *bufio.Scanner, RunThrough int) (x, y int, err error) {
	for {
		fmt.Printf("%d. Koordinate beginnend mit dem Buchstabe (Spalte) eingeben (z.B. A1): ", RunThrough)
		Input.Scan()
		CoordinateInput := Input.Text()
		CoordinateInput = strings.ToUpper(CoordinateInput) //Eingabe wird zu großbuchstabe gezwungen

		if len(CoordinateInput) > 3 || len(CoordinateInput) < 2 {
			fmt.Printf("Fehler! bitte mindestens 2 oder höchstens 3 Zeichen eingeben! \n\n")
			continue
		}

		x = int(CoordinateInput[0] - 'A')
		y, err = strconv.Atoi(CoordinateInput[1:])

		y = y - 1

		// Switch case für alle Fehler Zustände beim Eintippen
		switch {
		case err != nil:
			fmt.Printf("Fehler! Erst der Buchstabe dann die Zahl. \n\n")
			continue

		case y > 9:
			fmt.Printf("Fehler! Y muss von 1-10 gewählt werden. \n\n")
			continue

		case x > 9:
			fmt.Printf("Fehler! X muss von A-J gewählt werden. \n\n")
			continue
		}

		return x, y, err
	}
}
