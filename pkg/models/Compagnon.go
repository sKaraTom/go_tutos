package models

import "fmt"

type Compagnon struct {
	Nom        string
	Vie        int
	Puissance  int
	Mort       bool
	Inventaire [3]string
}

/*
Créer une instance de la classe Compagnon

@return: struct Compagnon
*/
func NewCompagnon(Nom string, Vie int, Puissance int, Mort bool, Inventaire [3]string) Compagnon {
	compagnon := Compagnon{Nom, Vie, Puissance, Mort, Inventaire}
	return compagnon
}

func (c Compagnon) AfficherNom() {
	fmt.Println("Nom du personnage :", c.Nom)
}

/*
Affiche des informations sur un personnage

@return: void
*/
func (c Compagnon) AffichageInfos() { // déclaration de ma méthode Affichage() liée à ma structure Compagnon
	fmt.Println("--------------------------------------------------")
	fmt.Println("Vie du personnage", c.Nom, ":", c.Vie)
	fmt.Println("Puissance du personnage", c.Nom, ":", c.Puissance)

	if c.Mort {
		fmt.Println("Vie du personnage", c.Nom, "est mort")
	} else {
		fmt.Println("Vie du personnage", c.Nom, "est vivant")
	}

	fmt.Println("\nLe personnage", c.Nom, "possède dans son inventaire :", c.Vie)

	for _, item := range c.Inventaire {
		fmt.Println("-", item)
	}
}
