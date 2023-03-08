package main

import "fmt"

func main() {
	var i uint8 = 21
	var persen string = "%"
	var j bool = true
	var unicodeRussia string = "Я"
	var nilaiBase10 uint = 21
	nilaiBase8 := 031
	var nilaiBase16 uint8 = 15
	var floatDesimal float32 = 123.456

	fmt.Printf("%d \n", i)
	fmt.Printf("%s \n", persen)
	fmt.Printf("%t \n", j)
	fmt.Printf("%q \n", unicodeRussia)
	fmt.Printf("%d \n", nilaiBase10)
	fmt.Println(nilaiBase8)
	fmt.Printf("%x \n", nilaiBase16)
	fmt.Printf("%X \n", nilaiBase16)
	fmt.Printf("%+q \n", unicodeRussia)
	fmt.Printf("%f \n", floatDesimal)
	fmt.Printf("%e \n", floatDesimal)
}
