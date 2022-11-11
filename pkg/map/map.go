package main

import "fmt"

func main() {

	// déclaration
	// var notes map[string]int
	var notes2 = make(map[string]int)
	notes3 := map[string]int{"Hatim": 20, "Alex": 18}

	// ajouter un élément
	notes3["Hatim"] = 100
	fmt.Println(notes3["Hatim"])

	notes2["Emilie"] = 1000
	fmt.Println(notes2["Emilie"])
	// copier une slice

}
