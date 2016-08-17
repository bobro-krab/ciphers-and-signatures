package poker

import (
	"fmt"
	"math/rand"
	"zi/crypto"
)

const deck_size = 52
const suits_count = 4 
const card_shift = 3 // 0 and 1 is numbers, that persistent to any
// encryption and decryption, so we shift it by this constant.

func init() {
	fmt.Println("Initialized")
}

type PokerTable struct {
	PlayerCount int
	P           int
	CKeys       []int
	DKeys       []int
	Cards       []int
	Taken       []int
}

func GetCard(card int) string {
	Suits := [...]string{"♠", "♥", "♣", "♦"}
	Values := [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	value := (card - card_shift) % (deck_size / suits_count)
	suit := (card - card_shift) / (deck_size / suits_count)
	return Values[value] + Suits[suit]
}

func NewTable(playerCount int) PokerTable {
	var Table PokerTable
	Table.P = int(crypto.GenPrime16())
	Table.Cards = make([]int, deck_size)
	Table.CKeys = make([]int, playerCount)
	Table.DKeys = make([]int, playerCount)
	Table.Taken = make([]int, playerCount)
	Table.PlayerCount = playerCount

	fmt.Println("PlayerCount", playerCount)
	for i := 0; i < playerCount; i++ {
		Table.CKeys[i] = rand.Int() % (Table.P - 1)
		for crypto.Gcd(Table.CKeys[i], Table.P-1) != 1 {
			Table.CKeys[i] = rand.Int() % (Table.P - 1)
		}
		Table.DKeys[i] = crypto.Reverse(Table.CKeys[i], Table.P-1)
	}
	for i := 0; i < deck_size; i++ {
		Table.Cards[i] = i + card_shift
	}
	return Table
}

func Shuffle(cards []int) {
	for i, _ := range cards {
		j := rand.Int() % deck_size
		temp := cards[i]
		cards[i] = cards[j]
		cards[j] = temp
	}
}

func PowCards(cards []int, a, module int) {
	for i, _ := range cards {
		cards[i] = crypto.Pow(cards[i], a, module)
	}
}

func (Table *PokerTable) Play() {
	// Encypt deck
	for i := 0; i < Table.PlayerCount; i++ {
		PowCards(Table.Cards, Table.CKeys[i], Table.P)
		Shuffle(Table.Cards)
	}

	// Take cards
	for i := 0; i < Table.PlayerCount; i++ {
		index := int(rand.Int()) % deck_size
		for Table.Cards[index] == 0 {
			index = int(rand.Int()) % deck_size
		}
		Table.Taken[i] = Table.Cards[index]
		Table.Cards[index] = 0
	}

	for player := 0; player < Table.PlayerCount; player++ {
		count := 0
		for i := (player + 1) % Table.PlayerCount; count < Table.PlayerCount; i = (i + 1) % Table.PlayerCount {
			Table.Taken[player] = crypto.Pow(Table.Taken[player], Table.DKeys[i], Table.P)
			count++
		}
		fmt.Println(player, "player card is", GetCard(Table.Taken[player]))
	}

	fmt.Println("cards ", Table.Cards)
}
