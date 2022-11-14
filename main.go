package main

import (
	"fmt"

	"github.com/sKaraTom/go_tutos/pkg/affichage"
	. "github.com/sKaraTom/go_tutos/pkg/models"
)

func main() {
	// Affichage
	fmt.Println("Je m'appelle", affichage.Nom)
	fmt.Println(affichage.AfficheSexe())

	// Classe Compagnon
	magicien := Compagnon{ // Instanciation de la classe Compagnon
		Nom:        "Gandalf",
		Vie:        100,
		Puissance:  20,
		Mort:       false,
		Inventaire: [3]string{"potions", "poisons", "bâton"},
	}
	magicien.AfficherNom()
	magicien.AffichageInfos()

	// compagnon instancié par méthode de constructeur.
	guerrier := NewCompagnon("Aragorn", 200, 100, false, [3]string{"potions", "poisons", "épée"})
	guerrier.AffichageInfos()

	tuerCompagnon := func(c *Compagnon) {
		c.Vie = 0
		c.Mort = true
		c.AffichageInfos()
	}

	tuerCompagnon(&guerrier)

	fmt.Println("après mort")
	guerrier.AffichageInfos()

}
