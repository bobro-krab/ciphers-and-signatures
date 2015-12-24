package main

import (
	"fmt"
	"zi/shifr"
	"zi/sign"
)

func main() {

	fmt.Println("Shifrator v0.3")

	var b shifr.RSA
	fmt.Println("\n\nRSA")
	sign.SignupFile("testfile", &b)
	sign.CheckupSignature("testfile", &b)

	var c shifr.Elgamal
	fmt.Println("\n\nElgamal")
	sign.SignupFile("testfile", &c)
	sign.CheckupSignature("testfile", &c)

	fmt.Println("\n\nGOST")
	var r shifr.Gost
	sign.SignupFile("testfile", &r)
	sign.CheckupSignature("testfile", &r)
}
