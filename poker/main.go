package main

import (
	"fmt"
	"zi/shifr"
	"zi/sign"
)

func main() {

	var r shifr.Gost

	fmt.Println("Poker v.1.0")
	sign.SignupFile("testfile", &r)
	sign.CheckupSignature("testfile", &r)
	return

	// if len(os.Args) > 2 {
	// 	shifr.DecryptFile(os.Args[1], &r)
	// } else {
	// 	shifr.EncryptFile(os.Args[1], &r)
	// }

}
