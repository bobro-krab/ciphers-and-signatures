package crypto

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().Second()))
}

func Pow64(A, B, Module int64) int64 {
	a := A
	b := B
	module := Module
	var result int64 = 1
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

// Multiple numbers without overflow by module
func Mul(A, B, module int) int {
	a := int64(A)
	b := int64(B)

	return int((a * b) % int64(module))
}

/*
   1) Функция быстрого возведения числа в степень по модулю.
*/
func Pow(A, B int, Module int) int {
	a := int64(A)
	b := int64(B)
	module := int64(Module)
	var result int64 = 1
	step_count := int(math.Log2(float64(b)))
	var i int = 0
	for i = 0; i <= step_count; i++ {
		if b%2 == 1 {
			result = (result * a) % module
		}
		b /= 2
		a = (a * a) % module
	}
	return int(result)
}

/*
   4) Функция, которая решает задачу нахождения дискретного логарифма при
      помощи алгоритма «Шаг младенца, шаг великана».
*/
func BabystepGiantstep(a, module, y int) int {
	fmt.Println("A is", a, ", module is", module, ", y is", y)
	k := int(math.Sqrt(float64(module))) + 1
	m := k
	m, k = k, m

	fmt.Println("m and k is", m, k)
	first := make(map[int]int)
	second := make(map[int]int)
	var i int = 3

	for i = 0; i <= m; i++ {
		first[(Pow(a, i, module)*y)%module] = i
	}
	for i = 1; i <= k; i++ {
		second[Pow(a, i*m, module)] = i
	}
	fmt.Println("first and second", first, second)
	i, j := find_indexes_m(first, second)
	fmt.Println("indexes are", i, j)
	x := i*m - j
	return x
}

func Reverse(a, module int) int {
	_, k_1, _ := Euclid(a, module)
	if k_1 < 0 {
		k_1 += module
	}
	return k_1
}

/*
   2) Функция, реализующая обобщённый алгоритм Евклида. Функция должна
      позволять находить наибольший общий делитель и обе неизвестных из
      уравнения.
*/
func Euclid(a, b int) (int, int, int) {
	U := [3]int{a, 1, 0}
	V := [3]int{b, 0, 1}

	for V[0] != 0 {
		q := U[0] / V[0]
		T := [3]int{U[0] % V[0], U[1] - q*V[1], U[2] - q*V[2]}
		U = V
		V = T
	}
	return U[0], U[1], U[2]
}

func Euclid64(a, b int64) (int64, int64, int64) {
	U := [3]int64{a, 1, 0}
	V := [3]int64{b, 0, 1}

	for V[0] != 0 {
		q := U[0] / V[0]
		T := [3]int64{U[0] % V[0], U[1] - q*V[1], U[2] - q*V[2]}
		U = V
		V = T
	}
	return U[0], U[1], U[2]
}

/*
   3) Функция построения общего ключа для двух абонентов по схеме Диффи-
      Хеллмана
*/
func DiffyHelman() {
	// choosing module
	module := GenPrime()
	fmt.Println("choosed module ", module)

	// choosing p and g
	p, g := GenPair()
	fmt.Println("generated parameters p =", p, "g =", g)

	// chooose secret numbers
	x_a := GenPrime()
	x_b := GenPrime()
	fmt.Println("choosed secret keys:", x_a, x_b)

	// calculating open keys with secret keys
	y_a := Pow(g, x_a, p)
	y_b := Pow(g, x_b, p)
	fmt.Println("generated open keys:", y_a, y_b)

	// checking results
	z_ab := Pow(y_b, x_a, p)
	z_ba := Pow(y_a, x_b, p)
	fmt.Println("connection key is similar: ", z_ab == z_ba, z_ab)
}

func Gcd64(a int64, b int64) int64 {
	var r int64
	for b != 0 {
		r = a % b
		a = b
		b = r
	}
	return a
}

// Great common divior
func Gcd(a int, b int) int {
	var r int
	for b != 0 {
		r = a % b
		a = b
		b = r
	}
	return a
}

// test for prime number
func Fermat(n int) bool {
	if n <= 0 {
		return false
	}
	if n <= 2 {
		return true
	}
	for i := 0; i < 100; i++ {
		// fmt.Println("Fermat n i", n, i)
		var a int = 3
		a = int(rand.Int())%(n-1) + 1
		if Pow(a, n-1, n) != 1 {
			return false
		}
		if Gcd(a, n) != 1 {
			return false
		}
	}
	return true
}

func GenPrime8() int8 {
	a := 64 + int8(rand.Int())
	for !Fermat(int(a)) {
		a = 64 + int8(rand.Int())
	}
	return a
}

func GenPrime16() int16 {
	a := int16(rand.Int())
	for !Fermat(int(a)) {
		a = int16(rand.Int())
	}
	return a
}

func GenPrime() int {
	a := int(rand.Int())
	for !Fermat(a) {
		a = int(rand.Int())
	}
	return a
}

// return 2 prime numbers p and q, where p = 2q + 1
func GenPair() (int, int) {
	var p int = 0
	var q int = 0
	var g int = 1
	for 1 == 1 {
		p = int(GenPrime16())
		q = (p - 1) / 2
		if Fermat(q) {
			break
		}
	}

	for g = 1; g < p-1; g++ {
		if Pow(g, q, p) != 1 {
			return p, g
		}
	}

	return -1, -1
}

func find_indexes_m(first map[int]int, second map[int]int) (int, int) {
	for k := range first {
		if _, ok := second[k]; ok {
			return first[k], second[k]
		}
	}
	return 0, 0
}
