package shifr

import (
	"fmt"
	// "math"
	"math/rand"
	"time"
	"zi/crypto"
)

func init() {
	rand.Seed(int64(time.Now().Second()))
}

type RSA struct {
	N int
	D int

	Phi int
	C   int
	P   int
	Q   int
}

func RSA_Init(in *RSA) {
	var r RSA = *in
	r.N = 0
	r.Phi = 0

	r.P = int(crypto.GenPrime16())
	r.Q = int(crypto.GenPrime16())
	for r.Q == r.P {
		r.Q = int(crypto.GenPrime())
	}
	fmt.Println("P is", r.P, "Q is", r.Q)

	r.N = r.P * r.Q
	fmt.Println("N is", r.N)

	r.Phi = (r.P - 1) * (r.Q - 1)
	fmt.Println("phi is", r.Phi)

	r.D = 3
	for {
		if crypto.Gcd(r.D, r.Phi) == 1 {
			break
		}
		r.D += 2
	}
	fmt.Println("D is", r.D)

	// c is a secret key
	_, r.C, _ = crypto.Euclid(r.D, r.Phi)
	fmt.Println("calculaatig for ", r.D, r.Phi)
	if r.C < 0 {
		r.C += r.Phi
	}
	fmt.Println("C is", r.C)
	*in = r
}

func RSA_Encrypt(message byte, r RSA) int {
	result := crypto.Pow(int(message), r.D, r.N)
	return result
}

func RSA_Decrypt(message int, r RSA) byte {
	result := crypto.Pow(message, int(r.C), int(r.N))
	return byte(result)
}
