package main

import (
	"fmt"
	"os"
	"path/filepath"
	"zi/shifr"
	// "zi/crypto"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DecryptFile(filename string, r shifr.Shifrator) {
	// fmt.Println(filepath.Ext(filename))
	// return
	if filepath.Ext(filename) != "."+r.FileType() {
		fmt.Println("wrong file to decrypt!")
		return
	}

	encryptedFile, err := os.Open(filename)
	check(err)
	defer encryptedFile.Close()

	keyFile, err := os.Open(filename + ".key")
	check(err)
	defer keyFile.Close()

	origFile, err := os.Create(filename + ".orig")
	check(err)
	defer origFile.Close()

	// read key, and set our setting to that
	stat, _ := keyFile.Stat()
	key := make([]byte, stat.Size())
	keyFile.Read(key)
	r.SetKey(key)

	stat, _ = encryptedFile.Stat()
	bytes := make([]byte, stat.Size())
	encryptedFile.Read(bytes)
	var i int64 = 0
	for i = 0; i < stat.Size(); i += int64(r.BlockSize()) {
		db := r.DecryptByte(bytes[i : i+int64(r.BlockSize())])
		origFile.WriteString(string(db))
		fmt.Print(string(db))
	}
}

func EncryptFile(filename string, r shifr.Shifrator) {
	r.Init()

	file, err := os.Open(filename)
	check(err)
	defer file.Close() // отложенное закрытие файла

	encryptedFile, err := os.Create(filename + "." + r.FileType())
	check(err)
	defer encryptedFile.Close()

	keyFile, err := os.Create(filename + "." + r.FileType() + ".key")
	check(err)
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
		encryptedFile.Write(r.EncryptByte(bytes[i]))
	}
	keyFile.Write(r.Key())
}

func main() {
	var r shifr.RSA

	if len(os.Args) > 2 {
		DecryptFile(os.Args[1], &r)
	} else {
		EncryptFile(os.Args[1], &r)
	}

	// shifr.RSA_Init(&r)
	// RSA_EncryptFile("testfile", r)
	// encrypted := shifr.RSA_Encrypt(5, r)
	// decrypted := shifr.RSA_Decrypt(encrypted, r)
	// fmt.Println("Rsa check", encrypted, decrypted, 5)
	//
	// shifr.Elgamal(21)
	// shifr.Shamir(24)

}
