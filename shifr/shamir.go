package shifr

import (
	"fmt"
	"math/rand"
	"zi/crypto"
)

// инврсия не может существовать, если она не взаимно проста с модулем
func Shamir_gen(module int) (int, int) {
	c := rand.Int()
	for crypto.Gcd(c, module) != 1 {
		c = rand.Int()
	}
	_, d, _ := crypto.Euclid(c, module)
	if d < 0 {
		d += module
	}
	return c, d
}

func Shamir(message int) int32 {
	// generation
	p := int(crypto.GenPrime16())
	ca, da := Shamir_gen(p - 1)
	cb, db := Shamir_gen(p - 1)

	// checking
	temp := crypto.Pow(message, ca, p)
	temp = crypto.Pow(temp, cb, p)
	temp = crypto.Pow(temp, da, p)
	temp = crypto.Pow(temp, db, p)

	fmt.Println("Shamir Message is", temp)

	return 1
}
