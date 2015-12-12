package poker

import (
	"fmt"
	"math/rand"
	"zi/crypto"
)

func init() {
	fmt.Println("Initialized")
}

type PokerTable struct {
	PlayerCount int
	P           int
	CKeys       []int
	DKeys       []int
	Cards       []int
}

func GetRaw(card int) int {
	return card % 13
}

func GetValue(card int) int {
	return card % 4
}

func NewTable(playerCount int) PokerTable {
	var Table PokerTable
	Table.P = int(crypto.GenPrime16())
	Table.Cards = make([]int, 52)
	Table.CKeys = make([]int, playerCount)
	Table.DKeys = make([]int, playerCount)
	fmt.Println("Playing", playerCount)
	for i := 0; i < playerCount; i++ {
		Table.CKeys[i] = rand.Int() % (Table.P - 1)
		Table.DKeys[i] = crypto.Reverse(Table.CKeys[i], Table.P-1)
		fmt.Println(i, "player generated", Table.CKeys[i], Table.DKeys[i])
	}
	for i := 0; i < 52; i++ {
		Table.Cards[i] = i
	}
	return Table
}

func Shuffle(cards []int) {
	for _, v := range cards {
		j := rand.Int() % 52
		v, cards[j] = cards[j], v
	}
}

func (Table *PokerTable) Play() {
	fmt.Println(Table.P, "general option")
	fmt.Println("Before shuffle", Table.Cards)
	Shuffle(Table.Cards)
	fmt.Println("After shuffle", Table.Cards)
}
