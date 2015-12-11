package shifr

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"math/big"
	// "math/rand"
	"zi/crypto"
)

type Gost struct {
	P, B, Q, D, E, R big.Int
}

func (r *Gost) Key() []byte {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(r)
	return b.Bytes()
}

func (r *Gost) LoadKey(key []byte) {
	m := Gost{}
	b := bytes.Buffer{}
	b.Write(key)
	d := gob.NewDecoder(&b)
	err := d.Decode(&m)
	if err != nil {
		panic(err)
	}
	*r = m
}

func Pow(A, B, Module big.Int) big.Int {
	a := A
	b := B
	module := Module
	var result big.Int = 1
	step_count := int(math.Log2(float64(b)))
	for i := 0; i <= step_count; i++ {
		if b%2 == 1 {
			result = (result * a) % module
		}
		b /= 2
		a = (a * a) % module
	}
	return result
}

func Fermat(n big.Int) bool {
	z := new(big.Int)
	a := new(big.Int)

	z.SetString("0", 10)
	if n.Cmp(z) < 0 {
		return false
	}

	z.SetString("2", 10)
	if n.Cmp(z) < 0 {
		return true
	}
	for i := 0; i < 100; i++ {
		// fmt.Println("Fermat n i", n, i)
		a, _ = rand.Int(rand.Reader, &n)
		if Pow(a, n-1, n) != 1 {
			return false
		}
		if Gcd(a, n) != 1 {
			return false
		}
	}
	return true

}

func (r *Gost) Init() {
	fmt.Println("Initialize")
	r.P.SetString("4", 10)
	for !crypto.Fermat(r.P) {
		r.Q = int(crypto.GenPrime16())
		// r.B = rand.Int() % 65536 // 16 bit random
		r.B = rand.Int() % 256 // 24 bit random

		if r.B < 0 {
			continue
		}
		r.P = r.B*r.Q + 1
	}
	fmt.Println(r.P, r.Q, r.B)
	fmt.Println("done")
}

func (r *Gost) GenSign(hash int) []int {
	result := make([]int, 3)
	result[0] = 1
	result[1] = 2
	result[2] = 3
	return result
}

func (r *Gost) CheckSign(sign []int, fileHash int) bool {
	// s := sign[0]
	// y := sign[1]
	// R := sign[2]
	return false
}

func (r *Gost) FileType() string {
	return "elg"
}
