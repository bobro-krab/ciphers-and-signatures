package shifr

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"
	"unsafe"
	"zi/crypto"
)

type Gost struct {
	P, B, Q, D, E, R int
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
	r.Q = crypto.GenPrime()
}

func (r *Gost) GenSign(hash int) []int {
	fmt.Println("Generating sign")
	hash %= r.P

	x := rand.Int()%(r.P-2) + 1
	y := crypto.Pow(r.G, x, r.P) // public

	k := rand.Int()%(r.P-2) + 1
	R := crypto.Pow(r.G, k, r.P)

	u := (hash - x*R%(r.P-1)) % (r.P - 1)

	_, k_1, _ := crypto.Euclid(k, r.P-1)
	s := k_1 * u % (r.P - 1)

	result := make([]int, 3)
	result[0] = s
	result[1] = y
	result[2] = R
	return result
}

func (r *Gost) CheckSign(sign []int, fileHash int) bool {
	fmt.Println("Checking sign")
	s := sign[0]
	y := sign[1]
	R := sign[2]
	hash1 := crypto.Pow(r.D, y, r.P) * crypto.Pow(y, s, r.P) % r.P
	hash2 := crypto.Pow(R, fileHash, r.P)
	return hash1 == hash2
}
