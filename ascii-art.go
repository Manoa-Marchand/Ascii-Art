package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	Args := os.Args
	nbArgs := len(Args)

	if nbArgs != 2 {

		fmt.Println("ERREUR: Veuillez entrer un argument")

	} else {

		Arr := []rune(os.Args[1])
		for _, i := range Arr {

			if i < 32 || i > 127 {

				fmt.Println("ERREUR: Argument non valide")
				return
			}
		}
		lAscii(Args)
	}
}

//fonction principale

func lAscii(Args []string) {

	Arg := []rune(Args[1])
	var nligne int
	var ligne int

	for i1 := 0; i1 < 8; i1++ {

		for i2 := 0; i2 < len(Arg); i2++ {

			nligne = 0
			ligne = GLigne(Arg[i2])
			file, err := os.Open("standard.txt")

			if err != nil {

				fmt.Println("ERREUR: " + err.Error())

			} else {

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {

					if nligne == ligne+i1 {

						fmt.Print(scanner.Text())

					}

					nligne++

				}
			}
		}
		fmt.Println()
	}
}

//Fonction pour récupérer la ligne du caractère

func GLigne(char rune) int {

	var ligne int

	for i := 0; i < 95; i++ {

		if rune(i+32) == char {

			ligne = i
			break
		}
	}
	ligne = ligne*9 + 1
	return ligne
}
