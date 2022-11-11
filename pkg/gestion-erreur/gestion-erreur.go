package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func division() {
	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entre un chiffre : ")
		scanner.Scan()
		nbr, err := strconv.Atoi(scanner.Text())
		if err != nil { // Gestion de l'erreur de la fonction  Atoi()
			fmt.Println("Vous devenez rentrer un nombre et non une chaîne de caractères !")
		} else if nbr <= 0 {
			fmt.Println("[division par zéro impossible] Votre valeur doit être supérieur ou égal à 0")
		} else {
			fmt.Println("Résultat :", 1000/nbr)
			break
		}
	}
}

func verificationDivision(nbr float64) (float64, error) {
	if nbr <= 0 {
		return 0, errors.New("Erreur: Il est impossible de diviser par 0 !") //création de l'erreur
	} else {
		return nbr, nil // on retourne nil si aucune erreur est détectée
	}
}

func main() {

	fmt.Println(aucunpackage.affiche())

	nbr := 0.0
	nbr, err := verificationDivision(0.0)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Aucune erreur, nombre:", nbr)
	}

	division()
	fmt.Println("Fin")

}
