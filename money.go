package main

import (
	"fmt"
	// "zi/money"
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
	var bank shifr.RSA
	var client Clientushka
	bank.Init()

	client.n = crypto.Random(2, bank.N-1)

	client.r = crypto.Random(1, bank.N)
	for crypto.Gcd(client.r, bank.N) != 1 {
		client.r = crypto.Random(1, bank.N)
	}
	nn := crypto.Mul(IntHash(client.n), crypto.Pow(client.r, bank.D, bank.N), bank.N)

	// Bank side
	ss := crypto.Pow(nn, bank.C, bank.N)

	// Client side again
	r_inverted := crypto.Reverse(client.r, bank.N)
	s := crypto.Mul(ss, r_inverted, bank.N)

	// Bank checks banknote
	a := crypto.Pow(s, bank.D, bank.N)
	b := IntHash(client.n)
	fmt.Println("a", a)
	fmt.Println("b", b)

	fmt.Println("All given is ")

	return

}
