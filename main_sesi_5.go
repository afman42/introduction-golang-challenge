package main

import (
	"fmt"
	"sync"
)

func main() {
	textBisa := []string{"bisa1", "bisa2", "bisa3"}
	textCoba := []string{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(id int) {
			if id%2 == 1 {
				fmt.Printf("%v %d\n", textBisa, id)
			} else {
				fmt.Printf("%v %d\n", textCoba, id)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
