package sign

import (
	"fmt"
	"os"
	"path/filepath"
)

type Shifrator interface {
	Init()
	BlockSize() int
	FileType() string

	EncryptByte(byte) []byte
	DecryptByte([]byte) byte

	Key() []byte
	LoadKey([]byte)
}

/*
Reads filename and save it containts into byte array
*/
func getBytesFromFile(filename string) []byte {
	encryptedFile, err := os.Open(filename)
	check(err)
	defer encryptedFile.Close()

	stat, _ := encryptedFile.Stat()
	bytes := make([]byte, stat.Size())
	encryptedFile.Read(bytes)
	return bytes
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DecryptFile(filename string, r Shifrator) {
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
	r.LoadKey(key)

	stat, _ = encryptedFile.Stat()
	bytes := make([]byte, stat.Size())
	encryptedFile.Read(bytes)

	var i int64 = 0
	for i = 0; i < stat.Size(); i += int64(r.BlockSize()) {
		db := r.DecryptByte(bytes[i : i+int64(r.BlockSize())])
		some := make([]byte, 1)
		some[0] = db
		origFile.Write(some)
		// fmt.Print(string(db))
	}
}

func EncryptFile(filename string, r Shifrator) {
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
		// fmt.Print(string(bytes[i]))
		encryptedFile.Write(r.EncryptByte(bytes[i]))
	}
	keyFile.Write(r.Key())
}
