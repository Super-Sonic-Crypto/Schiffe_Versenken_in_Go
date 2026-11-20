// Koordinaten Eingabe, abfragen nach Fehlerzuständen
// Achtung keine Abfrage nach den Schiffen, ob das Setzen logisch ist

// 1x Schlachtschiff (Länge 5)
// 2x Kreuzer (Länge 4)
// 3x Zerstörer (Länge 3)
// 4x U-Boot (Länge 2)

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ParseCoordinate(Input *bufio.Scanner, RunThrough int) (x, y int, err error) {
	for {
		switch RunThrough {
		case 1, 2, 3: // Zerstörer setzen
			fmt.Printf("%d. Schiff-Koordinate setzen: ", RunThrough)
			Input.Scan()
			CoordinateInput := Input.Text()
			CoordinateInput = strings.ToUpper(CoordinateInput) //Eingabe wird zu Großbuchstabe gezwungen

			if len(CoordinateInput) > 3 || len(CoordinateInput) < 2 {
				fmt.Printf("Fehler! bitte mindestens 2 oder höchstens 3 Zeichen eingeben! \n\n")
				continue
			}

			x = int(CoordinateInput[0] - 'A')
			y, err = strconv.Atoi(CoordinateInput[1:])
			y = y - 1 // Array zählt von 0-9, für bessere Rechnung wird 1 zu 0, usw.

			/*
				case 4: // Angriff
					fmt.Printf("Angriffs-Koordinate setzen: ")
					Input.Scan()
					AttackInput := Input.Text()
					AttackInput = strings.ToUpper(AttackInput)

					if len(AttackInput) != 2 {
						fmt.Printf("Fehler! bitte mindestens 2 oder höchstens 3 Zeichen eingeben! \n\n")
						continue
					}

					x = int(AttackInput[0] - 'A')
					y, err = strconv.Atoi(AttackInput[1:])
					y = y - 1 // Array zählt von 0-9, für bessere Rechnung wird 1 zu 0, usw.

			*/
		}

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
