package main

import (
	"fmt"
	"zi/poker"
	// "zi/shifr"
	// "zi/sign"
)

func main() {
	fmt.Println("Poker v0.1")
	Table := poker.NewTable(3)
	Table.Play()
	return

	// var r shifr.Gost
	// fmt.Println("Shifrator v0.1")
	// sign.SignupFile("testfile", &r)
	// sign.CheckupSignature("testfile", &r)
	// return

	// if len(os.Args) > 2 {
	// 	shifr.DecryptFile(os.Args[1], &r)
	// } else {
	// 	shifr.EncryptFile(os.Args[1], &r)
	// }

}
