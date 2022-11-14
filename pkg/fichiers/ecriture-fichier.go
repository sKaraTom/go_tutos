package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// chemin relatif par rapport à la racine projet
	file, err := os.OpenFile("./pkg/fichiers/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close() // on ferme automatiquement à la fin de notre programme

	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("test\n") // écrire dans le fichier
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("i love test\n")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("./pkg/fichiers/test.txt") // lire le fichier
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(data))
}
