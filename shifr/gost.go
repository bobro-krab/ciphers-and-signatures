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
		r.Q = crypto.GenPrime()
		r.B = rand.Int()
		r.P = r.B*r.Q + 1
	}
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
