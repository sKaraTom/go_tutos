package affichage

// la variable a une majuscule pour être accessible d'un autre fichier.
var Nom string = "Hatim"

// sans majuscule la variable est privée.
var sexe string = "masculin"

// la majuscule en début de nom de fonction est nécessaire pour être accessible depuis l'extérieur.
func AfficheSexe() string {
	return Nom + " est de sexe " + sexe
}
