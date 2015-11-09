package shifr

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"math/rand"
	"unsafe"
	"zi/crypto"
)

type Elgamal struct {
	P, G, C, D, E, R int
}

func (r *Elgamal) Init() {
	r.P, r.G = crypto.GenPair()
	r.C = rand.Int()%(r.P-2) + 1
	r.D = crypto.Pow(r.G, r.C, r.P)
}

func (r *Elgamal) BlockSize() int {
	return int(unsafe.Sizeof(int(1)) * 2 / unsafe.Sizeof(byte(1)))
}

func (r *Elgamal) EncryptByte(message byte) []byte {
	k := rand.Int()%(r.P-3) + 1
	r.R = crypto.Pow(r.G, k, r.P)
	r.E = int(int64(int64(message)*int64(crypto.Pow(r.D, k, r.P))) % int64(r.P))
	if r.E < 0 {
		r.E += r.P
	}
	firsthalf := make([]byte, r.BlockSize()/2)
	secondhalf := make([]byte, r.BlockSize()/2)
	binary.LittleEndian.PutUint32(firsthalf, uint32(r.E))
	binary.LittleEndian.PutUint32(secondhalf, uint32(r.R))
	firsthalf = append(firsthalf, secondhalf...)
	return firsthalf
}

func (r *Elgamal) DecryptByte(message []byte) byte {
	second := message[:r.BlockSize()/2]
	first := message[r.BlockSize()/2:]
	r.E = int(binary.LittleEndian.Uint32(first))
	r.R = int(binary.LittleEndian.Uint32(second))
	temp := r.P - 1 - r.C
	if temp < 0 {
		temp += r.P
	}
	m := int(int64(r.E) * int64(crypto.Pow(r.R, temp, r.P)) % int64(r.P))
	fmt.Println("Decrypted", m)
	return byte(m)
}

func Elgamalold(message byte) {
	fmt.Println("Message is", message)

	p, g := crypto.GenPair()
	// p, g := 23, 5
	fmt.Println("P and G is", p, g)

	c := rand.Int()%(p-2) + 1
	fmt.Println("C is", c)

	d := crypto.Pow(g, c, p)
	fmt.Println("D is", d)

	// encrypt
	k := rand.Int()%(p-3) + 1
	r := crypto.Pow(g, k, p)
	e := int(int64(int64(message)*int64(crypto.Pow(d, k, p))) % int64(p))
	if e < 0 {
		e += p
	}
	fmt.Println("Encrypted and k is ", e, k)

	// decrypt
	temp := p - 1 - c
	if temp < 0 {
		temp += p
	}
	m := int(int64(e) * int64(crypto.Pow(r, temp, p)) % int64(p))
	fmt.Println("Decrypted", m)

}

func (r *Elgamal) FileType() string {
	return "elg"
}

func (r *Elgamal) Key() []byte {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(r)
	return b.Bytes()
}

func (r *Elgamal) LoadKey(key []byte) {
	m := Elgamal{}
	b := bytes.Buffer{}
	b.Write(key)
	d := gob.NewDecoder(&b)
	err := d.Decode(&m)
	if err != nil {
		panic(err)
	}
	*r = m

}
