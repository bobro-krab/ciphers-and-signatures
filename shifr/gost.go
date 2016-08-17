package shifr

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"zi/crypto"
)

type Gost struct {
	P, B, Q, D, E, R int
	A, G, X, Y       int
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
		r.Q = crypto.GenPrimeBounds(30000, 65536)
		r.B = crypto.Random(30000, 65536)

		if r.B < 0 {
			continue
		}
		r.P = r.B*r.Q + 1
	}

	fmt.Println("P, Q, B")
	fmt.Println(r.P, r.Q, r.B)
	r.A = 2
	for 1 == 1 {
		r.G = crypto.GenPrimeBounds(2, r.P)
		r.A = crypto.Pow(r.G, r.B, r.P)
		if r.A > 1 {
			break
		}
	}
	fmt.Println("A:", r.A)
	r.X = crypto.Random(1, r.Q)
	r.Y = crypto.Pow(r.A, r.X, r.P)
	fmt.Println("Initialization complete")
}

func (r *Gost) GenSign(hash int) []int {
	fmt.Println("Generating sign")
	hash %= r.Q
	for hash < 0 {
		hash += r.Q
	}
	R := 0
	S := 0
	for {
		k := crypto.Random(1, r.Q)
		R = crypto.Pow(r.A, k, r.P) % r.Q
		if R == 0 {
			fmt.Println("Bad R k", R, k)
			continue
		}
		S = crypto.Mul(k, hash, r.Q) + crypto.Mul(r.X, R, r.Q)
		if S == 0 {
			fmt.Println("Bad R S k", R, S, k)
			continue
		}
		break
	}
	result := make([]int, 3)
	result[0] = R
	result[1] = S
	result[2] = 3
	return result
}

func (r *Gost) CheckSign(sign []int, fileHash int) bool {
	fmt.Println("Checking signature")
	fileHash %= r.Q
	for fileHash < 0 {
		fileHash += r.Q
	}
	R := sign[0]
	for R < 0 {
		R += r.Q
	}
	S := sign[1]
	hash_1 := crypto.Reverse(fileHash, r.Q)
	u_1 := crypto.Mul(S, hash_1, r.Q)
	u_2 := crypto.Mul(R, hash_1, r.Q)
	v := crypto.Mul(crypto.Pow(r.A, u_1, r.P), crypto.Pow(r.Y, u_2, r.P), r.P) % r.Q
	fmt.Println("R, v", R, v)
	return R == v
}

func (r *Gost) FileType() string {
	return "gost"
}
