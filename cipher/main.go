package main

import (
	"fmt"
	"os"
	"zi/shifr"
	"zi/sign"
)

func main() {
	var r shifr.RSA
	fmt.Println("Shifrator v0.1")
	sign.SignupFile("testfile", &r)
	sign.CheckSign("testfile", &r)

	return

	if len(os.Args) > 2 {
		shifr.DecryptFile(os.Args[1], &r)
	} else {
		shifr.EncryptFile(os.Args[1], &r)
	}

}
