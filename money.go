package main

import (
	"fmt"
	"zi/crypto"
	"zi/shifr"
)

type Monetary interface {
	CheckBanknote([]int) bool
	MakeBanknote() []int

	FileType() string
	Init()
	Key() []byte
	LoadKey([]byte)
}

func IntHash(a int) int {
	return a
}

type Clientushka struct {
	n, r int // so secret, wow
}

func main() {

	fmt.Println("Money v0.1")
	var client Clientushka
	values := [...]string{"100", "500", "1000", "5000", "1M"}

	bank := make([]shifr.RSA, 5)

	for i := 0; i < 5; i++ {
		fmt.Println("\nCurrent banknote: ", values[i])
		bank[i].Init()

		// client side, encrypt banknote
		client.n = crypto.Random(2, bank[i].N-1)
		client.r = crypto.Random(1, bank[i].N)
		for crypto.Gcd(client.r, bank[i].N) != 1 {
			client.r = crypto.Random(1, bank[i].N)
		}
		nn := crypto.Mul(IntHash(client.n), crypto.Pow(client.r, bank[i].D, bank[i].N), bank[i].N)

		// bank[i] side, encrypt with bank key
		ss := crypto.Pow(nn, bank[i].C, bank[i].N)

		// Client side again, decrypte bankote
		r_inverted := crypto.Reverse(client.r, bank[i].N)
		s := crypto.Mul(ss, r_inverted, bank[i].N)

		// bank[i] checks banknote, for its authenticity
		a := crypto.Pow(s, bank[i].D, bank[i].N)
		b := IntHash(client.n)
		fmt.Println("a", a)
		fmt.Println("b", b)
	}

	return

}
