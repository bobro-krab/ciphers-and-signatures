package shifr

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"math/rand"
	"time"
	"unsafe"
	"zi/crypto"
)

type Elgamal struct {
	P, G, C, D, E, R int
}

func init() {
	rand.Seed(int64(time.Now().Second()))
}

func (r *Elgamal) Init() {
	fmt.Println("INIT")
	r.P, r.G = crypto.GenPair()
	r.C = rand.Int()%(r.P-2) + 1    // x in signature
	r.D = crypto.Pow(r.G, r.C, r.P) // y in signature
	fmt.Println("INIT DONE...")
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

func (r *Elgamal) GenSign(hash int) []int {
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

func (r *Elgamal) CheckSign(sign []int, fileHash int) bool {
	fmt.Println("Checking sign")
	s := sign[0]
	y := sign[1]
	R := sign[2]
	hash1 := crypto.Pow(r.D, y, r.P) * crypto.Pow(y, s, r.P) % r.P
	hash2 := crypto.Pow(R, fileHash, r.P)
	return hash1 == hash2
}
