package main

import (
	"fmt"
	"zi/crypto"
	"zi/shifr"
	"zi/sign"
)

func main() {

	var r shifr.Elgamal

	fmt.Println(crypto.GenPrime16())

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
