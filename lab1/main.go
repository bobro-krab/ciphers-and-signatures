package main

import (
	"fmt"
	"os"
	// "zi/crypto"
	"zi/shifr"
)

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

	// for _, b := range bytes {
	// 	encodedByte := shifr.RSA_Encrypt(b, r)
	// 	encryptedFile.Write(encodedByte)
	// }

}

func main() {
	fmt.Println("hel")

	shifr.Elgamal(3)

}
