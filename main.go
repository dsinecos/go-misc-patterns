package main

import (
	"fmt"

	"github.com/dsinecos/go-misc-patterns/dice"
)

func main() {
	diceValue := make(chan int)

	dice.Throw(diceValue)

	select {
	case a := <-diceValue:
		fmt.Printf("Dice rolled to %v \n", a)
	}

	// To allow all the goroutines to execute and exit
	fmt.Scanln()
}
