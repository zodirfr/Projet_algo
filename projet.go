package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var tableau [10]int
	for i := range tableau {
		tableau[i] = r.Intn(5) // nombre entre 0 et 4
	}

	var occu [5]int
	for _, v := range tableau {
		occu[v]++
	}

	fmt.Println("Tableau :", tableau)
	fmt.Println("Occurrences :", occu)
}

