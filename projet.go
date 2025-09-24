package main

import (
	"fmt"
	"math/rand"
)

func main() {
	source := rand.NewSource(1)
	r := rand.New(source)

	var tableau [10]int
	for i := 0; i < 10; i++ {
		tableau[i] = r.Intn(6)
	}

	var occu [6]int
	for i := 0; i < 10; i++ {
		occu[tableau[i]]++
	}

	fmt.Println(tableau)
	fmt.Println(occu)
}
