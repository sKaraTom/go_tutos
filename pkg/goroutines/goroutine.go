package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

// fonction qui dure 3*1 secondes
func run(name string) {
	// defer garantit l'exécution de la méthode Done() même si ça plante (évite boucle infinie)
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second) // attendre 1 seconde
		fmt.Println(name, " : ", i)
	}
}

func main() {
	debut := time.Now()
	wg.Add(1)
	go run("Hatim")
	wg.Add(1)
	go run("Robert")
	wg.Add(1)
	go run("Alex")
	// attendre que tout le groupe ait terminé
	wg.Wait()

	fin := time.Now()
	fmt.Println(fin.Sub(debut))
}
