package main

import (
	"fmt"
	"math/big"
	"zi/shifr"
	"zi/sign"
)

func main() {

	// a := new(big.Int)
	// fmt.Sscan("12345678901234567890", a)
	// fmt.Println(a.String())
	// return

	var r shifr.Gost

	fmt.Println("Shifrator v0.1")
	sign.SignupFile("testfile", &r)
	sign.CheckupSignature("testfile", &r)
	return

	// if len(os.Args) > 2 {
	// 	shifr.DecryptFile(os.Args[1], &r)
	// } else {
	// 	shifr.EncryptFile(os.Args[1], &r)
	// }

}
