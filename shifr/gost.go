package shifr

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"zi/crypto"
)

type Gost struct {
	P, B, Q, D, E, R int
	A                int
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

func (r *Gost) Init() {
	fmt.Println("Initialize")
	r.P = 4
	for !crypto.Fermat(r.P) {
		r.Q = int(crypto.GenPrime8())
		r.B = 16777216 + rand.Int()%16777216 // 24 bit random

		if r.B < 0 {
			continue
		}
		r.P = r.B*r.Q + 1
	}
	fmt.Println(r.P, r.Q, r.B)
	r.A = int(rand.Int())%r.P + 16777216
	for crypto.Pow(r.A, r.Q, r.P) != 1 {
		r.A += 1
		fmt.Println("A, Q, Pow(a, q, p)", r.A, r.Q, crypto.Pow(r.A, r.Q, r.P))
	}
	fmt.Println("A:", r.A)

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
