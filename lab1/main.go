package main

import (
	"fmt"
	"os"
	"zi/shifr"
	// "zi/crypto"
)

func EncryptFile(filename string, r shifr.Shifrator) {
	r.Init()

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close() // отложенное закрытие файла

	encryptedFile, err := os.Create(filename + "." + r.FileType())
	defer encryptedFile.Close()

	keyFile, err := os.Create(filename + "." + r.FileType() + ".key")
	defer keyFile.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return
	}

	var i int64 = 0
	for i = 0; i < stat.Size(); i++ {
		fmt.Print(string(bytes[i]))
		encryptedFile.WriteString(r.EncryptByte(bytes[i]))
	}
	keyFile.Write(r.Key())
}

func main() {
	var r shifr.RSA
	r.C = 15
	EncryptFile("testfile", &r)
	fmt.Println("r.c is", r.C)

	// shifr.RSA_Init(&r)
	// RSA_EncryptFile("testfile", r)
	// encrypted := shifr.RSA_Encrypt(5, r)
	// decrypted := shifr.RSA_Decrypt(encrypted, r)
	// fmt.Println("Rsa check", encrypted, decrypted, 5)
	//
	// shifr.Elgamal(21)
	// shifr.Shamir(24)

}
