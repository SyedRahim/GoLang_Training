package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var firstName = "syed"
	var lastName = "ibrahim"
	fmt.Println("This is my first program in Go using GoLand!!")
	fmt.Println(rand.Intn(100))
	fmt.Println(firstName + " " + lastName)
	fmt.Println(rand.Intn(1000))
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is even number\n", i)
		} else {
			fmt.Printf("%d is odd number\n", i)
		}
	}
}
