package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
	   Déclaration des variables dynamiques
	*/
	flt := 15.5   //  sera automatiquement de type float
	in := 5       //  sera automatiquement de type int
	st := "hello" //  sera automatiquement de type string
	bol := true   //  sera automatiquement de type boolean

	fmt.Printf("Le type de la varialbe flt est %T\n", flt)
	fmt.Printf("Le type de la varialbe in est %T\n", in)
	fmt.Printf("Le type de la varialbe st est %T\n", st)
	fmt.Printf("Le type de la varialbe bol est %T\n", bol)

	var a int = 4
	a++ // incrémentation
	fmt.Println("incrémentation de 1 : ", a)

	var x int = 50
	var y float32 = 30.5

	fmt.Printf("x + y = ", float32(x)+y) // convertir le type de la variable x de int en float32

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez un nombre entier : ")
	scanner.Scan()

	nbr, _ := strconv.Atoi(scanner.Text()) // conversion du type string en int

	// time, err := strconv.Atoi(times)
	// if err != nil {
	// 	// Add code here to handle the error!
	// }

	fmt.Printf("res : %d\n", (nbr + 6))

	max := 20
	nombre := 0

	for true { // boucle infinie

		if nombre < max {
			fmt.Println("nombre ", nombre, " plus petit que ", max, " on continue à itérer")
			nombre++
			continue
		} else {
			fmt.Println("Max atteint on quitte la boucle")
			break // on quitte la boucle
		}
	}

}
