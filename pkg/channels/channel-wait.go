package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	now := time.Now()
	const tailleBuffer int = 10
	ch := make(chan int, tailleBuffer)

	// 5 expéditeurs
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second * 2)
			fmt.Println(<-ch)
			wg.Done()
		}()
	}

	// 10 récepteurs
	for j := 0; j < tailleBuffer; j++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second * 2)
			ch <- 50
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(now))

	// itérer sur un channel buffered
	ch2 := make(chan string, 2) // buffer de taille 2

	go func() {
		defer close(ch2) // on indique à notre compilateur qu'on a finit d'écrire sur le channel même si le taille du buffer n'est pas respecté.
		ch2 <- "test"
		// ch2 <- "test2"
	}()

	// for elem := range ch2 {
	// 	fmt.Println(elem)

	// }

	elem, ok := <-ch2
	fmt.Println(elem, ok)
	elem2, ok2 := <-ch2
	fmt.Println(elem2, ok2)
}
