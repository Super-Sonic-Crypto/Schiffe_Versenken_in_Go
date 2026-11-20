// Funktion für das Darstellen des Spielfelds
/*
	   A  B  C  D  E  F  G  H  I  J
	1  -  -  -  -  -  -  -  -  -  -
	2  -  -  -  -  -  -  -  -  -  -
	3  -  -  -  -  -  -  -  -  -  -
	4  -  -  -  -  -  -  -  -  -  -
	5  -  -  -  -  -  -  -  -  -  -
	6  -  -  -  -  -  -  -  -  -  -
	7  -  -  -  -  -  -  -  -  -  -
	8  -  -  -  -  -  -  -  -  -  -
	9  -  -  -  -  -  -  -  -  -  -
   10  -  -  -  -  -  -  -  -  -  -
*/

package main

import "fmt"

func RenderBoard(BoardSize int, PlayingField [10][10]string, RadarField [10][10]string) {

	// Befüllung des Arrays am Anfang mit "-" (leer)
	for i := range BoardSize {
		for j := range BoardSize {
			// Abfragen sind getrennt und nicht &&, sonst würde nach dem 1. Beschreiben von PlayingField unwahr sein
			if PlayingField[i][j] == "" {
				PlayingField[i][j] = "-"
			}

			if RadarField[i][j] == "" {
				RadarField[i][j] = "-"
			}
		}
	}

	//=====	Ausgabe Spielfeld =====
	fmt.Printf("Dein Spielfeld: \n")
	fmt.Printf("  ")
	for coordinate := 'A'; coordinate < 'A'+10; coordinate++ {
		fmt.Printf("  %c", coordinate)
	}
	fmt.Println()

	for i := range BoardSize {
		if i != 9 { // Abfrage wegen der Formatierung, sonst ist die 10 komisch aus der Reihe!
			fmt.Printf(" %d  ", i+1)
		} else {
			fmt.Printf("%d  ", i+1)
		}

		for j := range BoardSize {
			fmt.Print(PlayingField[i][j], "  ")
		}
		fmt.Println()
	}
	fmt.Printf("\n\n")
	// =====

	// ===== Ausgabe Radarfeld =====
	fmt.Printf("Radarfeld: \n")
	fmt.Printf("  ")
	for coordinate := 'A'; coordinate < 'A'+10; coordinate++ {
		fmt.Printf("  %c", coordinate)
	}
	fmt.Println()

	for i := range BoardSize {
		if i != 9 { // Abfrage wegen der Formatierung, sonst ist die 10 komisch aus der Reihe!
			fmt.Printf(" %d  ", i+1)
		} else {
			fmt.Printf("%d  ", i+1)
		}

		for j := range BoardSize {
			fmt.Print(RadarField[i][j], "  ")
		}
		fmt.Println()
	}
	//=====
}
