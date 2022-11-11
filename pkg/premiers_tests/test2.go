package main

import (
	"fmt"
	"math"
)

var GlobalVariable string

func sayHello(nom string) {
	fmt.Println("Hi", nom)
}

// retourne deux types int
func addThree(a int, b int) (int, int) {
	return a + 3, b + 3
}

// déclaration d'une fonction avec des paramètres infinis
func addition(param ...int) int {
	total := 0
	for _, value := range param { //j'ai mis un underscore "_" car je ne souhaite pas récupérer l'index de la range
		total += value
	}
	return total
}

func showEachIndexNumber(param ...int) {
	for index, value := range param {
		fmt.Println("index", index, "value", value)
	}
}

// Déclaration d'une fonction qui prend en paramètres un float64 et une fonction anonyme
func rajouterDix(a float64, fAnonyme func(float64) float64) /**/ {
	operation := fAnonyme(a) // Appel à notre fonction anonyme
	result := operation + 10
	fmt.Println(result)
}

func main() {
	sayHello("Bob")
	var a int
	var b int
	a, b = addThree(5, 10)
	fmt.Println("a", a, "b", b)

	showEachIndexNumber(1, 2, 3, 4, 5, 6)

	// stockage de notre fonction anonyme dans une variable
	racineCarree := func(x float64) float64 { return math.Sqrt(x) }
	rajouterDix(9, racineCarree)

	/*
	 il est possible aussi d'utiliser directement une fonction anonyme
	 dans une variable sans forcement la stocker dans une variable
	*/
	rajouterDix(5, func(x float64) float64 { return math.Pow(x, 2) })


	type Personnage struct {
        nom string
        age int
    }

    perso := Personnage{"Hatim", 20}
    fmt.Println(perso)

    // accès juste à l'attribut age de notre structure Personnage
    fmt.Println("je ne veux afficher que l'age du personnage => ", perso.age)

    perso.age = 23 // modification de l'age
    fmt.Println("3 ans plus tard ...", perso.age)


}
