package main

import (
	"fmt"
	"os"
	"zi/shifr"
	// "zi/crypto"
)

type Shifrator interface {
	BlockSize()
	EncryptByte(byte) string
	DecryptByte(string) byte
}

func RSA_EncryptFile(filename string, r shifr.RSA) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close() // отложенное закрытие файла

	encryptedFile, err := os.Open(filename + ".enc")
	defer encryptedFile.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return
	}

}

func main() {
	fmt.Println("hel")

	var r shifr.RSA

	shifr.RSA_Init(&r)
	RSA_EncryptFile("testfile", r)
	// encrypted := shifr.RSA_Encrypt(5, r)
	// decrypted := shifr.RSA_Decrypt(encrypted, r)
	// fmt.Println("Rsa check", encrypted, decrypted, 5)
	//
	// shifr.Elgamal(21)
	// shifr.Shamir(24)

}
