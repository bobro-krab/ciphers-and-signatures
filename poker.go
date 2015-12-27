package main

import (
	"fmt"
	"zi/poker"
)

func main() {

	fmt.Println("Poker v0.1")
	Table := poker.NewTable(3)
	Table.Play()
	return

}
