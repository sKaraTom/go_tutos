package main

import "fmt"

type Forme interface { // création de L'interface Forme
	air() float64       // signature de la méthode Air()
	perimetre() float64 // signature de la méthode Perimetre()
}

/*----------------Début Structure Cercle-----------------*/
type Rectangle struct {
	largeur  float64
	longueur float64
}

/*
Pour implémenter une interface dans Go, il suffit
d'implémenter toutes les méthodes de l'interface. Ici on
implémente la méthode Air() de l'interface Forme.
*/
func (r Rectangle) air() float64 {
	return r.largeur * r.longueur
}

/*
On implémente la méthode Perimetre() de l'interface Forme
*/
func (r Rectangle) perimetre() float64 {
	return 2 * (r.largeur * r.longueur)
}

/*----------------Fin Structure Cercle-----------------*/

func main() {

	var f Forme = Rectangle{5.0, 4.0} // affectation de la structure Rectangle à l'interface Forme
	r := Rectangle{5.0, 4.0}
	fmt.Println("Valeur de f :", f)

	fmt.Println("Aire du rectangle f :", f.air())
	fmt.Println("Aire du rectangle r :", r.air())
	fmt.Println("f == r ? ", f == r)

}
