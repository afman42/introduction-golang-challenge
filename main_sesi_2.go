package main

import "fmt"

func main() {
	var (
		iFor int    = 10
		word string = "CA日APBO"
	)

	// for looping Nilai i
	for i := 0; i < iFor; i++ {
		if i < 5 {
			fmt.Println("Nilai i = ", i)
		}
	}

	// for looping Nilai j
	for j := 0; j < iFor; j++ {
		if j < 5 {
			fmt.Println("Nilai j = ", j)
		}
	}

	// for looping character
	for index, alphabet := range word {
		fmt.Printf("character %#U starts at byte position %d\n", alphabet, index)
	}

	// for looping Nilai j
	for j := 6; j <= iFor; j++ {
		if j <= 10 {
			fmt.Println("Nilai j = ", j)
		}
	}
}
