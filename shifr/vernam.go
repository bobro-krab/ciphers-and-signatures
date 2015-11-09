package shifr

import (
	// "fmt"
	// "zi/crypto"
	// "bytes"

	"math/rand"
)

type Vernam struct {
	ByteKey     []byte
	CurrentByte int
}

func (r *Vernam) Init() {
	r.ByteKey = make([]byte, 0)
	r.CurrentByte = 0
}

func (r *Vernam) FileType() string {
	return "ver"
}

func xor(a, b byte) byte {
	return a ^ b
}

func (r *Vernam) EncryptByte(message byte) []byte {
	k := byte(rand.Int())
	r.ByteKey = append(r.ByteKey, k)
	res := make([]byte, 1)
	res[0] = xor(message, k)
	return res
}

func (r *Vernam) DecryptByte(b []byte) byte {
	x := xor(r.ByteKey[r.CurrentByte], b[0])
	// fmt.Println("current byte", r.CurrentByte)
	r.CurrentByte += 1
	return x
}

func (r *Vernam) BlockSize() int {
	return 1
}

func (r *Vernam) Key() []byte {
	return r.ByteKey
}

func (r *Vernam) LoadKey(k []byte) {
	r.ByteKey = k
}
