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

func (r *RSA) Key() []byte {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(r)
	return b.Bytes()
}

func (r *RSA) LoadKey(key []byte) {
	m := RSA{}
	b := bytes.Buffer{}
	b.Write(key)
	d := gob.NewDecoder(&b)
	err := d.Decode(&m)
	if err != nil {
		panic(err)
	}
	*r = m
}

func (r *RSA) FileType() string {
	return "rsa"
}

func (r *RSA) BlockSize() int {
	return int(unsafe.Sizeof(int(4)) / unsafe.Sizeof(byte(1)))
}

func (r *RSA) Init() {
	r.N = 0
	r.Phi = 0

	r.P = int(crypto.GenPrime16())
	r.Q = int(crypto.GenPrime16())
	for r.Q == r.P {
		r.Q = int(crypto.GenPrime())
	}

	r.N = r.P * r.Q
	r.Phi = (r.P - 1) * (r.Q - 1)

	r.D = 3
	for {
		if crypto.Gcd(r.D, r.Phi) == 1 {
			break
		}
		r.D += 2
	}

	// c is a secret key
	_, r.C, _ = crypto.Euclid(r.D, r.Phi)
	if r.C < 0 {
		r.C += r.Phi
	}
}

func (r *RSA) EncryptByte(message byte) []byte {
	result := crypto.Pow(int(message), r.D, r.N)
	bs := make([]byte, r.BlockSize())
	binary.LittleEndian.PutUint32(bs, uint32(result))
	return bs
}

func (r *RSA) DecryptByte(message []byte) byte {
	m := binary.LittleEndian.Uint32(message)
	result := crypto.Pow(int(m), int(r.C), int(r.N))
	return byte(result)
}

//----------------------------------------------------
// significator
//----------------------------------------------------

func (r *RSA) GenSign(hash int) []int {
	s := crypto.Pow(hash, r.C, r.N)
	result := make([]int, 1)
	result[0] = s
	return result
}

func (r *RSA) CheckSign(sign []int, fileHash int) bool {
	temp := sign[0]
	a := crypto.Pow(temp, r.D, r.N)
	b := fileHash
	fmt.Println("RSA_Check_sign temp", temp)
	fmt.Println("a", a)
	fmt.Println("b", b)
	return b == a
}
