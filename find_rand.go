package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var choose int

	var random = rand.Intn(99) + 1

	var a = true

	for a == true {
		fmt.Print("Choisi un nombre   ")
		fmt.Scanf("%d", &choose)
		if choose == random {
			fmt.Println("Bravo")
			a = false
		} else if choose > random {
			fmt.Println("Trop grand")
		} else if choose < random {
			fmt.Println("Trop petit")
		}
	}
}
