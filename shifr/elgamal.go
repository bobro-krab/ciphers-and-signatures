package shifr

import (
	"fmt"
	"math/rand"
	"zi/crypto"
)

func Elgamal(message byte) {
	fmt.Println("Message is", message)

	p, g := crypto.GenPair()
	fmt.Println("P and G is", p, g)

	c := rand.Int()%(p-2) + 1
	fmt.Println("C is", c)

	d := crypto.Pow(g, c, p)
	fmt.Println("D is", d)

	// encrypt
	k := rand.Int()%(p-3) + 1
	r := crypto.Pow(g, k, p)
	e := int(int64(int(message)*crypto.Pow(d, k, p)) % int64(p))
	if e < 0 {
		e += p
	}
	fmt.Println("Encrypted and k is ", e, k)

	// decrypt
	m := (e * crypto.Pow(r, p-1-c, p)) % p
	fmt.Println("Decrypted", m)

}
