package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var n int

	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100) + 1
	i:= false
	for i == false {
		fmt.Print("Choisi un nombre ")
		fmt.Scanf("%d", &n)

		if randomNumber == n {
			fmt.Println("Bravo tu as gagné.")
			i = true
		} 
		if randomNumber < n {
			fmt.Println("Ton nombre est supérieur au nombre mystère.")
		}
		if randomNumber > n {
			fmt.Println("Ton nombre est inférieur au nombre mystère.")
		}
	}
}
