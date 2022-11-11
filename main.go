package main

import (
	"fmt"

	"github.com/sKaraTom/go_tutos/pkg/affichage"
)

func main() {
	fmt.Println("Je m'appelle", affichage.Nom)
	fmt.Println(affichage.AfficheSexe())
}
