package main

import "fmt"

func main() {

	// function make pour déclarer en donnant nombre d'éléments initiaux
	slice1 := make([]string, 0)
	fmt.Println(slice1)
	// déclaration comme tableau sans taille --> slice
	var slice2 []string

	// ajouter un élément
	slice1 = append(slice1, "ajout1")
	fmt.Println(slice1)

	// copier une slice
	var sliceCopiee []string
	copy(slice1, sliceCopiee)
	fmt.Println(slice1)

}
