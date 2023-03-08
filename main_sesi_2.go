package main

import "fmt"

func main() {
	var iFor int = 10
	var character string = "CA日APBO"
	// for
	for i := 0; i < iFor; i++ {
		if i < 5 {
			fmt.Println("Nilai i = ", i)
		}
	}
	for j := 0; j < iFor; j++ {
		if j < 5 {
			fmt.Println("Nilai j = ", j)
		}
	}
	for index, word := range character {
		fmt.Printf("character %#U starts at byte position %d\n", word, index)
	}
	for j := 6; j <= iFor; j++ {
		if j <= 10 {
			fmt.Println("Nilai j = ", j)
		}
	}
}
